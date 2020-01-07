/* tree.go

	an implementation of the unix commmand tree
	https://linux.die.net/man/1/tree

	Gavin Inglis (ginglis)

*/


package main


import (
	"fmt"
	"io/ioutil"
	"flag"
)


type Options struct {
	listAll bool
	dirOnly bool
	fullPath bool
	noIndent bool
	listSz bool
}


func tree(root string, options Options, level int) {
	finfos, err := ioutil.ReadDir(root)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, finfo := range finfos {
		/* Check options */
		if finfo.Name() == "." || finfo.Name() == ".." {
			continue
		}

		if finfo.Name()[0] == '.' && !options.listAll {
			continue
		}

		if !finfo.IsDir() && options.dirOnly {
			continue
		}

		ind := ""
		if !options.noIndent {
			branch := "|---"
			for i:=0; i<level; i++ {
				ind = fmt.Sprintf("%s%s", ind, "    ")
			}
			ind = fmt.Sprintf("%s%s", ind, branch)
		}

		/* Constructs path for recursion and printing */
		path := fmt.Sprintf("%s/%s", root, finfo.Name())
		treepath  := fmt.Sprintf("%s%s", ind, finfo.Name())

		fmt.Println(treepath)

		/* Recurse through subdirectories */
		if finfo.IsDir() {
			tree(path, options, level + 1)
		}
	}
}


func main() {

	/* Parse Args */
	listAll  := flag.Bool("a", false, "All files are listed")
	dirOnly  := flag.Bool("d", false, "[TODO] List directories only")
	fullPath := flag.Bool("f", false, "[TODO] Print full path prefixes")
	noIndent := flag.Bool("i", false, "Do not print any indentation prefixes")
	listSz   := flag.Bool("s", false, "[TODO] Print size of each file")

	flag.Parse()

	options := Options { *listAll, *dirOnly, *fullPath, *noIndent, *listSz }

	dirs := flag.Args()
	arglen := len(dirs)

	if arglen == 0 {
		/* Default to current directory */
		fmt.Println(".")
		tree(".", options, 0)

	} else {
		/* Allow user to supply list of paths */
		for _, root := range dirs {
			fmt.Println(root)
			tree(root, options, 0)
		}
	}
}
