package main

import (
	"fmt"
	"strings"

	p "github.com/c-bata/go-prompt"
)

func main() {
	shell := NewShell(NewShellOpts{
		PromptText:  ">>> ",
		PromptColor: DarkBlue,
	})

	defer shell.Die(0)
	printGreeting()

	var ctrlCBinding = p.KeyBind{
		Key: p.ControlC,
		Fn: func(b *p.Buffer) {
			shell.Die(0)
		},
	}

	var staticOptions []p.Option = []p.Option{
		p.OptionLivePrefix(promptPrefix(&shell)),
		p.OptionInputTextColor(p.Cyan),

		p.OptionAddKeyBind(ctrlCBinding),

		p.OptionPreviewSuggestionTextColor(p.Green),
		p.OptionSuggestionBGColor(p.Black),
		p.OptionSuggestionTextColor(p.White),
		p.OptionSelectedSuggestionBGColor(p.White),
		p.OptionSelectedSuggestionTextColor(p.Black),
		p.OptionDescriptionBGColor(p.Black),
		p.OptionDescriptionTextColor(p.DarkBlue),
		p.OptionSelectedDescriptionBGColor(p.White),
		p.OptionSelectedDescriptionTextColor(p.Black),
	}

	for {
		options := append([]p.Option{
			p.OptionPrefixTextColor(getPrefixColor(&shell)),
		}, staticOptions...)

		input := p.Input("", completer, options...)
		args := strings.Split(strings.Trim(input, " "), " ")
		shell.ClearErr()

		if len(args) < 1 {
			continue
		}

		command := strings.ToUpper(args[0])
		if command == "" {
			continue
		}

		if command[0] == '@' {
			builtInCommand(&shell, args[0], command, args[1:])
			continue
		}

		switch command {
		case "ECHO":
			{
				msg := strings.Join(args[1:], " ")
				fmt.Println(shell.Client().Echo(msg))
				continue
			}
		case "PING":
			{
				res := shell.Client().Ping()
				fmt.Println(res)
				continue
			}
		case "?", "HELP":
			{
				printHelp()
				continue
			}
		case "EXIT", "Q", "QUIT":
			{
				return
			}
		}

		shell.SetErrF("command not found: %s\n", args[0])
	}
}

func builtInCommand(s *Shell, input string, command string, args []string) {
	argCount := len(args)

	switch command {
	case "@SET":
		{
			if argCount < 2 {
				s.SetErrF("invalid amount of arguments. @set requires atleast 2 but got %d\n", argCount)
				return
			}

			target, text := strings.ToUpper(strings.Trim(args[0], " ")), strings.Trim(args[1], " ")
			if target == "PROMPT" {
				if text[len(text)-1] != ' ' {
					s.SetPromptText(text + " ")
				} else {
					s.SetPromptText(text + " ")
				}

				if argCount == 3 {
					setColorFromStr(s, strings.ToUpper(strings.Trim(args[2], " ")))
				}
			} else {
				s.SetErrF("invalid target for @set: %s\n", strings.Trim(args[0], " "))
			}

			return
		}

	case "@LET":
		{

			return
		}
	}

	s.SetErrF("invalid bultin command: %s\n", input)
}

func promptPrefix(s *Shell) func() (string, bool) {
	return func() (string, bool) { return s.GetPromptText(), true }
}

func completer(d p.Document) []p.Suggest {
	s := []p.Suggest{
		{Text: "echo", Description: "display a line of text"},
		{Text: "ping", Description: "display PONG"},
	}
	return p.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
