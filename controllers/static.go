package controllers

import (
	"html/template"
	"net/http"

	"github.com/mnmozi/lenslocked/views"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{Question: "Is there a free version?",
			Answer: "sure there is."},
		{
			Question: "when will you finish this course man?",
			Answer:   "Soon insha2 alah",
		}, {
			Question: "How do I contact support?",
			Answer:   `Email us - <a href="mailto:suppoert@lenslocked.com"> support@lenslocked.com<\a>`,
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
