package cmd

import "github.com/urfave/cli"

func init() {
	create := cli.Command{
		Name:    "create",
		Aliases: []string{"c"},
		Usage:   "create a new repo",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "public",
				Usage: "make the repo public",
			},
		},
		Category: "REPOS",
		Action:   create,
	}
	Cmds = append(Cmds, create)
}

func create(ctx *cli.Context) error {
	return nil
}
