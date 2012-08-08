package main

import (
	"net/http"
	"html/template"
//	"strings"
	"bytes"
)

func renderTemplate(w http.ResponseWriter, tmpl string, c *content) {
	var buf bytes.Buffer
	err := templates.ExecuteTemplate(&buf, tmpl + ".html", c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(getPage(tmpl + ".html", c))
}

func parseTemplate(file string, data interface{}) (out []byte, error error) {
	var buf bytes.Buffer
	t, err := template.ParseFiles(file)
	if err != nil {
		return nil, err
	}
	err = t.Execute(&buf, data)
	if err!= nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func getPage(file string, data interface{}) []byte {
//	var active string
//	if strings.Contains(file, "project") {
//		active = "Projects"
//	} else if strings.Contains(file, "about") {
//		active = "About"
//	} else if strings.Contains(file, "post") {
//		active = "Archive"
//	} else if strings.Contains(file, "blog") || strings.Contains(file, "home") {
//		active = "Blog"
//	} else {
//		active = ""
//	}
	menu, error := parseTemplate("templates/menu.html", map[string]interface{} {"Nodes": loadMenu()})
	if error != nil {
		print(error.Error())
	}
	toc, error := parseTemplate("templates/toc.html", map[string]interface{} {"Nodes": loadToc()})
	if error != nil {
		print(error.Error())
	}	
	content, error := parseTemplate("templates/" + file , data)
	if error != nil {
		print(error.Error())
	}
	base, error := parseTemplate("templates/main.html", page{MenuHTML:template.HTML(menu), TocHTML:template.HTML(toc), ContentHTML:template.HTML(content)})
	if error != nil {
		return []byte("Internal server error...")
	}
	return base
}

/**** HANDLERS ****/
func rootHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home", nil)
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	p := loadRoute(r.URL.Path[lenPath:])
	renderTemplate(w, "route", p)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadContent(title)
	if err != nil {
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadContent(title)
	if err != nil {
		p = &content{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &content{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r * http.Request) {
		title := r.URL.Path[lenPath:]
		if !titleValidator.MatchString(title) {
			http.NotFound(w, r)
			return
		}
		fn(w, r, title)
	}
}
