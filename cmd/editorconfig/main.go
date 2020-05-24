package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	version = "dev"
	commit  = ""
	date    = ""
)

func main() {
	flagVersion := flag.Bool("v", false, "print the version and exit")
	flag.Usage = func() {
		fmt.Println("Usage of editorconfig:")
		fmt.Println("Flags:")
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()
	if *flagVersion {
		fmt.Printf("editorconfig %s built at %s commit %s\n", version, date, commit)
		os.Exit(0)
	}

	source := []byte(`# EditorConfig helps developers define and maintain consistent
# coding styles between different editors and IDEs
# http://editorconfig.org

root = true

[*]
charset = utf-8
indent_style = space
indent_size = 2
end_of_line = lf
max_line_length = 120
trim_trailing_whitespace = true
insert_final_newline = true

[*.bat]
end_of_line = crlf

[{Makefile,*.mk}]
indent_size = 4
indent_style = tab

[{*.txt,*.md,*.mkdn,*.mdown,*.markdown}]
max_line_length = off
trim_trailing_whitespace = false
`)

	flagArgs := flag.Args()

	filename := ".editorconfig"
	if len(flagArgs) > 0 {
		filename = flagArgs[0]
	}
	err := ioutil.WriteFile(filename, source, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
