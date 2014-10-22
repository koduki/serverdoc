package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"log"
	"os/exec"
  "io/ioutil"
//  "github.com/bitly/go-simplejson"
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


	// rf, err := ioutil.ReadFile("./params.json") //[]byte型での読み込み
	// if err != nil {
	//    log.Fatal(err)
	// }
	// json, err := simplejson.NewJson(rf)
	// desc, err := json.Get("Description").String()
	//
	// fmt.Printf(desc)


	type Person struct {
    Name  string
    Age   int
		Height int
		Weight int
	}

	rf, err := ioutil.ReadFile("./a.json")
	var person Person
	json.Unmarshal(rf, &person)

  fmt.Println(result)
	fmt.Printf("%s is %d years old, %d kg \n", person.Name, person.Age, person.Weight)

	type Params struct {
		Description string
	}

	rf2, err := ioutil.ReadFile("./datasheet.json")
	var params Params
	json.Unmarshal(rf2, &params)
	fmt.Printf("%s \n", params.Description)
}
