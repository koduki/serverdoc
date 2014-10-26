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
	Host [2]string
	Ip   [3]string
}

type JDK struct {
	Version [2]string
	Path    [2]string
}

type DataSheet struct {
	Description string
	Environment string
	OSName      [2]string
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

	if data.OSName[0] != "" {
		assertOSName(data.OSName[0])
	}

	if (data.Network != Network{}) {
		assertNetwork(data.Network)
	}
}

func run(command string, args ...string) string {
	var cmd *exec.Cmd
	cmd = exec.Command(command, args...)

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

	return strings.TrimSpace(string(result))
}

func equals(name string, actual string, expected string) {
	if actual == expected {
		fmt.Printf("CHECKED: '%s' is '%s' \n", name, expected)
	} else {
		fmt.Printf("FAILED: '%s' expected '%s', got '%s'\n", name, expected, actual)
	}
}

func contains(name string, actual string, expected string) {
	if strings.Contains(actual, expected) {
		fmt.Printf("CHECKED: '%s' contains '%s' \n", name, expected)
	} else {
		fmt.Printf("FAILED: '%s' expected '%s', got '%s'\n", name, expected, actual)
	}
}

func assertOSName(expected string) {
	actual := run("uname", "-a")
	contains("OS Name", actual, expected)
}

func assertNetwork(expected Network) {
	actual_host := run("hostname")
	actual_ip := run("ip", "a", "show", "dev", expected.Ip[0])

	equals("Netowrk.Host", actual_host, expected.Host[0])
	contains("Netowrk.Ip", actual_ip, expected.Ip[1])
}
