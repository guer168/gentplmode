package code_gen

import (
	"bytes"
	"errors"
	"strings"
	"text/template"
	"time"

	"github.com/guer168/gentplmode/code_gen/mysql"
	"github.com/guer168/gentplmode/code_gen/postgresql"
	"github.com/guer168/gentplmode/utils"
)

func NewDbCodeGen(t string) (IDBMetaData, error) {
	switch strings.ToLower(t) {
	case "mysql":
		return &mysql.Gen{}, nil
	case "pg", "postgresql":
		return &postgresql.PGGen{}, nil
	}
	return nil, errors.New("invalid type")
}

func GenerateTemplate(templateText string, templateData interface{}, params map[string]interface{}) ([]byte, error) {
	t, err := template.New("tableTemplate").Funcs(template.FuncMap{
		"CamelizeStr":    utils.CamelizeStr,
		"FirstCharacter": utils.FirstCharacter,
		"FirstLowerWord": utils.FirstLowerWord,
		"RemovePrefix":   utils.RemovePrefix,
		"StrToLower":     utils.StrToLower,
		"Replace": func(old, new, src string) string {
			return strings.ReplaceAll(src, old, new)
		},
		"Add": func(a, b int) int {
			return a + b
		},
		"now": func() string {
			return time.Now().Format(time.RFC3339)
		},
		"param": func(name string) interface{} {
			//fmt.Printf("%s\n",params)
			if v, ok := params[name]; ok {
				return v
			}
			return ""
		},
	}).Parse(templateText)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, templateData); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
