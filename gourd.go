// Gourd is my personal linux configuration tool.
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

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
		writeErr := utils.WriteJSON("test.txt", &osInfo)
		checkerr(writeErr)

		readErr := utils.ReadJSON("test.txt", &osInfo)
		checkerr(readErr)
		fmt.Printf("%+v\n", osInfo)

		path, err := exec.LookPath("ls")
		checkerr(err)

		success, err := utils.UserCanExec(path)
		checkerr(err)

		if success {
			log.Printf("$USER can exec %s", path)
		}

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
