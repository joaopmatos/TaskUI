package main

import (
	"os"
)

func main() {

	//Read the arguments
	arguments := os.Args

	//Check if we have any. If we don't have any arguments, we pass the --help output
	if len(arguments) <= 1 {
		init_tui()
		return
	}

	switch arguments[1] {
	case "list":
		list()
	case "run":
		run()
	case "help":
		help()
	default:
		init_tui()
	}

}

func init_tui() {
	println("TODO: implement TUI!")
}

func help() {
	const TextHelp = "TaskUI: lean Task Scheduler CLI tool.\n\nOptions\nlist: TaskUI list - shows a list of existing tasks\nrun: TaskUI run <id1> <id2> <id...> - runs the tasks enumerated in the parameters (separated by spaces)"

	println(TextHelp)
}

func list() {
	println("TODO: implement list function!")
}

func run() {
	println("TODO: implement list function!")
}

func enable() {
	println("TODO: implement list function!")
}

func disable() {
	println("TODO: implement list function!")
}

func delete() {
	println("TODO: implement list function!")
}
