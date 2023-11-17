package bot

import (
	"fmt"
	"github.com/lxzan/gws"
	"jjbot/core/logger"
	"net/http"
	"time"
)

func websocket(addr, path, accessToken string) {
	ws := WebSocket{
		Addr:        addr,
		Path:        path,
		AccessToken: accessToken,
	}
	header := http.Header{}
	if accessToken != "" {
		header.Set("Authorization", "Bearer "+accessToken)
	}

	logger.SugarLogger.Infof("[%s]连接OneBot中...", path)
	socket, _, err := gws.NewClient(&ws, &gws.ClientOption{
		Addr:          fmt.Sprintf("ws://%s/%s", addr, path),
		RequestHeader: header,
	})
	if err != nil {
		logger.SugarLogger.Infof("[%s]连接OneBot失败, 5秒后尝试重连...(%s)", path, err)
		ws.Restart(5)
		return
	}
	logger.SugarLogger.Infof("[%s]连接成功", path)
	go socket.ReadLoop()
}

type WebSocket struct {
	Addr        string
	Path        string
	AccessToken string
}

func (c *WebSocket) OnClose(socket *gws.Conn, err error) {
	logger.SugarLogger.Infof("[%s]连接断开, 5秒后尝试重连...(%s)", c.Path, err)
	c.Restart(5)
}

func (c *WebSocket) OnPong(socket *gws.Conn, payload []byte) {
}

func (c *WebSocket) OnOpen(socket *gws.Conn) {
}

func (c *WebSocket) OnPing(socket *gws.Conn, payload []byte) {
}

func (c *WebSocket) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()
	fmt.Printf("recv: %s\n", message.Data.String())
}

func (c *WebSocket) Restart(sleep ...time.Duration) {
	if len(sleep) > 0 {
		time.Sleep(sleep[0] * time.Second)
	}
	websocket(c.Addr, c.Path, c.AccessToken)
}
