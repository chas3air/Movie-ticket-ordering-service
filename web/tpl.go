package web

import (
	"html/template"
	"path/filepath"
)

func GetTPL() *template.Template {
	// Строим путь к папке шаблонов относительно исполняемого файла
	tplPath := filepath.Join("web", "templates", "*")

	tpl := template.Must(template.ParseGlob(tplPath))

	return tpl

}
