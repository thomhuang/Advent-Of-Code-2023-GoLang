package main

import (
	"fmt"
	"io"
	"os"
)

type block struct {
	id        int
	fileSpace int
	freeSpace int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	byteDiskMap, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var blocks []block
	totalFileSpaces := 0
	for i := 0; i < len(byteDiskMap); i += 2 {
		fileSpace := int(byteDiskMap[i] - '0')
		freeSpace := 0
		if i+1 < len(byteDiskMap) {
			freeSpace = int(byteDiskMap[i+1] - '0')
		}

		totalFileSpaces += fileSpace

		blocks = append(blocks, block{id: i / 2, fileSpace: fileSpace, freeSpace: freeSpace})
	}

	checkSum := 0
	left, right := 0, len(blocks)-1
	for iter := 0; iter < totalFileSpaces; {

		if blocks[left].fileSpace > 0 {
			checkSum += iter * blocks[left].id
			blocks[left].fileSpace--
			iter++
			continue
		}

		if blocks[left].freeSpace > 0 && blocks[right].fileSpace > 0 {
			checkSum += iter * blocks[right].id
			blocks[left].freeSpace--
			blocks[right].fileSpace--
			iter++
			continue
		}

		if blocks[left].fileSpace == 0 && blocks[left].freeSpace == 0 {
			left++
			continue
		}

		if blocks[right].fileSpace == 0 {
			right--
			continue
		}
	}

	fmt.Println(checkSum)
}
