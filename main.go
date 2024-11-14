package main

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"log"
)

func lockScreen() error {
	// Connect to the session bus
	conn, err := dbus.SessionBus()
	if err != nil {
		return fmt.Errorf("failed to connect to session bus: %v", err)
	}

	// Define the object path and interface
	objectPath := "/org/gnome/ScreenSaver"
	interfaceName := "org.gnome.ScreenSaver"

	// Create a new method call to Lock
	call := conn.Object(interfaceName, dbus.ObjectPath(objectPath)).Call(interfaceName+".Lock", 0)

	// Check for errors in the method call
	if call.Err != nil {
		return fmt.Errorf("failed to call Lock method: %v", call.Err)
	}

	return nil
}

func main() {
	if err := lockScreen(); err != nil {
		log.Fatalf("Error locking screen: %v", err)
	} else {
		fmt.Println("Screen locked successfully.")
	}
}
