package main

import (
	"github.com/animenotifier/arn"
	"github.com/blitzprog/color"
)

func main() {
	defer arn.Node.Close()

	color.Yellow("Deleting all sessions...")
	arn.DB.Clear("Session")
	color.Green("Finished.")
}
