package run

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(r *gin.Engine, serverName string, addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	go func() {
		logrus.Printf("%s running in %s \n", serverName, srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	//SIGINT 用户发送INTR字符(Ctrl+C)触发
	//SIGTERM 结束程序(可以被捕获、阻塞或忽略)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Printf("Shutting Down project %s... \n", serverName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatalf("%s Shutdown, cause by : %v \n", serverName, err)
	}
	select {
	case <-ctx.Done():
		logrus.Println("关闭超时")
	}
	logrus.Printf("%s stop success...", serverName)
}
