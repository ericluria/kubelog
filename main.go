package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	pod, err := getPod()
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("kubectl", "logs", pod, "-f")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	stdout, _ := cmd.StdoutPipe()
	cmd.Run()

	buff := make([]byte, 1024)
	for {
		_, err := stdout.Read(buff)
		if err != nil {
			break
		}
		fmt.Println(string(buff))
		buff = make([]byte, 1024)
	}
}

func getPod() (string, error) {
	cmd := exec.Command("fzf", "--ansi", "--no-preview")
	var out bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = &out

	cmd.Env = append(os.Environ(),
		fmt.Sprintf("FZF_DEFAULT_COMMAND=%s", "kubectl get pods --no-headers -o custom-columns=':metadata.name'"),
	)
	if err := cmd.Run(); err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			return "", err
		}
	}
	choice := strings.TrimSpace(out.String())
	return choice, nil
}
