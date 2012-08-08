package main

import (
	"io/ioutil"
	"encoding/xml"
	"fmt"
)

func (c *content) save() error {
	filename := viewPath + c.Title + ".txt"
	return ioutil.WriteFile(filename, c.Body, 0600)
}

func loadContent(title string) (*content, error) {
	filename := viewPath + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &content{Title: title, Body: body}, nil
}

func loadRoute(id string) (*content) {
	bytes, _ := ioutil.ReadFile("routesws.xml")
	var routes Routes
	title := "Yap"
	var pos int
	err := xml.Unmarshal([]byte(string(bytes)), &routes)
	if err != nil {
		fmt.Println("FUCK ", err.Error())
	}
	for i := range routes.Routes {
		if routes.Routes[i].Id == id {
			pos = i
		}
	}
	return &content{Title: title, Body: nil, Route: routes.Routes[pos]}
}

func loadMenu() (data interface{}) {
	bytes, _ := ioutil.ReadFile("dummy.xml")
	var menu Menu
	xml.Unmarshal([]byte(string(bytes)), &menu)
	return menu.Item
}

func loadToc() (data interface{}) {
	bytes, _ := ioutil.ReadFile("routesws.xml")
	var routes Routes
	//var lang interface{}
	err := xml.Unmarshal([]byte(string(bytes)), &routes)
	if err != nil {
		fmt.Println("FUCK ", err.Error())
	}
	for i := range routes.Routes {
		for j := range routes.Routes[i].Names.Name {
			if routes.Routes[i].Names.Name[j].Lang == "en" {
				//add(lang, routes.Routes[i])
			}
		}
	}
	
	return routes.Routes
}
