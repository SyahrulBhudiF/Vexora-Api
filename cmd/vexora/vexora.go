package main

import "github.com/SyahrulBhudiF/Vexora-Api/cmd/commands"

func main() {
	err := commands.Execute()
	if err != nil {
		return
	}
}
