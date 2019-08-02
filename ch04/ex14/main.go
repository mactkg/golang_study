// Issueshtml prints an HTML table of issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"html/template"

	"github.com/mactkg/golang_study/ch04/ex14/github"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{ len . }} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
  <th>MileStone</th>
</tr>
{{range . }}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'><img src='{{.User.HTMLURL}}.png' style="width: 1em; height: 1em;">{{.User.Login}}</a></td>
  <td><a href='/issue?id={{.Number}}'>{{.Title}}</a></td>
  <td>
  {{ if .MileStone }}
	<a href='{{.MileStone.HTMLURL}}'>{{.MileStone.Title}}</a>
  {{ else }}
	 -   
  {{end}}
  </td>
</tr>
{{end}}
</table>
`))

var issueView = template.Must(template.New("issueview").Parse(`
<h1><a href='{{.HTMLURL}}'>#{{.Number}}</a> {{.Title}}</h1>
<h2>By <a href='{{.User.HTMLURL}}'><img src='{{.User.HTMLURL}}.png' style="width: 2em; height: 2em;">{{.User.Login}}</a></h2>
<span>milestone: {{ if .MileStone }}<a href='{{.MileStone.HTMLURL}}'>{{.MileStone.Title}}</a>{{ else }} - {{end}}</span>
<p>{{.Body}}</p>
<a href="/">Back</a>
`))

func main() {
	data := make(map[int]*github.Issue)
	owner := os.Args[1]
	repo := os.Args[2]
	result, err := github.ListIssue(owner, repo)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	for _, issue := range result {
		data[issue.Number] = issue
	}
	fmt.Printf("Loaded %d issues\n", len(result))

	http.HandleFunc("/issue", func(w http.ResponseWriter, r *http.Request) {
		ids, ok := r.URL.Query()["id"]

		if !ok || len(ids[0]) < 1 {
			w.Write([]byte("Url Param 'id' is missing"))
			return
		}
		id, err := strconv.ParseInt(ids[0], 10, 64)
		if err != nil {
			w.Write([]byte("Can't parse param"))
			return
		}

		if err = issueView.Execute(w, data[int(id)]); err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err = issueList.Execute(w, result); err != nil {
			log.Fatal(err)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
