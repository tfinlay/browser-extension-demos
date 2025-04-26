package main

import (
	"fmt"
	"os"

	"github.com/tfinlay/browser-extension-demos/native-messaging/app/handler"
	"github.com/tfinlay/browser-extension-demos/native-messaging/app/password"
)

// For simplicity, I'm building this exclusiely for connectionless messaging, so an instance of the application is started for each message,
func main() {
	passwordManager := password.NewLoginManager()
	messageHandler := handler.NewMessageHandler(passwordManager)

	if err := messageHandler.HandleMessage(os.Stdin, os.Stdout); err != nil {
		panic(fmt.Errorf("failed to handle message: %w", err))
	}
}
