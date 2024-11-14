package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/godbus/dbus/v5"
)

func setScreenSaverActive(active bool) error {
	// Connect to the session bus
	conn, err := dbus.SessionBus()
	if err != nil {
		return fmt.Errorf("failed to connect to session bus: %v", err)
	}

	// Define the object path and interface
	objectPath := "/org/gnome/ScreenSaver"
	interfaceName := "org.gnome.ScreenSaver"

	// Create a new method call to SetActive
	call := conn.Object(interfaceName, dbus.ObjectPath(objectPath)).Call(interfaceName+".SetActive", 0, active)

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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/on`, onHandler)
	mux.HandleFunc(`/off`, offHandler)

	serverAddr := `:57650`
	err := http.ListenAndServe(serverAddr, mux)
	if err != nil {
		panic(err)
	}
	log.Printf("listening %v", serverAddr)
}
