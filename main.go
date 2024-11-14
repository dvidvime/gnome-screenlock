package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/godbus/dbus/v5"
)

const (
	serverAddr    = `:57650`
	objectPath    = "/org/gnome/ScreenSaver"
	interfaceName = "org.gnome.ScreenSaver"
)

// Function to get the current status of the screen saver
func getScreenSaverStatus() (bool, error) {
	// Connect to the session bus
	conn, err := dbus.SessionBus()
	if err != nil {
		return false, fmt.Errorf("failed to connect to session bus: %v", err)
	}

	// Create a new method call to GetActive
	var active bool
	call := conn.Object(interfaceName, objectPath).Call(interfaceName+".GetActive", 0)

	// Check for errors in the method call
	if call.Err != nil {
		return false, fmt.Errorf("failed to call GetActive method: %v", call.Err)
	}

	// Unmarshal the result into the active variable
	if err := call.Store(&active); err != nil {
		return false, fmt.Errorf("failed to store result: %v", err)
	}

	return active, nil
}

func setScreenSaverActive(active bool) error {
	// Connect to the session bus
	conn, err := dbus.SessionBus()
	if err != nil {
		return fmt.Errorf("failed to connect to session bus: %v", err)
	}

	// Create a new method call to SetActive
	call := conn.Object(interfaceName, objectPath).Call(interfaceName+".SetActive", 0, active)

	// Check for errors in the method call
	if call.Err != nil {
		return fmt.Errorf("failed to call SetActive method: %v", call.Err)
	}

	return nil
}

func onHandler(res http.ResponseWriter, _ *http.Request) {
	// Set the screen saver active (true)
	if err := setScreenSaverActive(true); err != nil {
		http.Error(res, "Error setting screen saver active", http.StatusInternalServerError)
		log.Printf("Error setting screen saver active: %v", err)
	} else {
		res.Header().Set("content-type", "text/plain; charset=UTF-8")
		res.WriteHeader(http.StatusOK)
		fmt.Println("Screen saver activated successfully.")
		_, err := res.Write([]byte("Screen saver activated successfully."))
		if err != nil {
			log.Printf("failed to write response: %v", err)
			return
		}
	}
}

func offHandler(res http.ResponseWriter, _ *http.Request) {
	// Set the screen saver inactive (false)
	if err := setScreenSaverActive(false); err != nil {
		http.Error(res, "Error setting screen saver inactive", http.StatusInternalServerError)
		log.Printf("Error setting screen saver inactive: %v", err)
	} else {
		res.Header().Set("content-type", "text/plain; charset=UTF-8")
		res.WriteHeader(http.StatusOK)
		fmt.Println("Screen saver inactivated successfully.")
		_, err := res.Write([]byte("Screen saver inactivated successfully."))
		if err != nil {
			log.Printf("failed to write response: %v", err)
			return
		}
	}
}

func statusHandler(res http.ResponseWriter, _ *http.Request) {
	//Get the screen saver status
	if active, err := getScreenSaverStatus(); err != nil {
		http.Error(res, "Error fetching screen saver status", http.StatusInternalServerError)
		log.Printf("Error fetching screen saver status: %v", err)
	} else {
		res.Header().Set("content-type", "text/plain; charset=UTF-8")
		res.WriteHeader(http.StatusOK)
		fmt.Printf("Screen saver status %v.\n", active)
		_, err := res.Write([]byte(strconv.FormatBool(active)))
		if err != nil {
			log.Printf("failed to write response: %v", err)
			return
		}
	}
}

func main() {
	fmt.Printf("Running at %v\n", serverAddr)

	mux := http.NewServeMux()
	mux.HandleFunc(`/on`, onHandler)
	mux.HandleFunc(`/off`, offHandler)
	mux.HandleFunc(`/status`, statusHandler)

	err := http.ListenAndServe(serverAddr, mux)
	if err != nil {
		panic(err)
	}
}
