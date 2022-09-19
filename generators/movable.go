package generators

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/bindasov/spaceBattle/adapters"
	"github.com/fatih/camelcase"
	"go/format"
	"html/template"
	"log"
	"os"
	"reflect"
	"strings"
)

type MethodName struct {
	Action   string
	Property string
}

type Method struct {
	Name       *MethodName
	InputParam string
	Output     string
}

func ExtractMethodParams(s string) string {
	i := strings.Index(s, "(")
	if i >= 0 {
		j := strings.Index(s, ")")
		if j >= 0 {
			return s[i+1 : j]
		}
	}
	return ""
}

func ExtractOutput(s string) string {
	i := strings.Index(s, ")")
	if i >= 0 {
		j := len(s)
		if j >= i {
			return s[i+1 : j]
		}
	}
	return ""
}

func processTemplate(fileName string, outputFile string, data []*Method) {
	tmpl := template.Must(template.New("").Funcs(sprig.FuncMap()).ParseFiles(fileName))
	var processed bytes.Buffer
	err := tmpl.ExecuteTemplate(&processed, fileName, data)
	if err != nil {
		log.Fatalf("Unable to parse data into template: %v\n", err)
	}
	fmt.Println(string(processed.Bytes()))
	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		log.Fatalf("Could not format processed template: %v\n", err)
	}
	outputPath := "./adapters/" + outputFile
	fmt.Println("Writing file: ", outputPath)
	f, _ := os.Create(outputPath)
	w := bufio.NewWriter(f)
	w.WriteString(string(formatted))
	w.Flush()
}

func Generate() {
	var methods []*Method

	t := reflect.TypeOf((*adapters.MovableAdapter)(nil)).Elem()
	for i := 0; i < t.NumMethod(); i++ {
		splitted := camelcase.Split(t.Method(i).Name)
		method := &Method{
			Name: &MethodName{
				Action:   splitted[0],
				Property: splitted[1],
			},
			InputParam: ExtractMethodParams(t.Method(i).Type.String()),
			Output:     ExtractOutput(t.Method(i).Type.String()),
		}
		methods = append(methods, method)
	}
	processTemplate("movableAdapter.tmpl", "movableGenerated.go", methods)

}
