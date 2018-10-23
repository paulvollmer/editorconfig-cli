package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const version = "0.1.0"

func main() {
	flagVersion := flag.Bool("v", false, "print the version and exit")
	flag.Parse()
	if *flagVersion {
		fmt.Printf("v%s\n", version)
		os.Exit(0)
	}

	source := []byte(`# EditorConfig helps developers define and maintain consistent
# coding styles between different editors and IDEs
# http://editorconfig.org

root = true

[*]
indent_style = space
indent_size = 2
end_of_line = lf
charset = utf-8
max_line_length = 120
trim_trailing_whitespace = true
insert_final_newline = true

[{Makefile,**.mk}]
indent_style = tab

[{*.txt,*.md,*.mkdn,*.mdown,*.markdown}]
max_line_length = 0
trim_trailing_whitespace = false
`)
	ioutil.WriteFile(".editorconfig", source, 0666)
}
