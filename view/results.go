package view

import (
	"text/template"
)

type Results struct {
	Result string
	NError int
	Errors []string
	NWarning int
	Warnings []string
}

var ResultsTmpl = template.Must(template.New("results").Parse(`
<!doctype html>
<html>
<head>
</head>
<body>
<h1>The genomes@fi GFF3 validating service</h1>
<h2>Results</h2>
<h3>{{.Result}}</h3>
<h3>Errors: {{.NError}}</h3>
{{if .NError}}
<p>(showing only first 10)</p>
<pre>{{range .Errors}}{{.}}<br>{{else}}{{end}}</pre>
{{end}}

<h3>Warnings: {{.NWarning}}</h3>
{{if .NWarning}}<pre>{{range .Warnings}}{{.}}<br>{{else}}{{end}}</pre>{{end}}
</body>
</html>
	`))

// func main() {
// 	result := Result{"result", 5, []string{"err0r1", "err0r2"}, 6, []string{"warning1", "warning2"}}
// 	err = tmpl.Execute(os.Stdout, result)
// 	if err != nil { panic(err) }
// }