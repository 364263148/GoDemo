package cmd

import(
	"fmt"
	"net/http"
	"github.com/urfave/cli"
	"time"
	"sync"
	"Cxl/route"
)

var CmdStart = cli.Command{
	Name:        "start",
	Usage:       "start web server",
	Description: "start web server",
	Action:      start,
}

func start(c *cli.Context) error {
	fmt.Println("Start web server!")
	serv := http.Server{
		Addr:         fmt.Sprintf("%s:%d", "localhost", 8888),
		Handler:      route.NewRouter(),
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
	}

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		serv.ListenAndServe()
	}()

	wg.Wait()
	return nil;
}
