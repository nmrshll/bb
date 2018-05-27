package cmd

import (
	"github.com/n-marshall/bb/promptui"
	"github.com/n-marshall/bb/validate"
)

// PasswordMask is the character used to mask password inputs
const PasswordMask = '●'

// PasswordPrompt prompts the user to input a password value
func PasswordPrompt(shouldConfirm bool, labelOverride *string) (string, error) {
	label := "Password"
	if labelOverride != nil {
		label = *labelOverride
	}

	prompt := promptui.Prompt{
		Label: label,
		Mask:  PasswordMask,
		Validate: func(input string) error {
			err := validate.Password(input)
			if err == nil {
				return nil
			}

			return promptui.NewValidationError(err.Error())
		},
	}

	password, err := prompt.Run()
	if err != nil {
		return "", err
	}
	if !shouldConfirm {
		return password, err
	}

	prompt = promptui.Prompt{
		Label: "Confirm " + label,
		Mask:  '●',
		Validate: func(input string) error {
			if len(input) > 0 {
				if input != password {
					return promptui.NewValidationError("Passwords do not match")
				}
				return nil
			}

			return promptui.NewValidationError("Please confirm your password")
		},
	}

	_, err = prompt.Run()
	if err != nil {
		return "", err
	}

	return password, nil
}

// EmailPrompt prompts the user to input an email
func EmailPrompt(defaultValue string) (string, error) {
	prompt := promptui.Prompt{
		Label: "Email",
		Validate: func(input string) error {
			err := validate.Email(input)
			if err == nil {
				return nil
			}
			return promptui.NewValidationError("Please enter a valid email address")
		},
	}
	if defaultValue != "" {
		prompt.Default = defaultValue
	}

	return prompt.Run()
}

// UsernamePrompt prompts the user to input a person's name
func UsernamePrompt(un string) (string, error) {
	prompt := promptui.Prompt{
		Label: "Username",
		Validate: func(input string) error {
			err := validate.Username(input)
			if err == nil {
				return nil
			}
			return promptui.NewValidationError("Please enter a valid username")
		},
	}
	if un != "" {
		prompt.Default = un
	}

	return prompt.Run()
}
