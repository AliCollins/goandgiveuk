// goandgive project goandgive.go
package goandgiveuk

import (
	"fmt"
	"net/http"

	// "appengine"
)

func init() {
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)

	fmt.Fprint(w, rootForm)
}

const rootForm = `
  <!DOCTYPE html>
    <html>
      <head>
        <meta charset="utf-8">
        <title>Go and Give UK</title>
        <link rel="stylesheet" href="/stylesheets/goandgive.css">        
      </head>
      <body>
        <h1>Go and Give</h1>
        <h2>The start of the Go and Give website</h2>
        <p>More information will be put here as it becomes available!</p>
		Just a small change to check that everything gets through OK...and again.  No way!!
      </body>
    </html>
`
