package main

import "github.com/aztecher/vellun/daemon/cmd"

func main() {
	cmd.Execute(cmd.NewAgentCmd())
}
