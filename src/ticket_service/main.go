package main

import (
	"fmt"
	"os"
)

func main() {
	if err := InitConfig(); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		return
	}
	createEngine().Run("localhost:" + Conf.APIPort)
}
