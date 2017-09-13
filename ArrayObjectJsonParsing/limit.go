package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Cmd struct {
	Command    string `json:"name,omitempty"`
	OutputFile string `json:"outputFile,omitempty"`
	Args       Args
}

type Args struct {
	Inputfolder  string `json:"inputfolder,omitempty"`
	Outputfolder string `json:"outputfolder,omitempty"`
	Orders       []string
}

type Mapping struct {
	Pool map[string]Cmd
}

func main() {
	mc := new(Mapping)
	b, err := ioutil.ReadFile("new.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b) // convert content to a 'string'
	if err = mc.FromJson(str); err != nil {
		fmt.Println(err)
	}

	fmt.Println(mc.Pool["csv2csv"].OutputFile)
}

func (mc *Mapping) FromJson(jsonStr string) error {
	var data = &mc.Pool
	b := []byte(jsonStr)
	return json.Unmarshal(b, data)
}
