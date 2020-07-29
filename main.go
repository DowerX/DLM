package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Downloads  string `yaml:"downloads"`
	Conditions []Condition
}

type Condition struct {
	Extension string `yaml:"ext"`
	Path      string `yaml:"path"`
}

var c Config

func main() {
	c = Config{}

	fmt.Println("Loading config from " + os.Args[1][4:] + ".")
	time.Sleep(1 * time.Second)

	data, err := ioutil.ReadFile(os.Args[1][4:])
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}

	items, err := ioutil.ReadDir(c.Downloads)
	if err != nil {
		panic(err)
	}

	for _, item := range items {
		if !item.IsDir() {
			for _, v := range c.Conditions {
				parts := strings.Split(strings.ToLower(item.Name()), ".")
				if v.Extension == parts[len(parts)-1] {
					err = os.Rename(c.Downloads+item.Name(), v.Path+item.Name())
					if err != nil {
						panic(err)
					}
					fmt.Println("Moved " + item.Name() + " to " + v.Path + item.Name())
				}
			}
		}
	}
}
