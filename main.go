package main

import (
	"os"
	"sort"

	"net/http"

	"github.com/dghubble/sling"
	"github.com/n-marshall/bb/cmd"
	"github.com/urfave/cli"
)

const (
	bitbucketAPI = "https://api.bitbucket.org/1.0/"
)

var (
	apiBase = sling.New().Base(bitbucketAPI).Client(http.DefaultClient)
)

// func Exec(cmd []string) (string, error) {
// 	spew.Dump(strings.Join(cmd, " "))
// 	output, err := exec.Command(cmd[0], cmd[1:]...).Output()
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(output), nil
// }

func main() {
	app := cli.NewApp()
	app.Name = "bb"
	app.Usage = "Bitbucket CLI interface"
	app.Version = "0.1.0"
	app.Commands = cmd.Cmds

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}
