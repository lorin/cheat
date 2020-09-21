package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func usage() {
	usage := `
usage: cheat <project>

`
	fmt.Printf(usage)
}

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	project := os.Args[1]
	dir := dirname(project)

	mkdir(dir)
	chdir(dir)

	makeGitIgnore()
	makeMakefile(project)
	makeRb(project)
	makeReadme(project)
}

func dirname(project string) string {
	return project + "-cheat-sheet"
}

func mkdir(project string) {
	err := os.Mkdir(project, 0755)
	check(err)
}

func chdir(project string) {
	err := os.Chdir(project)
	check(err)
}

// write templated text to a file
func write(text, project, fname string) {
	f, err := os.Create(fname)
	check(err)
	defer f.Close()
	tmpl, err := template.New(fname).Parse(text)
	check(err)

	err = tmpl.Execute(f, struct{ Project string }{project})
	check(err)
}

func makeReadme(project string) {
	text := `# {{.Project}} cheat sheet for Dash

This is a {{.Project}} [cheat sheet][1] for [Dash][2].

You need to have [cheatset][3] installed.

To install it on your local machine, do:

	make

[1]: https://github.com/Kapeli/cheatsheets
[2]: https://kapeli.com/dash
[3]: https://github.com/Kapeli/cheatset

`
	write(text, project, "README.md")
}

func makeMakefile(project string) {
	text := `
{{.Project}}.docset: {{.Project}}.rb
	cheatset generate $<
	open $@
`
	write(text, project, "Makefile")
}

func makeRb(project string) {
	text := `cheatsheet do
    title '{{.Project}}'
    docset_file_name '{{.Project}}'
    keyword '{{.Project}}'
    category do
        id ''
    end
end
`
	write(text, project, project+".rb")
}

func makeGitIgnore() {
	text := "*.docset\n"
	err := ioutil.WriteFile(".gitignore", []byte(text), 0755)
	check(err)
}
