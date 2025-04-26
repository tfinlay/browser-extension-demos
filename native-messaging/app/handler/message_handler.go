package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/tfinlay/browser-extension-demos/native-messaging/app/store"
)

type loginGetter interface {
	GetLogin(host string) (store.Login, error)
}

type MessageHandler struct {
	loginManager loginGetter
}

func NewMessageHandler(passwordManager loginGetter) *MessageHandler {
	return &MessageHandler{
		loginManager: passwordManager,
	}
}

func (h *MessageHandler) buildLoginResponse(login store.Login, err error) interface{} {
	if err == nil {
		return SuccessReponseMessage{
			Username: login.Username,
			Password: login.Password,
		}
	} else {
		if errors.Is(err, store.ErrUnknownHostname) {
			return ErrorResponseMessage{
				Error: "unknown host",
			}
		} else {
			fmt.Fprintf(os.Stderr, "Caught unhandled error while processing login request: %v\n", err)
			return ErrorResponseMessage{
				Error: "unhandled error",
			}
		}
	}
}

func (h *MessageHandler) HandleMessage(r io.Reader, w io.Writer) error {
	fmt.Fprintln(os.Stderr, "Handling a message!")

	rawMessage, err := readMessage(r)
	if err != nil {
		return fmt.Errorf("failed to read message: %w", err)
	}

	var message GetPasswordMessage
	if err := json.Unmarshal(rawMessage, &message); err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Getting login for hostname %q\n", message.Host)

	res := h.buildLoginResponse(h.loginManager.GetLogin(message.Host))

	if err := writeMessage(w, res); err != nil {
		return fmt.Errorf("failed to write response: %w", err)
	}
	return nil
}
