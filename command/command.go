package main

type Command interface {
	Execute()
}

type Invoker struct {
	commands chan Command
}

func NewInvoker() Invoker {
	return Invoker{
		commands: make(chan Command, 100),
	}
}

func (s *Invoker) SetCommand(command Command) (isOk bool) {
	defer func() {
		if r := recover(); r != nil {
			isOk = false
			return
		}
	}()

	s.commands <- command
	isOk = true

	return
}

func (s *Invoker) Run() {
	for command := range s.commands {
		command.Execute()
	}
}

func (s *Invoker) Stop() {
	close(s.commands)
}
