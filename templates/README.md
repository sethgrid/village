# Templates

Files in this directory are HTML templates, using Go's [Template Package](https://golang.org/pkg/html/template/).

Files that start with an underscore (henceforth 'private templates'), like `_base.tmpl`, are not designed to be used directly (as opposed to a file with no leading underscore (henceforth 'public templates'), like `donate.tmpl`).

Public templates should be wired up in `main.go` and referenced as a `.html` file and not a `.tmpl` file.