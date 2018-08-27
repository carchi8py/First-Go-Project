package cms

import (
	"net/http"
	"time"
)

//HandleNew Handles preview logic
func HandleNew(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Tmpl.ExecuteTemplate(w, "new", nil)
	case "POST":
		title := r.FormValue("title")
		content := r.FormValue("content")
		contentType := r.FormValue("content-type")
		r.ParseForm()

		if contentType == "page" {
			Tmpl.ExecuteTemplate(w, "page", &Page{
				Title: title,
				Content: content,
			})
			return
		}

		if contentType == "post" {
			Tmpl.ExecuteTemplate(w, "post", &Post{
				Title: title,
				Content: content,
			})
			return
		}
	default:
		http.Error(w, "Method not supported: "+r.Method, http.StatusMethodNotAllowed)
	}
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "Go Project CMS",
		Content: "Welcome to our home page!",
		Posts: []*Post{
			&Post{
				Title: "Hello, World",
				Content: "Hello world! Thank for coming to the site",
				DataPublished: time.Now(),
			},
			&Post{
				Title: "A post with Comments",
				Content: "Here's a controversial post. It's sure to attrack comments.",
				DataPublished: time.Now().Add(-time.Hour),
				Comments: []*Comment{
					&Comment{
						Author:        "Ben Tranter",
						Comment:       "Nevermind, I guess i just commented on my own post",
						DatePublished: time.Now().Add(-time.Hour / 2),
					},
				},
			},
		},
	}

	Tmpl.ExecuteTemplate(w, "page", p)
}
