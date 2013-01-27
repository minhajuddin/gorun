package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"syscall"
)

const PRETEMPLATE = `
package main

import (
	"fmt"
)

func main(){
	`
const POSTTEMPLATE = `
}
`

//Give a function p which maps to fmt.Printf or fmt.Println 
//based on the first input
func main() {
	//read from stdin and inspect it for functions and imports
	//write the pretemplate to tmp file
	//write the stdin code
	//write the posttemplate
	//rename the file to *.go
	//exec using go run
	f, err := ioutil.TempFile("", "goeval-stub")
	handle(err)
	tempfilepath := f.Name()
	f.WriteString(PRETEMPLATE)
	io.Copy(f, os.Stdin)
	f.WriteString(POSTTEMPLATE)
	f.Close()
	newfilepath := tempfilepath + ".go"
	os.Rename(tempfilepath, newfilepath)
	runFile(newfilepath)
}

func runFile(fp string) {
	binary, err := exec.LookPath("go")
	handle(err)
	args := []string{"", "run", fp}
	err = syscall.Exec(binary, args, os.Environ())
	handle(err)
}

func handle(err error) {
	if err != nil {
		log.Fatal("ERR: ", err)
	}
}
