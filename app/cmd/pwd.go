package cmd

import (
	"fmt"
	"os"
)

func HandlePwd(command string, args ...string) {
	wd, _ := os.Getwd()
	fmt.Println(wd)
}
