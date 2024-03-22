package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/effiong-jr/snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	snippets, err := app.snippets.Latest()

	if err != nil {
		if errors.Is(err, models.ErrorNoRecord) {
			app.notFound(w)
		} else {
			app.serveError(w, err)
		}

		return
	}

	for _, snippet := range snippets {

		fmt.Fprintf(w, "%v", snippet)
	}

	if err != nil {
		app.serveError(w, err)
	}

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrorNoRecord) {
			app.notFound(w)
		} else {
			app.serveError(w, err)

		}

		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.html",
		// "./ui/html/partials/nav.html",
		"./ui/html/pages/view.html",
	}

	t, err := template.ParseFiles(files...)

	if err != nil {
		app.serveError(w, err)
		return
	}

	err = t.ExecuteTemplate(w, "base", snippet)

	// fmt.Fprintf(w, "%v", snippet)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji\nBut slowly, slowly\n\n- Kobayashi Issa"
	expires := 7

	err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serveError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", 1), http.StatusSeeOther)

}
