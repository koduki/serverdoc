package main

import (
//	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"log"
	"os/exec"
  "io/ioutil"
  //"io"
)

func main() {
	cmd := exec.Command("ls")
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

  s, err := ioutil.ReadAll(stdout)
  fmt.Printf("%s", s)

	re, _ := regexp.Compile("main.go")
	result := re.Find(s)

  if string(result) == "" {
		error := errors.New("Not found")
    log.Fatal(error)
	}


  fmt.Printf("%s", result)
	//fmt.Printf("%s is %d years old\n", person.Name, person.Age)
}
