/* tree.go

	an implementation of the unix commmand tree
	https://linux.die.net/man/1/tree

	Gavin Inglis (ginglis)

*/


/* TODO

	1. Flags - some I'm gonna do:
		--help
		-a 			All files are listed
		-d 			List directories only
		-f          Print full path prefixes
		-i          Do not print any indendation prefixes
		-s          Print size of each file

	There's a bunch more, this is it for now

	2. Final report (X directories, X files)

*/

package main


import (
	"fmt"
	"io/ioutil"
	"os"
)


func tree(root string, level int) {
	finfos, err := ioutil.ReadDir(root)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, finfo := range finfos {
		if finfo.Name() == "." || finfo.Name() == ".." {
			continue
		}

		ind := ""
		for i:=0; i<level; i++ {
			ind = fmt.Sprintf("%s%s", ind, "    ")
		}

		path := fmt.Sprintf("%s/%s", root, finfo.Name())
		treepath  := fmt.Sprintf("%s|---%s", ind, finfo.Name())

		fmt.Println(treepath)

		if finfo.IsDir() {
			tree(path, level + 1)
		}
	}
}


func main() {
	root := os.Args[0]
	fmt.Println(root)
	tree(root, 0)
}
