package server

import (
	"Lv/internal/bootstrap"
	"Lv/router"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

func SetupRouter()(r *gin.Engine){
	r = gin.Default()
	router.RouteProvider(r)

	return r
}

func CreateServer(router *gin.Engine, port, host string, args ...string) *http.Server {
	// 建立 Server
	server := &http.Server{
		Addr:    port,
		Handler: router,
		// ReadTimeout:  5 * time.Second,
		// WriteTimeout: 10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	return server
}

// SignalListenAndServe 開啟Server & 系統信號監聽
func SignalListenAndServe(server *http.Server) (ctx context.Context) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	defer func() {
		if err := recover(); err != nil {
			errMessage := fmt.Sprintf("❌  Server 發生意外 Error: %v ❌", err)
			bootstrap.WriteLog("ERROR", errMessage)
		}
	}()

	l, err := net.Listen("tcp", server.Addr)
	if err != nil {
		bootstrap.WriteLog("ERROR", fmt.Sprintf("❌  Server 建立監聽連線失敗 (%v) ❌", err))
		return
	}

	go func() {
		receivedSignal := <-bootstrap.GracefulDown()
		bootstrap.WriteLog("INFO", fmt.Sprintf("🎃  接受訊號 <- %v 🎃", receivedSignal))
		server.Close()
	}()

	bootstrap.WriteLog("INFO", "🐳  Web Server 開始服務! "+l.Addr().String()+"🐳")
	defer bootstrap.WriteLog("INFO", "🔥  Web Server 結束服務!🔥")
	err = server.Serve(l)
	bootstrap.WriteLog("WARNING", fmt.Sprintf("🎃  Server 回傳 error (%v) 🎃", err))

	return
}