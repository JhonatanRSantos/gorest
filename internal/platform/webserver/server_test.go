package webserver

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
)

type webRouter struct {
	app *fiber.App
}

func (wr *webRouter) Start() {
	group := wr.app.Group("/")
	group.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
}

func TestWebServer(t *testing.T) {
	ws := NewWebServer(DefaultConfig(WebServerDefaultConfig{}))
	ws.AddRoutes(&webRouter{ws.GetApp()})

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("failed to get net listener. Cause: %s", err)
	}

	port := listener.Addr().(*net.TCPAddr).Port
	if err = listener.Close(); err != nil {
		t.Fatalf("failed to close listener. Cause: %s", err)
	}

	go func(ws *WebServer, port int) {
		if err := ws.Listen(fmt.Sprintf(":%d", port)); err != nil {
			panic(fmt.Errorf("failed to close webserver. Cause: %s", err))
		}
	}(ws, port)

	defer func() {
		if err := ws.GetApp().Shutdown(); err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second * 1)

	resp, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
	if err != nil {
		t.Fatalf("failed to GET /. Cause: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("invalid status code when calling GET /. Expected 200 but got %d", resp.StatusCode)
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body from GET / request. Cause: %s", err)
	}

	if string(bs) != "OK" {
		t.Fatalf("invalid response body when calling GET /. Expected OK but got %s", string(bs))
	}

	resp, err = http.Get(fmt.Sprintf("http://localhost:%d/swagger", port))
	if err != nil {
		t.Fatalf("failed to GET /. Cause: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("invalid status code when calling GET /. Expected 200 but got %d", resp.StatusCode)
	}
}
