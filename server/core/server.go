package core

import (
	"context"
	"daigou-admin/global"
	"daigou-admin/initialize"
	"daigou-admin/utils/zaplog"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type server interface {
	ListenAndServe() error
}

func Run() {
	Router := initialize.Routers()
	Router.Static("/form-generator", fmt.Sprintf("%s/resource/form-generator", global.G_Config.Gen.RootPath))

	//InstallPlugs(Router)
	// end 插件描述

	s := &http.Server{
		Addr:           global.G_Config.System.Addr,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zaplog.Fatalf("listen: %s\n", err)
		}
	}()
	zaplog.Info("server run success on ", global.G_Config.System.Addr)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	c := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	for {
		signInfo := <-c
		zaplog.Infof("server get a signal %s", signInfo.String())
		switch signInfo {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			zaplog.Info("shutting down server...")
			// The context is used to inform the server it has 5 seconds to finish
			// the request it is currently handling
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			defer global.G_DBManage.Close()
			defer global.G_CacheManage.Close()
			if err := s.Shutdown(ctx); err != nil {
				zaplog.Fatal("Server forced to shutdown:", err)
			}
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
