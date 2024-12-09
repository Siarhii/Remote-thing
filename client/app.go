package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// App struct for system actions
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Shutdown system after specified timer
func Shutdown(timer uint64) error {
	switch runtime.GOOS {
	case "windows", "linux", "darwin":
		duration := time.Duration(timer) * time.Minute
		time.Sleep(duration)
		
		switch runtime.GOOS {
		case "windows":
			return exec.Command("shutdown", "/s", "/t", "0").Run()
		case "linux", "darwin":
			return exec.Command("sudo", "shutdown", "now").Run()
		}
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)

	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

// Sleep system after specified timer
func Sleep(timer uint64) error {
	switch runtime.GOOS {
	case "windows", "linux", "darwin":
		duration := time.Duration(timer) * time.Minute
		time.Sleep(duration)
	
		switch runtime.GOOS {
		case "windows":
			return exec.Command("rundll32", "powrprof.dll,SetSuspendState", "0", "1", "0").Run()
		case "linux", "darwin":
			return exec.Command("systemctl", "suspend").Run()
		}
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)

	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

// Restart system after specified timer
func Restart(timer uint64) error {
	switch runtime.GOOS {
	case "windows", "linux", "darwin":
		duration := time.Duration(timer) * time.Minute
		time.Sleep(duration)
	
		switch runtime.GOOS {
		case "windows":
			return exec.Command("shutdown", "/r", "/t", "0").Run()
		case "linux", "darwin":
			return exec.Command("sudo", "reboot").Run()
		}
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)

	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

// Message struct to match server-side definition
type Message struct {
	Event   string `json:"event"`
	Content string `json:"content"`
}

// WebSocket connection details
type ClientConnection struct {
	conn     *websocket.Conn
	deviceID string
	userID   string
	done     chan struct{}
	mu       sync.Mutex
	app      *App
}

// sendMessage is a thread-safe method to send messages
func (c *ClientConnection) sendMessage(msg Message) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.conn.WriteJSON(msg)
}

// processSystemAction handles system actions based on received commands
func (c *ClientConnection) processSystemAction(action string, timer string) error {
	num, err := strconv.ParseUint(timer, 10, 16)
	if err != nil {
		log.Println("Error parsing timer:", err)
		return err
	}

	var responseContent string
	switch action {
	case "shutdown":
		go func() {
			err = Shutdown(num)
			if err != nil {
				log.Printf("Shutdown failed: %v", err)
				responseContent = fmt.Sprintf("Shutdown failed: %v", err)
			} else {
				responseContent = "System shutdown initiated"
			}
		}()
	case "sleep":
		go func() {
			err = Sleep(num)
			if err != nil {
				log.Printf("Sleep failed: %v", err)
				responseContent = fmt.Sprintf("Sleep failed: %v", err)
			} else {
				responseContent = "System sleep initiated"
			}
		}()
	case "restart":
		go func() {
			err = Restart(num)
			if err != nil {
				log.Printf("Restart failed: %v", err)
				responseContent = fmt.Sprintf("Restart failed: %v", err)
			} else {
				responseContent = "System restart initiated"
			}
		}()
	default:
		return fmt.Errorf("unsupported action: %s", action)
	}

	// Send response back to the server
	responseMsg := Message{
		Event:   "SystemActionResponse",
		Content: responseContent,
	}
	return c.sendMessage(responseMsg)
}

// messageRouter reads all incoming messages and routes them to appropriate handlers
func (c *ClientConnection) messageRouter() {
	for {
		select {
		case <-c.done:
			return
		default:
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				close(c.done)
				return
			}

			var msg Message
			err = json.Unmarshal(message, &msg)
			if err != nil {
				log.Println("Unmarshal error:", err)
				continue
			}

			switch msg.Event {
			case "Ping":
				// Immediate pong response
				pongMsg := Message{
					Event:   "Pong",
					Content: "testing",
				}
				err = c.sendMessage(pongMsg)
				if err != nil {
					log.Println("Pong write error:", err)
					close(c.done)
					return
				}
				log.Println("Received Ping, sent Pong")

			case "Command":
				log.Printf("Received Command: %s", msg.Content)
				
				// Parse the Command message (assuming format: "action:timer")
				parts := strings.Split(msg.Content, ":")
				if len(parts) != 2 {
					log.Println("Invalid command format")
					continue
				}
				
				action, timer := parts[0], parts[1]
				err = c.processSystemAction(action, timer)
				if err != nil {
					log.Printf("System action error: %v", err)
				}

			case "Stat":
				log.Printf("Received Stat request: %s", msg.Content)
				
				statsResponseMsg := Message{
					Event:   "StatResponse",
					Content: fmt.Sprintf("Stat response for: %s", msg.Content),
				}
				err = c.sendMessage(statsResponseMsg)
				if err != nil {
					log.Println("Stats response write error:", err)
					close(c.done)
					return
				}
				log.Println("Sent StatResponse")

			default:
				log.Printf("Received unexpected message: %+v", msg)
			}
		}
	}
}

func (a *App) ActionFromClient(action string,timer string) error {
	var err error
	num, err := strconv.ParseUint(timer, 10, 16)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Printf("Num : %v",num)
	switch action {
	case "shutdown":
		err = Shutdown(num)
	case "sleep":
		err = Sleep(num)
	case "restart":
		err = Restart(num)
	default:
		return fmt.Errorf("unsupported action: %s", action)
	}

	if err != nil {
		return fmt.Errorf("%s failed: %v, unsupported operating system: %s", action, err, runtime.GOOS)
	}
	return nil
}

// ConnectToServer function establishes the WebSocket connection
func (app *App) ConnectToServer(deviceID string, result func(string)) {
	// Server WebSocket endpoint
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/connect"}

	// Add query parameters for deviceID and userID
	query := u.Query()
	query.Set("deviceID", deviceID)
	query.Set("userID", "user1")
	u.RawQuery = query.Encode()

	log.Printf("Connecting to %s", u.String())

	// Create WebSocket connection
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		// Send error message back to the frontend if connection fails
		result(fmt.Sprintf("Error connecting to WebSocket: %v", err))
		return
	}
	defer c.Close()

	// Send success message back to frontend
	result(fmt.Sprintf("Successfully connected to WebSocket with deviceID: %s", deviceID))

	// Create client connection struct
	clientConn := &ClientConnection{
		conn:     c,
		deviceID: deviceID,
		userID:   "user1",
		done:     make(chan struct{}),
		app:      app,
	}

	// Channel to handle interrupt signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Single goroutine to handle all incoming messages
	go clientConn.messageRouter()

	// Wait for interrupt signal
	<-interrupt

	// Cleanly close the connection
	log.Println("Interrupt received, closing connection")
	close(clientConn.done)
	err = c.WriteMessage(websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("write close error:", err)
	}
}


// func main() {
// 	// Create App instance
// 	app := NewApp()

// 	// Server WebSocket endpoint
// 	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/connect"}

// 	// Add query parameters for deviceID and userID
// 	query := u.Query()
// 	query.Set("deviceID", "cli1")
// 	query.Set("userID", "user1")
// 	u.RawQuery = query.Encode()

// 	log.Printf("Connecting to %s", u.String())

// 	// Create WebSocket connection
// 	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
// 	if err != nil {
// 		log.Fatal("dial:", err)
// 	}
// 	defer c.Close()

// 	// Create client connection struct
// 	clientConn := &ClientConnection{
// 		conn:     c,
// 		deviceID: "cli1",
// 		userID:   "user1",
// 		done:     make(chan struct{}),
// 		app:      app,
// 	}

// 	// Channel to handle interrupt signal
// 	interrupt := make(chan os.Signal, 1)
// 	signal.Notify(interrupt, os.Interrupt)

// 	// Single goroutine to handle all incoming messages
// 	go clientConn.messageRouter()

// 	// Wait for interrupt signal
// 	<-interrupt

// 	// Cleanly close the connection
// 	log.Println("Interrupt received, closing connection")
// 	close(clientConn.done)
// 	err = c.WriteMessage(websocket.CloseMessage, 
// 		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
// 	if err != nil {
// 		log.Println("write close error:", err)
// 	}
// }