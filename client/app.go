package main

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

// App struct
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


func Shutdown(timer uint) error {
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

func Sleep(timer uint) error {
	
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

func Restart(timer uint) error {
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


// func (a *App) ControlFromClient(action string) error {
// 	var err error
// 	switch action {
// 	case "shutdown":
// 		err = Shutdown()
// 	case "sleep":
// 		err = Sleep()
// 	case "restart":
// 		err = Restart()
// 	default:
// 		return fmt.Errorf("unsupported action: %s", action)
// 	}

// 	if err != nil {
// 		return fmt.Errorf("%s failed: %v, unsupported operating system: %s", action, err, runtime.GOOS)
// 	}
// 	return nil
// }

func (a *App) ActionFromClient(action string,timer uint) error {
	var err error
	switch action {
	case "shutdown":
		err = Shutdown(timer)
	case "sleep":
		err = Sleep(timer)
	case "restart":
		err = Restart(timer)
	default:
		return fmt.Errorf("unsupported action: %s", action)
	}

	if err != nil {
		return fmt.Errorf("%s failed: %v, unsupported operating system: %s", action, err, runtime.GOOS)
	}
	return nil
}
