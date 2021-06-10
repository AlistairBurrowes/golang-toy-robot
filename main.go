package main

import (
	"bufio"
	"fmt"
	"golang-toy-robot/model"
	"golang-toy-robot/processing"
	"os"
)

func main() {
	table := model.Table{MaxCoordinate: model.Coordinate{5, 5}, Robot: nil}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("~~ Toy Robot ~~")

	for {
		scanner.Scan()
		readCommandAsText := scanner.Text()

		toRunCommand := processing.ParseCommand(readCommandAsText)

		if toRunCommand != nil {
			table = processing.ProcessCommand(table, toRunCommand)
		}
	}
}
