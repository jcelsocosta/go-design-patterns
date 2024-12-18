package bridge

import (
	"errors"
	"fmt"
	"io"
)

type PrinterAPI interface {
	PrintMessage(string) error
}

type PrinterImpl1 struct{}

func (p *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

type PrinterAPI2 struct {
	Writer io.Writer
}

func (d *PrinterAPI2) PrintMessage(msg string) error {
	if d.Writer == nil {
		return errors.New("you need to pass an io.Writer to PrinterAPI2")
	}
	fmt.Fprintf(d.Writer, "%s", msg)
	return nil
}

type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *NormalPrinter) Print() error {
	c.Printer.PrintMessage(c.Msg)
	return nil
}

type PackPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *PackPrinter) Print() error {
	c.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", c.Msg))
	return nil
}
