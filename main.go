package main

import(
	"fmt"
	"os"
	"github.com/urfave/cli"
)

func show(c *cli.Context) error {
	fmt.Println("show function!")
	return nil;
}

func main(){
	fmt.Println("GoDemo!");

	app := cli.NewApp()
	app.Name = "GoDemo App"
	app.Usage = "Just show some message"
	app.Version = "0.0.0.1"
	app.Commands = []cli.Command{{
		Name:        "show",
		Usage:       "Just show some message",
		Description: "Just show some message for v0.0.0.1",
		Action:      show},
	}
	app.Run(os.Args)

}