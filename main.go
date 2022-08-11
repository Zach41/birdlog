package main

import (
	"github.com/Zach41/birdlog/application"
	_ "github.com/Zach41/birdlog/application/help"
	_ "github.com/Zach41/birdlog/application/set"
	_ "github.com/Zach41/birdlog/application/show"
)

func main() {
	// app := application.NewAnalyzier()
	// p := prompt.New(app.Execute, app.GetSuggestions)

	// p.Run()
	// cmd := show.Command()
	// cmd.Execute()
	p := application.NewAnalyzier()
	p.Run()
}
