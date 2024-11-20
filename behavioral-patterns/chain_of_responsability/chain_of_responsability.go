package chainofresponsability

import (
	"fmt"
	"io"
	"strings"
)

type ChainLogger interface {
	Next(string)
}

type FirstLogger struct {
	NextChain ChainLogger
}

func (firstLogger *FirstLogger) Next(s string) {
	fmt.Printf("First Logger: %s\n", s)

	if firstLogger.NextChain != nil {
		firstLogger.NextChain.Next(s)
	}
}

type SecondLogger struct {
	NextChain ChainLogger
}

func (secondLogger *SecondLogger) Next(s string) {
	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Printf("Second Logger: %s\n", s)

		if secondLogger.NextChain != nil {
			secondLogger.NextChain.Next(s)
		}

		return
	}

	fmt.Printf("Finishing in second logging\n\n")
}

type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (writterLogger *WriterLogger) Next(s string) {
	if writterLogger.Writer != nil {
		writterLogger.Writer.Write([]byte("WriterLogger: " + s))
	}

	if writterLogger.NextChain != nil {
		writterLogger.NextChain.Next(s)
	}
}

type ClosureChain struct {
	NextChain ChainLogger
	Closure   func(string)
}

func (closureMain *ClosureChain) Next(s string) {
	if closureMain.Closure != nil {
		closureMain.Closure(s)
	}

	if closureMain.NextChain != nil {
		closureMain.Next(s)
	}
}
