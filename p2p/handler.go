package p2p

import (
	"fmt"
	"io"
)

type Handler interface {
	HandleMessage(*Message) error
}

type DefaultHandler struct{}

func (h *DefaultHandler) HandleMessage(msg *Message) error {
	b, err := io.ReadAll(msg.Payload)
	if err != nil {
		return err
	}
	fmt.Printf("Handling msg from %s: , Payload: %s", msg.From, string(b))
	return nil
}
