package shortener

import (
	"net/http"
	"appengine"
	"appengine/memcache"
	"html/template"
	"fmt"
	"log"
	"strings"
)

const BaseUrl  = "localhost:8080"


func render(w http.ResponseWriter, context interface{}, template_path string) {
	tmpl := template.New("PAGE")

	tmpl = template.Must(template.ParseFiles(template_path))
	template_paths := strings.Split(template_path, "/")
	template_name := template_paths[len(template_paths)-1]

	if err := tmpl.ExecuteTemplate(w, template_name, context); err != nil {
		fmt.Fprint(w, err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	context := map[string]interface{}{
		"Car": "value",
		"papa": "jeden",
	}
	render(w, context, "shortener/templates/index.html")

}

func CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusNotFound)
		return
	}
	originUrl := r.FormValue("originUrl")
	log.Print(r.Method)
	log.Print(originUrl)
	urlSuffix := GenerateUrlSuffix()
	ctx := appengine.NewContext(r)
	shorterUrl := BaseUrl + "/" + urlSuffix
	item := &memcache.Item{
		Key:   urlSuffix,
		Value: []byte(originUrl),
	}
	if err := memcache.Add(ctx, item); err != nil {
		fmt.Println(err)
	}
	fmt.Fprint(w, shorterUrl)
}
