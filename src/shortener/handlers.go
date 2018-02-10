package shortener

import (
	"appengine"
	"appengine/datastore"
	"appengine/memcache"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"strings"
	"shortener/models"
	"net/http"
	"time"
)

const BaseUrl  = "http://localhost:8080"


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
	render(w, nil, "shortener/templates/index.html")
}

func CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	originUrl := r.FormValue("originUrl")
	log.Print(originUrl)
	urlSuffix := GenerateUrlSuffix()
	ctx := appengine.NewContext(r)
	shorterUrl := BaseUrl + "/" + urlSuffix
	item := &memcache.Item{
		Key:   urlSuffix,
		Value: []byte(originUrl),
		Expiration: time.Second*60*15,
	}
	e := new(models.UrlHistory)
	e.OriginalUrl = originUrl
	e.ShortUrl = shorterUrl
	e.Created = time.Now()
	k := datastore.NewKey(ctx, "UrlHistory", "stringID", 0, nil)
	if _, err := datastore.Put(ctx, k, e); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := memcache.Add(ctx, item); err != nil {
		fmt.Println(err)
	}
	context := map[string]interface{}{
		"resultUrl": shorterUrl,
	}
	render(w, context, "shortener/templates/result.html")
}

func OriginalRedirect(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	ctx := appengine.NewContext(r)
	log.Println(vars["urlHash"])
	urlItem, err := memcache.Get(ctx, vars["urlHash"])
	if err != nil{
		fmt.Fprint(w, "Short url died")
		return
	}
	var url string
	if url = string(urlItem.Value); !strings.HasPrefix(url, "http"){
		url = "http://" + url
	}
	http.Redirect(w, r, url , http.StatusSeeOther)
}
