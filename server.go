package main

import (
	"fmt"
	"bytes"
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var upgrader = websocket.Upgrader{}
var socks []*websocket.Conn

// serve index.html
func serveIndex(c echo.Context) error {
	// index.html packaged in static.go by go-bindata
	data, err := Asset("index.html")
	if err != nil {
		return err
	}
	return c.HTML(http.StatusOK, string(data))
}

// serve static assets
func serveStatic(c echo.Context) error {
	// static assets packaged in static.go by go-bindata
	data, err := Asset(c.Param("asset"))
	if err != nil {
		return err
	}
	return c.HTML(http.StatusOK, string(data))
}

// establish websocket
func serveSocket(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	// collect websockets for later broadcasting
	socks = append(socks, ws)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// print client messages
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}

// receive and handle gamestate events
func catchEvent(c echo.Context) error {
	// TODO: validate json before broadcast
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request().Body)
	broadcastEvent(c, buf.String())
	return c.String(http.StatusOK, "ok")
}

// broadcast event to all websockets
func broadcastEvent(c echo.Context, event string) {
	// TODO: consider implementing event detection here
	// TODO: clean up dead websockets
	for _, sock := range socks {
		err := sock.WriteMessage(websocket.TextMessage, []byte(event))
		if err != nil {
			c.Logger().Error(err)
		}
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// event collection endpoint
	e.POST("/", catchEvent)
	// event feed endpoint, spawns websocket /ws
	e.GET("/", serveIndex)
	// serves websocket
	e.GET("/ws", serveSocket)
	// serves static assets
	e.GET("/static/:asset", serveStatic)
	e.Logger.Fatal(e.Start(":3000"))
}