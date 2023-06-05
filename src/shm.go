package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func main() {

	f, err := os.OpenFile("trace.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Printf("--------------")

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	//var stdin bytes.Buffer

	cmd := exec.Command("bash")
	cmd.Stdin = &FakeReader{} //os.Stdin //strings.NewReader("ls -l\necho $$\n")
	cmd.Stdout = &FakeWriter{name: "out"}
	cmd.Stderr = &FakeWriter{name: "err"}
	errsh := cmd.Run()
	fmt.Printf("Err %+v\n", errsh)
	fmt.Printf("Stdout: %s", stdout.String())
	fmt.Printf("Stderr: %s", stderr.String())

}

func cli_cli() {
	(&cli.App{}).Run(os.Args)

}

type FakeReader struct {
	xxx int
}

func (f *FakeReader) Read(p []byte) (n int, err error) {
	buf := make([]byte, 1)

	i_n, i_err := os.Stdin.Read(buf)
	copy(p, buf)
	log.Printf("I: %s\n", string(p[:]))
	return i_n, i_err
}

type FakeWriter struct {
	name string
}

func (f *FakeWriter) Write(p []byte) (n int, err error) {
	log.Printf("%s: %s\n", f.name, string(p[:]))
	return len(p), nil
}
