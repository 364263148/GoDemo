package main

import(
	"os"
	"github.com/urfave/cli"

	"Cxl/cmd"
)

func main(){

	app := cli.NewApp()
	app.Name = "GoDemo App"
	app.Usage = "Just show some message"
	app.Version = "0.0.0.1"
	app.Commands = []cli.Command{
		cmd.CmdStart,
	}
	app.Run(os.Args)

}