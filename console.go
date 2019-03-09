package prompt

import (
	"errors"
	"fmt"
	"strings"

	console "github.com/Songmu/prompter"
)

type ConsolePrompter struct{}

func (c ConsolePrompter) YesNo(message string, opts ...Option) (bool, error) {
	choice, err := c.Choose(message, []string{"y", "n", "cancel"}, opts...)
	if err != nil {
		return false, err
	}
	switch choice {
	case "Y", "y":
		return true, nil
	case "N", "n":
		return false, nil
	default:
		return false, ErrUserCancel
	}
}

func (c ConsolePrompter) Prompt(message string, opts ...Option) (string, error) {
	return c.prompt(message, true, opts)
}

func (c ConsolePrompter) Password(message string, opts ...Option) (string, error) {
	return c.prompt(message, true, opts)
}

func (c ConsolePrompter) prompt(message string, password bool, opts []Option) (string, error) {
	options, err := processOptions(opts)
	if err != nil {
		return "", err
	}

	prompter := console.Prompter{
		Message: message,
		Default: options.defaultChoice,
	}
	return prompter.Prompt(), nil
}

func (c ConsolePrompter) Choose(message string, choices []string, opts ...Option) (string, error) {
	options, err := processOptions(opts)
	if err != nil {
		return "", err
	}

	if options.defaultChoice != "" {
		def := strings.ToLower(options.defaultChoice)
		found := false
		for _, c := range choices {
			if def == strings.ToLower(c) {
				found = true
				break
			}
		}
		if !found {
			return "", fmt.Errorf("default choice not valid")
		}
	}

	prompter := console.Prompter{
		Message:    message,
		Choices:    choices,
		IgnoreCase: true,
		Default:    options.defaultChoice,
	}
	return prompter.Prompt(), nil
}

func Console() *ConsolePrompter {
	return new(ConsolePrompter)
}
