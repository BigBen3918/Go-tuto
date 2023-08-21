// Go offers built-in support for creating dynamic content or showing customized output to the user with the text/template package. A sibling package named html/template provides the same API but has additional security features and should be used for generating HTML.

package main

import (
    "os"
    "text/template"
)

func main() {

    t1 := template.New("t1")
    t1, err := t1.Parse("Value is {{.}}\n")
    if err != nil {
        panic(err)
    }

    t1 = template.Must(t1.Parse("Value: {{.}}\n"))

    t1.Execute(os.Stdout, "some text")
    t1.Execute(os.Stdout, 5)
    t1.Execute(os.Stdout, []string{
        "Go",
        "Rust",
        "C++",
        "C#",
    })

    Create := func(name, t string) *template.Template {
        return template.Must(template.New(name).Parse(t))
    }

    t2 := Create("t2", "Name: {{.Name}}\n")

    t2.Execute(os.Stdout, struct {
        Name string
    }{"Jane Doe"})

    t2.Execute(os.Stdout, map[string]string{
        "Name": "Mickey Mouse",
    })

    t3 := Create("t3",
        "{{if . -}} yes {{else -}} no {{end}}\n")
    t3.Execute(os.Stdout, "not empty")
    t3.Execute(os.Stdout, "")

    t4 := Create("t4",
        "Range: {{range .}}{{.}} {{end}}\n")
    t4.Execute(os.Stdout,
        []string{
            "Go",
            "Rust",
            "C++",
            "C#",
        })
}