package view

import (
	"text/template"
)

type Error struct {
	Error error
	Advice string
}

var ErrorTmpl = template.Must(template.New("error").Parse(`
<!doctype html>
<html>
<head>
</head>
<body>
<h1>The genomes@fi GFF3 validating service</h1>
<h2>Processing failed</h2>
<h3>{{.Error}}</h3>
<p>{{.Advice}}</p>
</body>
</html>
	`))