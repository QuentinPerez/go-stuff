package autosock

import (
	"fmt"
	"net"
	"sync"
	"sync/atomic"

	"github.com/Sirupsen/logrus"
)

const (
	closed = iota
	shoulBeClosed
	connected
)

type AutoReconnect struct {
	state          uint32 // close or not the connection
	conn           net.Conn
	reConnect      func() (net.Conn, error)
	waitWorker     sync.WaitGroup
	closeChan      chan struct{}
	closeReadChan  chan struct{}
	closeWriteChan chan struct{}
	readChan       chan struct{}
	failChan       chan struct{}
	failReadChan   chan struct{}
	connectChan    chan struct{}
}

func New(initConnection func() (net.Conn, error)) *AutoReconnect {
	ret := &AutoReconnect{
		reConnect:      initConnection,
		closeChan:      make(chan struct{}),
		closeReadChan:  make(chan struct{}),
		closeWriteChan: make(chan struct{}),
		readChan:       make(chan struct{}),
		failChan:       make(chan struct{}),
		failReadChan:   make(chan struct{}),
		connectChan:    make(chan struct{}),
	}
	atomic.StoreUint32(&ret.state, closed)
	ret.waitWorker.Add(4)
	go ret.read()
	go ret.write()
	go ret.reconnect()
	go ret.fail()
	ret.connectChan <- struct{}{}
	return ret
}

func (a *AutoReconnect) read() {
	defer a.waitWorker.Done()

	defer fmt.Println("return read()")

	buff := make([]byte, 128)

	for {
		select {
		case _, ok := <-a.closeReadChan:
			if ok {
				logrus.Fatal("autoconnect.Read(): An unexcepted error occured")
			}
			return
		case <-a.readChan:
			_, err := a.conn.Read(buff)
			if err != nil {
				atomic.StoreUint32(&a.state, shoulBeClosed)
				logrus.Error(err)
				a.failReadChan <- struct{}{}
			}
		}
	}
}

func (a *AutoReconnect) write() {
	defer a.waitWorker.Done()

	defer fmt.Println("return write()")
	for {
		select {
		case _, ok := <-a.closeWriteChan:
			if !ok {
				return
			}
		}
	}
}

func (a *AutoReconnect) fail() {
	defer a.waitWorker.Done()

	defer fmt.Println("return fail()")

	for {
		select {
		case _, ok := <-a.failChan:
			if !ok {
				close(a.connectChan)
				return
			}
			if atomic.LoadUint32(&a.state) == shoulBeClosed {
				a.conn.Close()
				atomic.StoreUint32(&a.state, closed)
			}
			a.connectChan <- struct{}{}
		case <-a.failReadChan:
			if atomic.LoadUint32(&a.state) == shoulBeClosed {
				a.conn.Close()
				atomic.StoreUint32(&a.state, closed)
			}
			a.connectChan <- struct{}{}
		}
	}
}

func (a *AutoReconnect) reconnect() {
	defer a.waitWorker.Done()

	defer fmt.Println("return reconnect()")

	var err error

	for {
		select {
		case <-a.closeChan:
			fmt.Println("Close readChan")
			close(a.closeReadChan)
			fmt.Println("Close failChan")
			close(a.failChan)
			fmt.Println("Close closeChan")
			close(a.closeChan)
			for {
				select {
				case _, ok := <-a.connectChan:
					if !ok {
						return
					}
				default:
					break
				}
			}
		case <-a.connectChan:
			a.conn, err = a.reConnect()
			if err != nil {
				logrus.Info(err)
				a.failChan <- struct{}{}
			} else {
				atomic.StoreUint32(&a.state, connected)
				a.readChan <- struct{}{}
			}
		}
	}
}

func (a *AutoReconnect) Close() {
	a.closeChan <- struct{}{}
	if _, ok := <-a.closeChan; ok {
		logrus.Fatal("autoconnect.Close(): An unexcepted error occured")
	}
	a.waitWorker.Wait()
}
