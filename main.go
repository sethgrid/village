package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

// NewSite gets the templates all set up.
// If you are adding a new page, you must set up the template/file.tmpl and then add it to this constructor before adding the handler.
func NewSite() *Site {
	tmpl := make(map[string]*template.Template)
	tmpl["index.html"] = template.Must(template.New("index").ParseFiles("templates/index.tmpl", "templates/_latest_news.tmpl", "templates/_base.tmpl"))
	tmpl["contact.html"] = template.Must(template.New("contact").ParseFiles("templates/contact.tmpl", "templates/_latest_news.tmpl", "templates/_base.tmpl"))
	tmpl["donate.html"] = template.Must(template.New("donate").ParseFiles("templates/donate.tmpl", "templates/_latest_news.tmpl", "templates/_base.tmpl"))

	return &Site{templates: tmpl}
}

func main() {
	port := flag.String("p", "8100", "port to serve on")
	flag.Parse()

	// could put NewSite() here and make the hanlders method receivers on Site{}. This will be more performat.
	// however, I'm opting for constructing NewSite() in each handler to allow for hot reloading of templates.
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/index.html", indexHandler)

	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/contact.html", contactHandler)

	http.HandleFunc("/donate", donateHandler)
	http.HandleFunc("/donate.html", donateHandler)

	log.Printf("Serving on HTTP port: %s\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if err := NewSite().Render("index.html", w); err != nil {
		log.Println(err)
	}
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	if err := NewSite().Render("contact.html", w); err != nil {
		log.Println(err)
	}
}
func donateHandler(w http.ResponseWriter, r *http.Request) {
	if err := NewSite().Render("donate.html", w); err != nil {
		log.Println(err)
	}
}

// Site holds any state that we need
type Site struct {
	templates map[string]*template.Template
}

// Render is a simple wrapper around template.*template.Execute
func (s *Site) Render(templateName string, w io.Writer) error {
	if _, ok := s.templates[templateName]; !ok {
		var allTemplates []string
		for templateName := range s.templates {
			allTemplates = append(allTemplates, templateName)
		}
		return fmt.Errorf("unknown template name. Must be one of %v", allTemplates)
	}
	return s.templates[templateName].ExecuteTemplate(w, "base", struct {
		Active string
	}{
		Active: templateName,
	})
}
