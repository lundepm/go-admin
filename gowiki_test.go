package main

import "testing"

func TestLoadContent(t *testing.T) {
	_, err := loadContent("dummy653")
	if err == nil{
		t.Log("Invalid URL")
		t.Fail()
	}
	if err != nil{
		t.Log("Success")
	}
}
