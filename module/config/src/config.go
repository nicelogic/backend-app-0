package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Init(path string, out interface{}) (err error){
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return 
	}

	err = yaml.Unmarshal(content, out)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
