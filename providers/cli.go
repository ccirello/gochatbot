package providers // import "cirello.io/gobot/providers"

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"cirello.io/gobot/messages"
)

var (
	stdin     io.Reader = os.Stdin
	inPrompt  io.Writer = os.Stdout
	outPrompt io.Writer = os.Stdout
)

type providerCLI struct {
	in  chan messages.Message
	out chan messages.Message
}

// CLI is the message provider meant to be used in development of rule sets.
func CLI() *providerCLI {
	cli := &providerCLI{
		in:  make(chan messages.Message),
		out: make(chan messages.Message),
	}
	go cli.loop()
	return cli
}

func (c *providerCLI) IncomingChannel() chan messages.Message {
	return c.in
}

func (c *providerCLI) OutgoingChannel() chan messages.Message {
	return c.out
}

func (c *providerCLI) loop() {
	go func() {
		fmt.Fprint(inPrompt, "in:> ")
		scanner := bufio.NewScanner(stdin)
		for scanner.Scan() {
			c.in <- messages.Message{
				Room:         "CLI",
				FromUserID:   "",
				FromUserName: "",
				Message:      scanner.Text(),
			}
			fmt.Fprint(inPrompt, "in:> ")
		}
	}()
	go func() {
		for msg := range c.out {
			fmt.Fprintln(outPrompt, "\nout:>", msg.Room, msg.FromUserID, msg.FromUserName, ":", msg.Message)
		}
	}()
}