package main

import "fmt"

type Command struct {
	Name string

}
type newCommand Command

func (c Command) GetName() string {
	return c.Name
}

func (c newCommand) GetName() string {
	return c.Name
}

type Commands []Command

func main() {
	var commandvar newCommand
	commandvar.Name  = "redifine the string test111"
	fmt.Println(commandvar.Name)
	fmt.Println(commandvar.GetName())

	var comvar Command
	comvar.Name = "redifine the string test222"
	fmt.Println(comvar.Name)
	fmt.Println(comvar.GetName())

	var comvar1 Commands {
		comvar
	}


	var commands(comvar1)
	fmt.Println(commands.GetName())


}