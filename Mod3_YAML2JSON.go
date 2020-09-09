package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/ghodss/yaml"
)

func PrintFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func prettyPrint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "	")
	return out.Bytes(), err
}

func main() {

	filePath := "E:\\Go Tasks\\Final\\Test2\\examples\\"
	JSONOutPath := "E:\\Go Tasks\\Final\\Test2\\JSONOut\\"

	files, err := ioutil.ReadDir(filePath)
	PrintFatalError(err)
	for _, file := range files {

		dat, err := ioutil.ReadFile(filePath + file.Name())
		PrintFatalError(err)
		fmt.Println(file.Name())

		m := make(map[string]interface{})

		err = yaml.Unmarshal([]byte(dat), &m)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		jsonOut, err := json.Marshal(m)
		PrintFatalError(err)

		pretty, _ := prettyPrint(jsonOut)

		err = ioutil.WriteFile(JSONOutPath+file.Name()+".json", pretty, 0644)
		time.Sleep(150)
		PrintFatalError(err)
	}

}
