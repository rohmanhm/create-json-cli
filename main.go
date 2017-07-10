/**
 * This is just WIP example implementation Create-json-cli with Golang
 * For a while don't use this package
 */

package main

import (
	"fmt"
	"os"
	"strings"

	"reflect"

	"github.com/Jeffail/gabs"
	"github.com/urfave/cli"
)

// RecursiveValue is value recursive
type RecursiveValue []string

func main() {
	program := cli.NewApp()
	program.Name = "create-json-cli"
	program.Version = "0.0.1"
	program.Usage = "CLI Help Create .babelrc file"
	program.Action = func(cmd *cli.Context) error {
		argsLen := len(cmd.Args())
		if argsLen > 0 {
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(pwd)

			cmdValue := RecursiveValue(cmd.Args())
			result := cmdValue.Serialize()

			jsonString := result.StringIndent("", "  ")
			fmt.Println(string(jsonString))
		}
		return nil
	}

	program.Run(os.Args)
}

// Serialize for Recursive value
func (rv RecursiveValue) Serialize() gabs.Container {
	result := gabs.New()
	for _, i := range rv {
		s := strings.Split(i, "=")
		key, value := s[0], s[1]

		svalue := strings.Split(value, ",")

		if len(svalue) > 1 {
			// rvalue := RecursiveValue(svalue)
			// buf := &bytes.Buffer{}
			// gob.NewEncoder(buf).Encode(value)
			// bs := buf.Bytes()
			// jsonParsed, _ := gabs.ParseJSON([]byte(bs))
			// var aa []string
			// some := json.Unmarshal(bs, &aa)
			fmt.Println(reflect.TypeOf(value))
			// result.Set(rvalue, key)
		} else {
			result.Set(value, key)
		}
	}

	return *result
}
