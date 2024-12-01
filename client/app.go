package main

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
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
