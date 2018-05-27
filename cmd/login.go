package cmd

import "github.com/urfave/cli"
import "github.com/lunux2008/xulu"

func init() {
	login := cli.Command{
		Name:     "login",
		Usage:    "Log in to a user account, authenticating the CLI",
		Category: "ACCOUNT",
		Action:   login,
	}
	Cmds = append(Cmds, login)
}

func login(ctx *cli.Context) error {
	email, err := EmailPrompt("")
	if err != nil {
		return err
	}

	password, err := PasswordPrompt(false, nil)
	if err != nil {
		return err
	}

	xulu.Use(email, password)
	return nil
}
