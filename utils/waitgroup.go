package utils

import (
	"github.com/sirupsen/logrus"
	"runtime/debug"
	"sync"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logrus.Error("WaitGroupWrapper Wrap", err)
				debug.PrintStack()
			}
			w.Done()
		}()
		cb()
	}()
}
