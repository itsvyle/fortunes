package main

import (
	"flag"
	"math/rand"
	"os"
)

var path = flag.String("path", "../test-fortunes", "Path to the folder containing fortunes and the `.vyle` file given by the fortune-generator")
var showSourceName = flag.Bool("s", true, "Show the source file name of the fortune")
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
	fourBytes := make([]byte, 4)

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

	fourBytes = make([]byte, 4)
	_, err = file.Read(fourBytes)
	if err != nil {
		panic(err)
	}
	entriesCount := readInt32(fourBytes)

	randomIndex := rand.Intn(entriesCount + 1)
	randomIndex = 0

	_, err = file.Seek(int64(11*randomIndex), 1)
	if err != nil {
		panic(err)
	}

	_, err = file.Read(oneByte)
	if err != nil {
		panic(err)
	}
	fortune := fortunes[oneByte[0]]

	if *showSourceName {
		println("From: " + fortune.name + "\n")
	}

	_, err = file.Read(fourBytes)
	if err != nil {
		panic(err)
	}
	fortuneOffset := readInt32(fourBytes)

	_, err = file.Read(fourBytes)
	if err != nil {
		panic(err)
	}
	fortuneLength := readInt32(fourBytes)

	fortuneFilePath := *path + "/" + fortune.name
	fortuneFile, err := os.OpenFile(fortuneFilePath, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer fortuneFile.Close()

	_, err = fortuneFile.Seek(int64(fortuneOffset), 0)
	if err != nil {
		panic(err)
	}

	fortuneContent := make([]byte, fortuneLength)
	_, err = fortuneFile.Read(fortuneContent)
	if err != nil {
		panic(err)
	}

	println(string(fortuneContent))
}

func readInt32(bytes []byte) int {
	return int(bytes[0])<<24 | int(bytes[1])<<16 | int(bytes[2])<<8 | int(bytes[3])
}

func main() {
	if *path == "" {
		panic("\"-path: Path to the folder containing fortunes and the `.vyle` file given by the fortune-generator\" is required")
	}
	for i := 0; i < *iterationsCount; i++ {
		GiveFortune()
	}
}
