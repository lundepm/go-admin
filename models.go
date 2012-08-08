package main

import (
	"html/template"
	"regexp"	
)

const tmplPath = "templates/"
const viewPath = "views/"
const lenPath = len("/view/")

var templates = template.Must(template.ParseFiles(tmplPath+"header.html", tmplPath+"home.html", tmplPath+"footer.html", tmplPath+"view.html", tmplPath+"edit.html", tmplPath+"route.html"))
var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")

type config struct {
	conf []conf
}

type conf struct {
	Type string
	Value string
}

type page struct {
	MenuHTML template.HTML
	TocHTML template.HTML
	ContentHTML template.HTML
}

type content struct {
	Title string
	Body  []byte
	Route Route
}

/**** MENUMAPPING ****/
type Menu struct {
	Item 			[]Item 		`xml:"item"`
}

type Item struct {
	Number 			string 		`xml:"number,attr"`
	MenuText 		string 		`xml:"text"`
	Link			string		`xml:"link"`
	ToolTip			string		`xml:"tooltip"`
}

/**** ROUTEMAPPING ****/
type Routes struct {
	Routes			[]Route		`xml:"Route"`
}

type Route struct {
	Id 				string 		`xml:"Id,attr"`
	Impl_repo 		string		`xml:"Impl_repo"`
	Src 			string		`xml:"Src"`
	Disabled 		bool		`xml:"Disabled"`
	Http_code 		int			`xml:"Http_code"`
	Redir_to_id 	string		`xml:"Redir_to_id"`
	Partial_redir 	string		`xml:"Partial_redir"`
	Names 			Names		`xml:"Names"`
	Meta_values 	Meta_values	`xml:"Meta_values"`
}

type Names struct {
	Name			[]Name		`xml:"Name"`
}

type Name struct {
	Lang			string		`xml:"Lang,attr"`
	Name			string		`xml:"Name"`
	Tooltip			string		`xml:"Tooltip"`
}

type Meta_values struct {
	MetaValue		[]MetaValue	`xml:"MetaValue"`
}

type MetaValue struct {
	Key				string		`xml:"Key,attr"`
	Value			string		`xml:"Value"`
}

