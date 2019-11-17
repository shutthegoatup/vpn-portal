package cli

import (
	"os"
	"os/signal"
	"syscall"
)

var appContinueChan = make(chan struct{})
var appDoneChan = make(chan struct{})

var signals = []os.Signal{
	syscall.SIGUSR1,
}

// InitSignalHandler handles signal interruption.
func InitSignalHandler() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, signals...)
	go func() {
		for {
			select {
			case <-appDoneChan:
				return
			case sig := <-sigChan:
				switch sig {
				case syscall.SIGUSR1:
					appContinueChan <- struct{}{}
				default:
				}
			}
		}
	}()
}
