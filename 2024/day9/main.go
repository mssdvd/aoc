package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	diskMap, err := os.ReadFile("./input")
	check(err)

	// Remove line feed
	diskMap = diskMap[:len(diskMap)-1]

	var disk1, disk2 []string
	id := 0

	for i, block := range diskMap {
		var c string
		if i%2 == 0 {
			// File blocks
			c = strconv.Itoa(id)
			id++
		} else {
			// Empty blocks
			c = "."
		}
		blockLen, err := strconv.Atoi(string(block))
		check(err)
		for range blockLen {
			disk1 = append(disk1, c)
		}
	}

	disk2 = append(disk2, disk1...)

	fileBlock := len(disk1) - 1
	emptyBlock := 0

	for range disk1 {
		for disk1[emptyBlock] != "." {
			emptyBlock++
		}

		for disk1[fileBlock] == "." {
			fileBlock--
		}

		if fileBlock <= emptyBlock {
			break
		}

		disk1[emptyBlock] = disk1[fileBlock]
		disk1[fileBlock] = "."

	}

	fileBlock = len(disk2) - 1
	emptyBlock = 0

	diskMapIndex := len(diskMap) - 1
	if diskMapIndex%2 != 0 {
		diskMapIndex--
	}

	for diskMapIndex >= 0 {
		fileSize, err := strconv.Atoi(string(diskMap[diskMapIndex]))
		check(err)

		for emptyBlock < len(disk2) && disk2[emptyBlock] != "." {
			emptyBlock++
		}

		currentFileID := strconv.Itoa(diskMapIndex / 2)
		for disk2[fileBlock] != currentFileID {
			fileBlock--
		}

		if fileBlock <= emptyBlock {
			emptyBlock = 0
			diskMapIndex -= 2
			continue
		}

		emptySize := 0
		for emptyBlock+emptySize < len(disk2) && disk2[emptyBlock+emptySize] == "." {
			emptySize++
		}

		if emptySize < fileSize {
			emptyBlock += emptySize
			continue
		}

		for range fileSize {
			disk2[emptyBlock] = disk2[fileBlock]
			disk2[fileBlock] = "."
			emptyBlock++
			fileBlock--
		}
		emptyBlock = 0
		diskMapIndex -= 2
	}

	checksum1 := checksum(disk1)
	fmt.Println("Part One:", checksum1)

	checksum2 := checksum(disk2)
	fmt.Println("Part Two:", checksum2)
}

func checksum(disk []string) int {
	var checksum int

	for i, block := range disk {
		if block == "." {
			continue
		}

		id, err := strconv.Atoi(block)
		check(err)
		checksum += i * id
	}

	return checksum
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
