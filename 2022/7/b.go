package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type File struct {
	name     string
	size     int
	fileType string
	parent   *File
	files    []*File
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Unable to open file")
	}

	root := &File{name: "/", fileType: "dir"}
	pwd := root

	s := bufio.NewScanner(f)

	var i int
	for s.Scan() {
		line := s.Text()
		// fmt.Printf("%d - %s\n", i+1, line)
		tokens := strings.Split(line, " ")
		if tokens[0] == "$" {
			switch tokens[1] {
			case "cd":
				pwd = changeDirectory(pwd, tokens[2])
				if pwd == nil {
					log.Fatal("WHAT")
				}
			}
		} else {
			if tokens[0] == "dir" {
				mkdir(pwd, tokens[1])
			} else {
				touch(pwd, tokens[1], tokens[0])
			}
		}
		i++
	}

	totalDiskSpace := 70000000
	totalNeeded := 30000000
	sums := make(map[string]int)
	var directorySizes []int
	totalSize := calculateDirectorySize(root, sums)
	fmt.Printf("Total Size: %d\n", totalSize)
	unused := totalDiskSpace - totalSize
	toDelete := totalNeeded - unused

	var answer int
	for _, v := range sums {
		directorySizes = append(directorySizes, v)
	}

	sort.Ints(directorySizes)
	for _, s := range directorySizes {
		if s >= toDelete {
			answer = s
			break
		}
	}

	fmt.Printf("%d directories; answer: %d\n", len(sums), answer)
}

func mkdir(pwd *File, name string) {
	f := &File{
		name:     name,
		fileType: "dir",
		size:     0,
		parent:   pwd,
		files:    []*File{},
	}
	pwd.files = append(pwd.files, f)
}

func touch(pwd *File, name, size string) {
	sizeAsInt, _ := strconv.Atoi(size)
	f := &File{
		name:     name,
		fileType: "file",
		size:     sizeAsInt,
		parent:   pwd,
	}
	pwd.files = append(pwd.files, f)
	pwd.size += sizeAsInt
}

func iterateTree(root *File, offset string) {
	var print string
	if root.fileType == "file" {
		print = fmt.Sprintf("%s- %s (%s, size=%d)", offset, root.name, root.fileType, root.size)
	} else {
		print = fmt.Sprintf("%s- %s (%s)", offset, root.name, root.fileType)
	}
	fmt.Println(print)

	for _, d := range root.files {
		iterateTree(d, offset+" ")
	}
}

func calculateDirectorySize(pwd *File, memo map[string]int) int {
	if pwd.fileType == "file" {
		if val, ok := memo[pwd.name]; ok {
			return val
		}
		return pwd.size
	}

	var pwdFileSize int
	for _, f := range pwd.files {
		pwdFileSize += calculateDirectorySize(f, memo)
		memo[pwd.name+string(pwd.files[0].size)] = pwdFileSize
	}
	return pwdFileSize
}

func changeDirectory(dir *File, changeTo string) *File {
	if changeTo == "/" && dir.name == "/" {
		return dir
	}

	if changeTo == ".." {
		return dir.parent
	}

	for _, d := range dir.files {
		if d.name == changeTo {
			return d
		}
	}

	return nil
}
