// Gourd is my personal linux configuration tool.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/zahfox/gourd/utils"
)

func main() {
	app := cli.NewApp()
	app.Name = "gourd"
	app.Usage = "linux configuration tool"

	app.Action = func(c *cli.Context) error {
		osInfo, err := utils.Os()
		checkerr(err)
		writeErr := utils.Write("test.txt", &osInfo)
		checkerr(writeErr)

		readErr := utils.Read("test.txt", &osInfo)
		checkerr(readErr)
		fmt.Printf("%+v\n", osInfo)
		return nil
	}

	err := app.Run(os.Args)
	checkerr(err)
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
