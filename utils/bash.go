package utils

import (
	"os"
	"os/exec"
)

func ClearBash() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}