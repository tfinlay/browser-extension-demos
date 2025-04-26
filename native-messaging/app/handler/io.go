package handler

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
)

func readMessage(r io.Reader) (json.RawMessage, error) {
	var messageLength uint32
	if err := binary.Read(r, binary.NativeEndian, &messageLength); err != nil {
		return nil, fmt.Errorf("failed to read message length: %w", err)
	}

	messageBuf := make([]byte, messageLength)
	if n, err := io.ReadFull(r, messageBuf); err != nil {
		return nil, fmt.Errorf("failed to read data (read %d of expected %d): %w", n, messageLength, err)
	}

	return json.RawMessage(messageBuf), nil
}

func writeMessage(w io.Writer, message interface{}) error {
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	if err := binary.Write(w, binary.NativeEndian, uint32(len(messageBytes))); err != nil {
		return fmt.Errorf("failed to write message length: %w", err)
	}
	if _, err := w.Write(messageBytes); err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}
	return nil
}
