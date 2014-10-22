package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

type Network struct {
	Host string
	Ip   string
}

type JDK struct {
	Version []string
	Path    []string
}

type DataSheet struct {
	Description string
	Environment string
	OS          string
	Network     Network
	Jdk         JDK
}

func main() {

	rf, err := ioutil.ReadFile("./datasheet.json")
	if err != nil {
		log.Fatal(err)
	}

	var params DataSheet
	json.Unmarshal(rf, &params)
	// fmt.Printf("%s \n", params.Description)
	// fmt.Printf("%s \n", params.Environment)
	// fmt.Printf("%s \n", params.OS)
	// fmt.Printf("%s \n", params.Network.Host)
	// fmt.Printf("%s \n", params.Network.Ip)
	// fmt.Printf("%s \n", params.Jdk.Version[0])
	// fmt.Printf("%s \n", params.Jdk.Path[0])

	parse(params)
}

func parse(data DataSheet) {

	if data.OS != "" {
		checkOS(data.OS)
	}

	// if (data.Network != Network{}) {
	// 	fmt.Printf("check %s \n", data.Network.Host)
	// 	fmt.Printf("check %s \n", data.Network.Ip)
	// }
}

func run(command string, arg string) string {
	cmd := exec.Command(command, arg)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	result, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}

	return string(result)
}

func equals(name string, actual string, expected string) {
	if strings.Contains(actual, expected) {
		fmt.Printf("'%s' is '%s' \n", name, expected)
	} else {
		fmt.Printf("'%s' expected '%s', got '%s'\n", name, expected, actual)
	}
}

func checkOS(expected string) {
	name := "OS"
	actual := strings.TrimSpace(run("cat", "/etc/redhat-release"))
	equals(name, actual, expected)
}
