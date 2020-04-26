package main

import (
	"../model"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const (
	apiPath             = "/articles"
	apiContentTypeKey   = "Content-Type"
	apiContentTypeValue = "application/json"
)

var articles []model.Article

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(apiContentTypeKey, apiContentTypeValue)
	_ = json.NewEncoder(w).Encode(articles)
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(apiContentTypeKey, apiContentTypeValue)
	var newArticle model.Article
	_ = json.NewDecoder(r.Body).Decode(&newArticle)
	newArticle.Id = string(len(articles))
	articles = append(articles, newArticle)
	_ = json.NewEncoder(w).Encode(newArticle)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(apiContentTypeKey, apiContentTypeValue)
	params := mux.Vars(r)
	for _, val := range articles {
		if val.Id == params["id"] {
			_ = json.NewEncoder(w).Encode(val)
			return
		}
	}
	_ = json.NewEncoder(w).Encode(&model.Article{})
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(apiContentTypeKey, apiContentTypeValue)
	params := mux.Vars(r)
	for index, val := range articles {
		if val.Id == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			var updateArticle model.Article
			_ = json.NewDecoder(r.Body).Decode(&updateArticle)
			updateArticle.Id = params["id"]
			articles = append(articles, updateArticle)
			_ = json.NewEncoder(w).Encode(updateArticle)
			return
		}
	}
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(apiContentTypeKey, apiContentTypeValue)
	params := mux.Vars(r)
	for index, val := range articles {
		if val.Id == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			return
		}
	}
	_ = json.NewEncoder(w).Encode(articles)
}

func initialData() ([]model.Article, error) {
	var art []model.Article

	art = append(art, model.Article{Id: "1", Title: "Why we should use Golang?", Content: "Lorem ipsum", Date: time.Now().String()})
	art = append(art, model.Article{Id: "2", Title: "Who is the best programmer and why Michael Kaczynski?", Content: "Lorem ipsum", Date: time.Date(2020, 01, 01, 15, 0, 0, 2525252, time.UTC).String()})
	art = append(art, model.Article{Id: "3", Title: "NaN != NaN, why?", Content: "Lorem ipsum", Date: time.Now().String()})
	art = append(art, model.Article{Id: "4", Title: "Oh God, who recommended me that language?!", Content: "Lorem ipsum", Date: time.Now().String()})

	if len(art) != 4 {
		return nil, errors.New("initialData: Cant load initial data")
	}
	return art, nil
}

func runApi() (string, error) {
	muxRouter := mux.NewRouter()

	initData, errInitData := initialData()
	if errInitData != nil {
		return "", errInitData
	} else {
		articles = initData

		muxRouter.HandleFunc(apiPath+"/getAllArticles", getAllArticles).Methods("GET")
		muxRouter.HandleFunc(apiPath+"/createArticle", createArticle).Methods("POST")
		muxRouter.HandleFunc(apiPath+"/getArticle/{id}", getArticle).Methods("GET")
		muxRouter.HandleFunc(apiPath+"/updateArticle/{id}", updateArticle).Methods("PUT")
		muxRouter.HandleFunc(apiPath+"/deleteArticle/{id}", deleteArticle).Methods("DELETE")

		log.Fatal(http.ListenAndServe(":8001", muxRouter))
		return "healthy!", nil
	}
}

func main() {
	apiHealth, errApi := runApi()

	if errApi != nil {
		log.Printf("%v", errApi)
	}
	log.Printf("API: %v", apiHealth)
}
