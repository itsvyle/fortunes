package main

import (
	"flag"
	"math/rand"
	"os"
)

var path = flag.String("path", "", "Path to the folder containing fortunes and the `.vyle` file given by the fortune-generator")
var showSourceName = flag.Bool("s", false, "Show the source file name of the fortune")
var iterationsCount = flag.Int("n", 1, "Number of fortunes to generate")
var maxLength = flag.Int("max", 0, "Max length of the generated fortune; 0 = no-limit")
var minLength = flag.Int("min", 0, "Min length of the generated fortune")

type FortuneFile struct {
	id     byte
	weight byte
	name   string
}

var triesCount = 0

func GiveFortune(file *os.File) {
	if triesCount > 10 {
		println("Could not find a fortune with the given constraints")
		return
	}
	_, err := file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

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

	fortuneFiles := make([]FortuneFile, fortunesCount)
	for i := range fortunesCount {
		fortuneFile := FortuneFile{}

		_, err = file.Read(oneByte)
		if err != nil {
			panic(err)
		}
		fortuneFile.id = oneByte[0]

		// weight
		_, err = file.Read(oneByte)
		if err != nil {
			panic(err)
		}
		fortuneFile.weight = oneByte[0]

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
		fortuneFile.name = string(name)
		fortuneFiles[i] = fortuneFile
	}
	// skip 10 empty bytes
	if _, err = file.Seek(10, 1); err != nil {
		panic(err)
	}

	_, err = file.Read(fourBytes)
	if err != nil {
		panic(err)
	}
	entriesCount := readInt32(fourBytes)

	randomIndex := rand.Intn(entriesCount + 1)

	if randomIndex > 0 {
		_, err = file.Seek(int64(10*randomIndex), 1)
		if err != nil {
			panic(err)
		}
	}

	_, err = file.Read(oneByte)
	if err != nil {
		panic(err)
	}
	fortuneFileInfo := fortuneFiles[oneByte[0]]

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

	if *maxLength > 0 && fortuneLength > *maxLength {
		triesCount++
		GiveFortune(file)
		return
	} else if *minLength > 0 && fortuneLength < *minLength {
		triesCount++
		GiveFortune(file)
		return
	}

	fortuneFilePath := *path + "/" + fortuneFileInfo.name
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

	if *showSourceName {
		println("From: " + fortuneFileInfo.name)
	}
	print(string(fortuneContent))
}

func readInt32(bytes []byte) int {
	return int(bytes[0])<<24 | int(bytes[1])<<16 | int(bytes[2])<<8 | int(bytes[3])
}

func main() {
	flag.Parse()

	if *path == "" {
		panic("\"-path: Path to the folder containing fortunes and the `.vyle` file given by the fortune-generator\" is required")
	}

	vyleFile := *path + "/fortunes.vyle"
	// get a handle to the file
	file, err := os.OpenFile(vyleFile, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i := 0; i < *iterationsCount; i++ {
		GiveFortune(file)
	}
}
