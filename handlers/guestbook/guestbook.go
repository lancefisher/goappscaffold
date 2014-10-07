package guestbook

import (
	"appengine"
	"html/template"
	"net/http"

	"github.com/lancefisher/goappscaffold/models"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	//c.Infof("guestbook.HandleFunc")
	var g *models.Greeting
	count := 10
	greetings := make([]models.Greeting, 0, count)
	if err := g.GetRecent(c, &greetings, count); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := guestbookTemplate.Execute(w, greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var guestbookTemplate = template.Must(template.New("book").Parse(`
<html>
  <head>
    <title>Go Guestbook</title>
  </head>
  <body>
    {{range .}}
      {{with .Author}}
        <p><b>{{.}}</b> wrote:</p>
      {{else}}
        <p>An anonymous person wrote:</p>
      {{end}}
      <pre>{{.Content}}</pre>
    {{end}}
    <form action="/sign" method="post">
      <div><textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Guestbook"></div>
    </form>
  </body>
</html>
`))
