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
	noReport bool
}


type Report struct {
	dirs int
	files int
}


func tree(root string, options *Options, report *Report, level int) {
	finfos, err := ioutil.ReadDir(root)

	if err != nil {
		fmt.Println("[error opening dir]") // TODO: output on same line as dir
		return
	}

	for fnum, finfo := range finfos {
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

			branch := "├── "
			if fnum == len(finfos) - 1  || options.dirOnly {
				branch = "└── "
			}

			for i:=0; i<level; i++ {
				ind = fmt.Sprintf("%s%s", ind, "│   ")
			}
			ind = fmt.Sprintf("%s%s", ind, branch)
		}

		/* Construct paths for recursion and printing */
		path := fmt.Sprintf("%s/%s", root, finfo.Name())
		var treepath string

		if options.listSz {
			szstr := fmt.Sprintf("[%v]", finfo.Size())
			treepath = fmt.Sprintf("%s  %s", szstr, treepath)
		}

		if options.fullPath {
			treepath = fmt.Sprintf("%s%s%s", ind, treepath, path)
		} else {
			treepath = fmt.Sprintf("%s%s%s", ind, treepath, finfo.Name())
		}

		fmt.Println(treepath)

		/* Recurse through subdirectories and count for report */
		if finfo.IsDir() {
			report.dirs++
			tree(path, options, report, level + 1)
		} else {
			report.files++
		}
	}
}


func main() {

	/* Parse Args */
	listAll  := flag.Bool("a", false, "All files are listed")
	dirOnly  := flag.Bool("d", false, "List directories only")
	fullPath := flag.Bool("f", false, "Print full path prefixes")
	noIndent := flag.Bool("i", false, "Do not print any indentation prefixes")
	listSz   := flag.Bool("s", false, "Print size of each file")
	// should be -- but that is a parse error in the flag pkg
	noReport := flag.Bool("noreport", false, "Omit printing of file and directory report at the end of tree listing")

	flag.Parse()

	options := Options { *listAll, *dirOnly, *fullPath, *noIndent, *listSz, *noReport }
	report := Report { 0, 0 }

	dirs := flag.Args()
	arglen := len(dirs)

	if arglen == 0 {
		/* Default to current directory */
		fmt.Println(".")
		tree(".", &options, &report, 0)

	} else {
		/* Allow user to supply list of paths */
		for _, root := range dirs {
			fmt.Println(root)
			tree(root, &options, &report, 0)
		}
	}

	if !options.noReport {
		if report.dirs > 0 {
			if report.files > 0 {
				fmt.Printf("\n%v directories, %v files\n", report.dirs, report.files)
			} else {
				fmt.Printf("\n%v directories\n", report.dirs)
			}
		} else if report.files > 0 {
			fmt.Printf("\n%v files\n", report.files)
		}
	}
}
