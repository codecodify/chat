package vars

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Quit 退出信号
var Quit chan struct{}

func init() {
	var q = make(chan os.Signal)
	Quit = make(chan struct{})
	go func() {
		// 监听到退出信号
		<-q
		fmt.Println("监听到系统退出信号")
		close(Quit)
		time.Sleep(3 * time.Second)
		os.Exit(0)
	}()

	go func() {
		signal.Notify(q, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	}()
}
