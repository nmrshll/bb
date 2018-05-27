// Package cmd contains all of the bitbucket cli commands
package cmd

import "github.com/urfave/cli"

// Cmds is the list of all cli commands
var Cmds []cli.Command

// chain allows easy sequential calling of BeforeFuncs and AfterFuncs.
// chain will exit on the first error seen.
func chain(funcs ...func(*cli.Context) error) func(*cli.Context) error {
	return func(ctx *cli.Context) error {

		for _, f := range funcs {
			err := f(ctx)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
