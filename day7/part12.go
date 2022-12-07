package main

import (
	. "aoc-2022/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// assumes commands are simulating DFS traversal!
type shell struct {
	input []string
	idx   int
}

func (sh *shell) up() {
	sh.idx++
}

func (sh *shell) readDir() []string {
	sh.idx += 2 // cd and ls
	var entries []string
	for ; sh.idx < len(sh.input) && sh.input[sh.idx][0] != '$'; sh.idx++ {
		entries = append(entries, sh.input[sh.idx])
	}
	return entries
}

type inode struct {
	size     int
	name     string
	children []inode
}

// create FS tree by DFS traversal. Calculate sizes while at it.
func (node *inode) createDir(name string, sh *shell) int {
	node.name = name
	entries := sh.readDir()
	for _, entry := range entries {
		f := strings.Split(entry, " ")
		if f[0] == "dir" {
			child := inode{}
			node.size += child.createDir(f[1], sh)
			node.children = append(node.children, child)
		} else {
			filesize, _ := strconv.Atoi(f[0])
			node.size += filesize
			child := inode{filesize, f[1], nil}
			node.children = append(node.children, child)
		}
	}
	sh.up()
	return node.size
}

func (node *inode) findLarge() int {
	if node.children == nil {
		return 0
	}
	var size = 0
	if node.size <= 100000 {
		size += node.size
	}
	for _, child := range node.children {
		size += child.findLarge()
	}
	return size
}

func (node *inode) findToDelete(deletionSize int, current int) int {
	if node.children == nil { // skip files
		return current
	}
	if node.size >= deletionSize && node.size < current {
		current = node.size
	}
	for _, child := range node.children {
		current = child.findToDelete(deletionSize, current)
	}
	return current
}

func main() {
	file, e := os.Open("input")
	PanicOnError(e)

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sh := &shell{lines, 0}
	root := inode{}
	root.createDir("/", sh)
	fmt.Println("part1", root.findLarge())
	free := 70000000 - root.size
	toDelete := 30000000 - free
	fmt.Println("part2", root.findToDelete(toDelete, root.size))
}
