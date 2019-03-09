package prompt

import (
	"errors"
)

var (
	// ErrUserCancel is returned when the user cancels the prompt instead of answering it.
	ErrUserCancel = errors.New("operation canceled by the user")
	// ErrNotSupported is returned by the prompter when it doesn't support some feature.
	ErrNotSupported = errors.New("prompt not supported")
)

type options struct {
	defaultChoice string
}

// Option is a prompt option.
type Option func(opts *options) error

// Default sets the default for a prompt.
func Default(s string) Option {
	return func(opts *options) error {
		opts.defaultChoice = s
		return nil
	}
}

func processOptions(opts []Option) (*options, error) {
	cfg := new(options)
	for _, o := range opts {
		if err := o(cfg); err != nil {
			return nil, err
		}
	}
	return cfg, nil
}

// Prompter is the interface satisfied by prompt agents.
type Prompter interface {
	// Prompt prompts the user for input.
	Prompt(message string, opts ...Option) (answer string, err error)

	// Password prompts the user for password input.
	Password(message string, opts ...Option) (answer string, err error)

	// YesNo asks the user a yes/no question.
	YesNo(message string, opts ...Option) (answer bool, err error)

	// Choose asks the user to choose from a set of options.
	Choose(message string, options []string, opts ...Option) (choice int, err error)
}
