package chainofresponsability

import (
	"fmt"
	"strings"
	"testing"
)

type myTestWriter struct {
	receivedMessage *string
}

func (m *myTestWriter) Write(p []byte) (int, error) {
	tmpMessage := string(p)
	m.receivedMessage = &tmpMessage
	return len(p), nil
}

func (m *myTestWriter) Next(s string) {
	m.Write([]byte(s))
}

func TestCreateDefaultCahin(t *testing.T) {
	myWriter := myTestWriter{}

	writerLogger := WriterLogger{Writer: &myWriter}
	second := SecondLogger{NextChain: &writerLogger}
	chain := FirstLogger{NextChain: &second}

	t.Run("3 loggers, 2 of them writes to console, second only if it founds the word 'hello', third wirtes to some variable if second found 'hello'",
		func(t *testing.T) {
			chain.Next("message that breaks the chain \n")

			if myWriter.receivedMessage != nil {
				t.Error("Last link should not receive any message")
			}

			chain.Next("Hello")

			if myWriter.receivedMessage == nil || !strings.Contains(*myWriter.receivedMessage, "Hello") {
				t.Fatal("Last Link didn't received expected message")
			}
		})

	t.Run("2 loggers, second uses to closure implementation", func(t *testing.T) {
		myWriter = myTestWriter{}

		closureLogger := ClosureChain{
			Closure: func(s string) {
				fmt.Printf("My closure logger! Message: %s\n", s)
				myWriter.receivedMessage = &s
			},
		}

		writerLogger.NextChain = &closureLogger

		chain.Next("Hello closure logger")

		if *myWriter.receivedMessage != "Hello closure logger" {
			t.Fatal("Expected message wasn't received in myWriter")
		}
	})
}
