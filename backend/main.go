package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Article is ...
type Article struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	ID       string
	Content  string `json:"content"`
	Creation string
}

type articleHandlers struct {
	sync.Mutex
	store map[string]Article
}

//CHECKS REQUEST GOING TO /articles/ IS POST OR GET
func (h *articleHandlers) getAllArticles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return

	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

//POST ONE ARTICLE (json body is in routes.rest)
//POST http://localhost:8080/articles
func (h *articleHandlers) post(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s' ", ct)))
		return
	}

	var article Article
	err = json.Unmarshal(bodyBytes, &article)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	//article.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	article.ID = fmt.Sprintf("%d", len(h.store)+1)
	article.Creation = fmt.Sprintf("%s", time.Now().String())

	h.Lock()
	h.store[article.ID] = article
	defer h.Unlock()
}

//GET ALL ARTICLES (with pagination)
//GET http://localhost:8080/articles?page=2
func (h *articleHandlers) get(w http.ResponseWriter, r *http.Request) {

	articles := make([]Article, 0, 1000)

	paginationURL := strings.Split(r.URL.String(), "?page=")
	pageNumber := paginationURL[1]
	num, err := strconv.Atoi(pageNumber)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	starting := (num - 1) * 3
	ending := starting + 2
	if ending > len(h.store)-1 {
		ending = len(h.store) - 1
	}

	h.Lock()
	for i := starting; i <= ending; i++ {
		articles = append(articles, h.store[strconv.Itoa(i+1)])
	}
	h.Unlock()

	jsonBytes, err := json.Marshal(articles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

//GET ONE ARTICLE
//GET http://localhost:8080/articles/1
func (h *articleHandlers) getOneArticle(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.String(), "/")

	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	h.Lock()

	article, ok := h.store[parts[2]]

	h.Unlock()

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(article)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

//SEARCH ARTICLES(case insensitive)
//GET GET http://localhost:8080/articles/search?q=abc
func (h *articleHandlers) searcharticle(w http.ResponseWriter, r *http.Request) {
	articles := make([]Article, 0, 1000)
	queryURL := strings.Split(r.URL.String(), "?q=")

	if len(queryURL) != 2 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	h.Lock()

	query := strings.ToLower(queryURL[1])

	//i := 0
	for _, element := range h.store {
		option1 := strings.Contains(strings.ToLower(element.Title), query)
		option2 := strings.Contains(strings.ToLower(element.Subtitle), query)
		option3 := strings.Contains(strings.ToLower(element.Content), query)
		if option1 || option2 || option3 {
			articles = append(articles, element)
		}
	}

	h.Unlock()

	if len(articles) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(articles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func newArticleHandlers() *articleHandlers {
	return &articleHandlers{
		store: map[string]Article{
			"1": Article{
				Title:    "PS4 Firmware Update 8.00 Available to Download Today",
				Subtitle: "PS4 firmware",
				ID:       "1",
				Content:  "Pre-set avatars are also getting a bit of a facelift with new pictures to choose from. They include profile pictures from PS4 titles such as Bloodborne, Journey, Ghost of Tsushima, The Last of Us: Part II, and Uncharted 4: A Thief’s End. You can check them all out through here.",
				Creation: "2020-10-14 00:40:32.0093369 +0530 IST m=+222.416472201",
			},
		},
	}
}

func main() {
	articleHandlers := newArticleHandlers()
	http.HandleFunc("/articles", articleHandlers.getAllArticles)
	http.HandleFunc("/articles/search", articleHandlers.searcharticle)
	http.HandleFunc("/articles/", articleHandlers.getOneArticle)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
