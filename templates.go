package main

import (
	"bytes"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
)

var (
	tmplMainPage  = "mainpage.html"
	tmplArticle   = "article.html"
	tmplArchive   = "archive.html"
	tmplNotesWeek = "notes_week.html"
	tmplNotesTag  = "notes_tag.html"
	tmplNotesNote = "notes_note.html"
	tmpl404       = "404.html"
	templateNames = [...]string{
		tmplMainPage,
		tmplArticle,
		tmplArchive,
		tmplNotesWeek,
		tmplNotesTag,
		tmplNotesNote,
		tmpl404,
		"analytics.html",
		"page_navbar.html",
		"tagcloud.js",
	}
	templatePaths   []string
	templates       *template.Template
	reloadTemplates = true
)

func getTemplates() *template.Template {
	if reloadTemplates || (nil == templates) {
		if 0 == len(templatePaths) {
			for _, name := range templateNames {
				templatePaths = append(templatePaths, filepath.Join("tmpl", name))
			}
		}
		templates = template.Must(template.ParseFiles(templatePaths...))
	}
	return templates
}

func serveTemplate(w http.ResponseWriter, templateName string, model interface{}) bool {
	var buf bytes.Buffer
	if err := getTemplates().ExecuteTemplate(&buf, templateName, model); err != nil {
		logger.Errorf("Failed to execute template %q, error: %s", templateName, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}

	// at this point we ignore error
	w.Header().Set("Content-Length", strconv.Itoa(len(buf.Bytes())))
	w.Write(buf.Bytes())
	return true
}
