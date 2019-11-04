package main

import (
	"fmt"

	p "github.com/c-bata/go-prompt"
)

func printGreeting() {
	fmt.Println("Welcome to gsh.\nType \"help\" for more information.")
}

func printHelp() {
	fmt.Println(`Usage: [COMMAND] [ARGUMENTS]...
	gsh is an interactive shell for gourd
	
	Commands:
	  echo		display a line of text
	  ping		display PONG
	  quit		exit gsh
	  exit		exit gsh
	`)
}

func getPrefixColor(shell *Shell) p.Color {
	if shell.HasErr() {
		return p.DarkRed
	}

	switch shell.GetPromptColor() {
	case Default:
		{
			return p.DefaultColor
		}
	case Red:
		{
			return p.Red
		}
	case DarkRed:
		{
			return p.DarkRed
		}
	case Blue:
		{
			return p.Blue
		}
	case DarkBlue:
		{
			return p.DarkBlue
		}
	case Green:
		{
			return p.Green
		}
	case Black:
		{
			return p.Black
		}
	case White:
		{
			return p.White
		}
	}

	return p.DefaultColor
}

func setColorFromStr(shell *Shell, s string) {
	switch s {
	case "DEFAULT":
		{
			shell.SetPromptColor(Default)
			break
		}
	case "RED":
		{
			shell.SetPromptColor(Red)
			break
		}
	case "DARKRED":
		{
			shell.SetPromptColor(DarkRed)
			break
		}
	case "BLUE":
		{
			shell.SetPromptColor(Blue)
			break
		}
	case "DARKBLUE":
		{
			shell.SetPromptColor(DarkBlue)
			break
		}
	case "GREEN":
		{
			shell.SetPromptColor(Green)
			break
		}
	case "BLACK":
		{
			shell.SetPromptColor(Black)
			break
		}
	case "WHITE":
		{
			shell.SetPromptColor(White)
			break
		}
	}
}

func colorToDisplay(c Color) string {
	switch c {
	case Default:
		{
			return "default"
		}
	case Red:
		{
			return "red"
		}
	case DarkRed:
		{
			return "darkred"
		}
	case Blue:
		{
			return "blue"
		}
	case DarkBlue:
		{
			return "darkblue"
		}
	case Green:
		{
			return "green"
		}
	case Black:
		{
			return "black"
		}
	case White:
		{
			return "white"
		}
	}
	return "unknown"
}
