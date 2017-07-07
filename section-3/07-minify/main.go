package main

import (
	"bytes"
	"fmt"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
)

func main() {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/javascript", js.Minify)

	buf := bytes.Buffer{}
	m.Minify("text/html", &buf, bytes.NewReader([]byte(`
		<!doctype html>
		<html>
			<head>
				<title>Test Minify</title>
			</head>
			<body>
				<p>Hello</p>
				<ul>
					<li>Item 1</li>
					<li>Item 2</li>
					<li>Item 3</li>
				</ul>
			</body>
		</html>
	`)))
	fmt.Println(buf.String())
}
