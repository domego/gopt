package main

import (
  "syscall"
  "time"
  "os"

  "github.com/fvbock/endless"
  "github.com/gin-gonic/gin"
  "github.com/domego/ginkits/middleware"
  "github.com/domego/gokits/log"
)

func handleApp() {
	gin.SetMode(conf.Env)
	route := gin.New()
	route.Use(middleware.Logger())
	route.Use(gin.RecoveryWithWriter(os.Stderr))
	registRoute(route)
	endless.DefaultHammerTime = time.Second * 2
	server := endless.NewServer(conf.Address, route)
	// 捕获进程USR1信号，reloadConfig
	server.SignalHooks[endless.PRE_SIGNAL][syscall.SIGUSR1] = append(server.SignalHooks[endless.PRE_SIGNAL][syscall.SIGUSR1], func() {
		handleReloadSignal()
	})
	server.SetKeepAlivesEnabled(true)

	if conf.Env != gin.DebugMode {
		log.Infof("%s: Listening and serving HTTP on %s", appName, conf.Address)
	}
	err := server.ListenAndServe()
	if err != nil {
		handleEnd()
		log.Errorf("%s", err)
	}
}
