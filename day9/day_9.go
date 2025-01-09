package main

import (
	"fmt"
	"io"
	"os"
)

type block struct {
	id         int
	fileSpace  int
	freeSpace  int
	movedSpace int
	pos        int
	used       bool
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
	totalSpaces := 0
	for i := 0; i < len(byteDiskMap); i += 2 {
		pos := totalSpaces
		fileSpace := int(byteDiskMap[i] - '0')
		freeSpace := 0
		if i+1 < len(byteDiskMap) {
			freeSpace = int(byteDiskMap[i+1] - '0')
		}

		totalFileSpaces += fileSpace
		totalSpaces += fileSpace + freeSpace

		blocks = append(blocks, block{id: i / 2, fileSpace: fileSpace, freeSpace: freeSpace, pos: pos})
	}

	part1CheckSums := part1(blocks, totalFileSpaces)
	part2CheckSums := part2(blocks)

	fmt.Println(part1CheckSums)
	fmt.Println(part2CheckSums)
}

func part1(blocks []block, totalSpaces int) int {
	copiedBlocks := make([]block, len(blocks))
	copy(copiedBlocks, blocks)

	checkSum := 0
	left, right := 0, len(blocks)-1

	// Close in from left --> right
	// if curr left has no free spaces, move one more to the left
	// if curr right has no more file spaces, move one more to the right
	for iter := 0; iter < totalSpaces; {
		// If left block has file spaces available, then we haven't processed the check sum so far at that position/id
		if copiedBlocks[left].fileSpace > 0 {
			checkSum += iter * copiedBlocks[left].id
			copiedBlocks[left].fileSpace--
			iter++
			continue
		}

		// If we have free space on the left and file space on the right, move right file space to left free space + update
		if copiedBlocks[left].freeSpace > 0 && copiedBlocks[right].fileSpace > 0 {
			checkSum += iter * copiedBlocks[right].id
			copiedBlocks[left].freeSpace--
			copiedBlocks[right].fileSpace--
			iter++
			continue
		}

		// If we have no file space to compute checksum, and no free space on left to move any file space from right, move left one
		if copiedBlocks[left].fileSpace == 0 && copiedBlocks[left].freeSpace == 0 {
			left++
			continue
		}

		// If we have no more file space on right, then we've moved all possible files, move right one
		if copiedBlocks[right].fileSpace == 0 {
			right--
			continue
		}
	}

	return checkSum
}

func part2(blocks []block) int {
	copiedBlocks := make([]block, len(blocks))
	copy(copiedBlocks, blocks)

	// Move files starting from end of file system IDs, id = 0 doesn't move
	for i := len(copiedBlocks) - 1; i >= 1; i-- {
		fileBlock := copiedBlocks[i]
		// We can only move to the freespace of a block leftward of the current block we're trying to move
		for j := 0; j < i; j++ {
			freeBlock := copiedBlocks[j]
			// If it's ...
			// - identical blocks
			// - used/moved previously
			// - there isn't enough free space for the file
			// don't move
			if fileBlock == freeBlock || freeBlock.used || fileBlock.fileSpace > freeBlock.freeSpace {
				continue
			}

			// Set moved block to used so we don't move it again
			copiedBlocks[i].used = true
			// keep track of position of the file after movement
			// Take position of the block with free space moved into + the file space of that given block + any file space moved previously
			copiedBlocks[i].pos = copiedBlocks[j].pos + copiedBlocks[j].fileSpace + copiedBlocks[j].movedSpace

			// for the block that had open space, update movedSpace to account for added file
			// Remove the free space taken up by the file moved
			copiedBlocks[j].movedSpace += copiedBlocks[i].fileSpace
			copiedBlocks[j].freeSpace -= copiedBlocks[i].fileSpace
			break
		}
	}

	checkSum := 0
	for i := 0; i < len(copiedBlocks); i++ {
		currPos := copiedBlocks[i].pos
		for range copiedBlocks[i].fileSpace {
			checkSum += currPos * copiedBlocks[i].id
			currPos++
		}
	}

	return checkSum
}
