package controller

import (
	"os"
	"html/template"
	"io/ioutil"
	"net/http"
	"fmt"
	"errors"
)

// PopulateTemplates func
// Create map templates name to templates.Template
func PopulateTemplate() map[string]*template.Template {
	const basePath = "templates"
	result := make(map[string]*template.Template)

	layout := template.Must(template.ParseFiles(basePath + "/_base.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}

//session

//set
func setSessionUser(w http.ResponseWriter, r *http.Request, userName string) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Values["user"] = userName
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

//get
func getSessionUser(r *http.Request) (string, error) {
	var userName string
	session, err := store.Get(r, sessionName)
	if err != nil {
		return "", err
	}
	val := session.Values["user"]
	fmt.Println("val:", val)
	userName, ok := val.(string)
	if !ok {
		return "", errors.New("can not get session user")
	}
	fmt.Println("username:", userName)
	return userName, nil
}

//clear
func clearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil

}
