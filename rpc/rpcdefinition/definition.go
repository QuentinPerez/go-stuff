package rpcdefinition

import "time"

type Call struct {
	NbOfCall int // no guard :/
}

type Args struct {
	Start time.Time
}

type Result struct {
	NB     int
	Start  time.Time
	Middle time.Time
}

func (c *Call) FastCall(a *Args, res *Result) (err error) {
	c.NbOfCall++
	*res = Result{
		c.NbOfCall,
		a.Start,
		time.Now(),
	}
	return
}

func (c *Call) SlowCall(a *Args, res *Result) (err error) {
	c.NbOfCall++
	time.Sleep(1 * time.Second) // I waste your time :)
	*res = Result{
		c.NbOfCall,
		a.Start,
		time.Now(),
	}
	return
}
