package main

import "github.com/keysonZZZ/kgo/kgoConsole"

func main() {
	kgoConsole.AddAction(kgoConsole.Command{
		Name:   "build",
		Runner: buildCmd,
	})

	kgoConsole.Main()
}
