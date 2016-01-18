package main

import "sync"

type singleton struct {
}

var (
	instance *singleton
	once     sync.Once
)

func New() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
