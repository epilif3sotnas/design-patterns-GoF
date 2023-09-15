package command


type CommandHistory struct {
	commandsHistory []Command
}


func NewCommandHistory(commandsHistory []Command) *CommandHistory {
	return &CommandHistory{
		commandsHistory: commandsHistory,
	}
}

func (self *CommandHistory) Push(command Command) {
	self.commandsHistory = append(self.commandsHistory, command)
}

func (self *CommandHistory) Pop() Command {
	var commandPopped Command = self.commandsHistory[len(self.commandsHistory)-1]
	self.commandsHistory = self.commandsHistory[:len(self.commandsHistory)-1]
	return commandPopped
}

func (self *CommandHistory) GetCommandsHistory() []Command {
	return self.commandsHistory
}

func (self *CommandHistory) SetCommandsHistory(commandsHistory []Command) {
	self.commandsHistory = commandsHistory
}