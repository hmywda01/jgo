package web

import (
	"io"
	"text/template"
	"errors"
	"path/filepath"
	"strings"
	"io/ioutil"
	"os"
)

type Renderer struct {
	funcMap      template.FuncMap
	templateText string
}

func (r *Renderer) getTemplate() *template.Template {
	return template.Must(template.New("_base").Funcs(r.funcMap).Parse(r.templateText))
}

func (r *Renderer) SetFuncMap(funcMap map[string]interface{}) {
	r.funcMap = template.FuncMap(funcMap)
}

func (r *Renderer) Render(names []string, writer io.Writer, data interface{}) error {
	t := r.getTemplate()
	for _, name := range names {
		if t.Lookup(name) != nil {
			return t.ExecuteTemplate(writer, name, data)
		}
	}
	return errors.New("Unable to find template.")
}

func GetRenderer(directory string) (*Renderer, error) {
	fileList := []string{}
	err := filepath.Walk(directory, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})
	check(err)

	templateText := ""
	for _, file := range fileList {
		if !strings.HasSuffix(file, ".html") {
			continue
		}

		contents, err := ioutil.ReadFile(file)
		check(err)

		filename := file[len(directory) + 1:]
		templateText += "{{ define \"" + filename + "\" }}" + string(contents) + "{{ end }}\n"
	}

	return &Renderer{
		templateText: templateText,
	}, nil
}
