package util

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func Cmder(command string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", command)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		//return simple error msg
		s, _ := stderr.ReadString('\n')
		return "", fmt.Errorf("%s. %v", strings.TrimSpace(s), err)
	}
	s, _ := stderr.ReadString('\n')
	if s == "" {
		return string(stdout.Bytes()), nil
	}
	return string(stdout.Bytes()), fmt.Errorf("%s", strings.TrimSpace(s))
}
