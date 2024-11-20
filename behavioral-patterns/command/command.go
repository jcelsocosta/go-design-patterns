package main

import "fmt"

type Command interface {
	Execute()
}

type ConsoleOutput struct {
	message string
}

func (cmd *ConsoleOutput) Execute() {
	fmt.Println(cmd.message)
}

func CreateCommand(s string) Command {
	fmt.Println("Creating command")

	return &ConsoleOutput{
		message: s,
	}
}

type CommandQueue struct {
	queue []Command
}

func (p *CommandQueue) AddCommand(c Command) {
	p.queue = append(p.queue, c)

	if len(p.queue) == 3 {
		for _, command := range p.queue {
			command.Execute()
		}

		p.queue = make([]Command, 3)
	}
}

func main(){
	queue := CommandQueue{}

	queue.AddCommand(CreateCommand("First Message"))
	queue.AddCommand(CreateCommand("Second Message"))
	queue.AddCommand(CreateCommand("Third Message"))

	queue.AddCommand(CreateCommand("Fourth Message"))
	queue.AddCommand(CreateCommand("Fifth Message"))
}