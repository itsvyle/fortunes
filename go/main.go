package main

import (
	"flag"
	"os"
)

var path = flag.String("path", "../test-fortunes", "Path to the folder containing fortunes and the `.vyle` file given by the fortune-generator")
var maxLength = flag.Int("max", 0, "Max length of the generated fortune; 0 = no-limit")
var minLength = flag.Int("min", 0, "Min length of the generated fortune")
var showSourceName = flag.Bool("s", false, "Show the source file name of the fortune")
var iterationsCount = flag.Int("n", 1, "Number of fortunes to generate")

type Fortune struct {
	id     byte
	weight byte
	name   string
}

func GiveFortune() {
	vyleFile := *path + "/fortunes.vyle"
	// get a handle to the file
	file, err := os.OpenFile(vyleFile, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	oneByte := make([]byte, 1)

	_, err = file.Read(oneByte)
	if err != nil {
		panic(err)
	}
	if oneByte[0] != 1 {
		panic("Invalid version byte")
	}

	_, err = file.Read(oneByte)
	if err != nil {
		panic(err)
	}
	fortunesCount := int(oneByte[0])

	fortunes := make([]Fortune, fortunesCount)
	for i := range fortunesCount {
		fortune := Fortune{}

		_, err = file.Read(oneByte)
		if err != nil {
			panic(err)
		}
		fortune.id = oneByte[0]

		// weight
		_, err = file.Read(oneByte)
		if err != nil {
			panic(err)
		}
		fortune.weight = oneByte[0]

		// name len
		_, err = file.Read(oneByte)
		if err != nil {
			panic(err)
		}

		name := make([]byte, oneByte[0])
		_, err = file.Read(name)
		if err != nil {
			panic(err)
		}
		fortune.name = string(name)
		fortunes[i] = fortune
	}
	// skip 10 empty bytes
	if _, err = file.Seek(10, 1); err != nil {
		panic(err)
	}

	entriesCount_ := make([]byte, 4)
	_, err = file.Read(entriesCount_)
	if err != nil {
		panic(err)
	}
	entriesCount := int(entriesCount_[0])<<24 | int(entriesCount_[1])<<16 | int(entriesCount_[2])<<8 | int(entriesCount_[3])
	print(entriesCount)
	print(fortunes)

}

func main() {
	if *minLength > *maxLength {
		panic("Min length cannot be greater than max length")
	}
	if *path == "" {
		panic("\"-path: Path to the folder containing fortunes and the `.vyle` file given by the fortune-generator\" is required")
	}
	for i := 0; i < *iterationsCount; i++ {
		GiveFortune()
	}
}
