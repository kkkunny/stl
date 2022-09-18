package os

import (
	"os"
	"os/exec"
)

// Exec 执行系统命令
func Exec(cmd string, args ...string)error{
	command := exec.Command(cmd, args...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}
