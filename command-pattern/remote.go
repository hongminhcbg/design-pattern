package main
type remote struct {
	Command command
}

func (r *remote) SetCommand(c command)  {
	r.Command = c
}

func (r *remote) ButtonWasPressed()  {
	r.Command.Execute()
}
