package main

import (
	"../model"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "os"
	"time"

	"github.com/rs/cors"
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

	art = append(art, model.Article{Id: "1", Title: "Why should we use Golang?", Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas rhoncus enim in velit vulputate mollis. Cras et ipsum vulputate elit dignissim suscipit id eu felis. Integer pulvinar facilisis interdum. Nulla luctus tincidunt velit id sodales. Sed in faucibus lacus, nec venenatis risus. Integer dignissim vestibulum arcu. Proin elementum dapibus euismod. Vivamus bibendum risus in arcu pretium viverra. Donec sollicitudin elementum ex laoreet congue. Donec in porta lorem. Sed scelerisque interdum lectus in congue. Cras lacinia velit in dolor feugiat lobortis. Morbi et tellus augue. Phasellus euismod eleifend mauris, et laoreet tellus viverra ac. Donec sit amet suscipit tortor.\n\nSed tortor ex, tempor vel facilisis in, lobortis in magna. Nunc interdum, purus ut lacinia convallis, enim sem consequat purus, eget pharetra tellus tellus vitae massa. Pellentesque bibendum varius sapien, a fermentum est pharetra ut. Praesent libero sapien, iaculis nec justo ac, maximus vehicula metus. Fusce vitae ultrices est. Nulla mollis mauris non sem posuere, nec consequat ligula pharetra. Aenean at tincidunt enim, non pharetra risus. Donec at magna molestie, scelerisque tellus quis, venenatis mi. Suspendisse enim odio, volutpat a facilisis vitae, iaculis id nisl.\n\nCurabitur mattis id felis sit amet malesuada. Nam at massa et tellus scelerisque efficitur eu eu est. Suspendisse condimentum elit ut dignissim suscipit. Nam lacinia, felis sit amet accumsan porttitor, lectus justo consequat ex, ut tincidunt odio mauris sed arcu. Cras dignissim ipsum turpis, et tempus metus bibendum in. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Aliquam vehicula, nulla eu sodales maximus, lorem orci suscipit tortor, vitae rutrum sapien velit eu est.\n\nProin finibus in nulla at sollicitudin. In augue nisi, auctor in eros et, ullamcorper laoreet eros. Etiam mattis ligula eget magna tempor, vitae cursus nunc elementum. Aliquam pharetra turpis nec massa ultrices bibendum. Mauris sit amet vestibulum leo. Phasellus sem elit, malesuada lobortis ornare in, vestibulum molestie orci. Vestibulum commodo, neque nec aliquet viverra, purus turpis tempus ex, id maximus purus mauris ac ligula. Integer dignissim velit a nibh ullamcorper pharetra. Morbi lectus leo, tempor sed felis at, vulputate imperdiet dui. Duis placerat turpis tellus, a aliquam enim porttitor id.\n\nDonec ut aliquet augue, in dictum purus. In sed tellus nisi. Nunc at dolor sem. Nullam blandit et augue tempus tempor. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Aenean egestas dui fermentum sodales pulvinar. Integer suscipit egestas leo vitae porttitor. Praesent in euismod neque, vitae ultricies libero. Nullam eget nulla massa. Cras nec nunc nec sem pharetra scelerisque ac sodales mi. Donec in rhoncus metus, et imperdiet tortor. Pellentesque malesuada tortor non diam gravida, vitae pellentesque libero tempor.\n\nSuspendisse enim nulla, imperdiet at scelerisque pretium, bibendum vitae mi. In volutpat nibh nibh, at pretium libero condimentum ut. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Donec gravida, elit tempus iaculis cursus, mauris leo fringilla elit, et malesuada nibh diam quis mi. Quisque ut lacus justo. Suspendisse non molestie lorem, id finibus augue. Aenean vitae metus vehicula, semper lacus et, aliquet lacus. Vivamus eget molestie urna. Sed id ipsum non sem facilisis luctus vitae ut erat. Nullam vitae dui sit amet tellus semper sodales ac ac diam. Mauris at lacus suscipit nunc dictum accumsan eget sit amet libero. Maecenas suscipit elit non mi viverra, sit amet lacinia turpis vestibulum. Vestibulum est quam, tincidunt ac euismod vitae, sollicitudin vel urna.\n\n", Date: time.Now().Unix()})
	art = append(art, model.Article{Id: "2", Title: "Who is the best programmer and why Michael Kaczynski?", Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas rhoncus enim in velit vulputate mollis. Cras et ipsum vulputate elit dignissim suscipit id eu felis. Integer pulvinar facilisis interdum. Nulla luctus tincidunt velit id sodales. Sed in faucibus lacus, nec venenatis risus. Integer dignissim vestibulum arcu. Proin elementum dapibus euismod. Vivamus bibendum risus in arcu pretium viverra. Donec sollicitudin elementum ex laoreet congue. Donec in porta lorem. Sed scelerisque interdum lectus in congue. Cras lacinia velit in dolor feugiat lobortis. Morbi et tellus augue. Phasellus euismod eleifend mauris, et laoreet tellus viverra ac. Donec sit amet suscipit tortor.\n\nSed tortor ex, tempor vel facilisis in, lobortis in magna. Nunc interdum, purus ut lacinia convallis, enim sem consequat purus, eget pharetra tellus tellus vitae massa. Pellentesque bibendum varius sapien, a fermentum est pharetra ut. Praesent libero sapien, iaculis nec justo ac, maximus vehicula metus. Fusce vitae ultrices est. Nulla mollis mauris non sem posuere, nec consequat ligula pharetra. Aenean at tincidunt enim, non pharetra risus. Donec at magna molestie, scelerisque tellus quis, venenatis mi. Suspendisse enim odio, volutpat a facilisis vitae, iaculis id nisl.\n\nCurabitur mattis id felis sit amet malesuada. Nam at massa et tellus scelerisque efficitur eu eu est. Suspendisse condimentum elit ut dignissim suscipit. Nam lacinia, felis sit amet accumsan porttitor, lectus justo consequat ex, ut tincidunt odio mauris sed arcu. Cras dignissim ipsum turpis, et tempus metus bibendum in. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Aliquam vehicula, nulla eu sodales maximus, lorem orci suscipit tortor, vitae rutrum sapien velit eu est.\n\nProin finibus in nulla at sollicitudin. In augue nisi, auctor in eros et, ullamcorper laoreet eros. Etiam mattis ligula eget magna tempor, vitae cursus nunc elementum. Aliquam pharetra turpis nec massa ultrices bibendum. Mauris sit amet vestibulum leo. Phasellus sem elit, malesuada lobortis ornare in, vestibulum molestie orci. Vestibulum commodo, neque nec aliquet viverra, purus turpis tempus ex, id maximus purus mauris ac ligula. Integer dignissim velit a nibh ullamcorper pharetra. Morbi lectus leo, tempor sed felis at, vulputate imperdiet dui. Duis placerat turpis tellus, a aliquam enim porttitor id.\n\nDonec ut aliquet augue, in dictum purus. In sed tellus nisi. Nunc at dolor sem. Nullam blandit et augue tempus tempor. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Aenean egestas dui fermentum sodales pulvinar. Integer suscipit egestas leo vitae porttitor. Praesent in euismod neque, vitae ultricies libero. Nullam eget nulla massa. Cras nec nunc nec sem pharetra scelerisque ac sodales mi. Donec in rhoncus metus, et imperdiet tortor. Pellentesque malesuada tortor non diam gravida, vitae pellentesque libero tempor.\n\nSuspendisse enim nulla, imperdiet at scelerisque pretium, bibendum vitae mi. In volutpat nibh nibh, at pretium libero condimentum ut. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Donec gravida, elit tempus iaculis cursus, mauris leo fringilla elit, et malesuada nibh diam quis mi. Quisque ut lacus justo. Suspendisse non molestie lorem, id finibus augue. Aenean vitae metus vehicula, semper lacus et, aliquet lacus. Vivamus eget molestie urna. Sed id ipsum non sem facilisis luctus vitae ut erat. Nullam vitae dui sit amet tellus semper sodales ac ac diam. Mauris at lacus suscipit nunc dictum accumsan eget sit amet libero. Maecenas suscipit elit non mi viverra, sit amet lacinia turpis vestibulum. Vestibulum est quam, tincidunt ac euismod vitae, sollicitudin vel urna.\n\n", Date: time.Date(2020, 01, 01, 15, 0, 0, 2525252, time.UTC).Unix()})
	art = append(art, model.Article{Id: "3", Title: "NaN != NaN, why?", Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas rhoncus enim in velit vulputate mollis. Cras et ipsum vulputate elit dignissim suscipit id eu felis. Integer pulvinar facilisis interdum. Nulla luctus tincidunt velit id sodales. Sed in faucibus lacus, nec venenatis risus. Integer dignissim vestibulum arcu. Proin elementum dapibus euismod. Vivamus bibendum risus in arcu pretium viverra. Donec sollicitudin elementum ex laoreet congue. Donec in porta lorem. Sed scelerisque interdum lectus in congue. Cras lacinia velit in dolor feugiat lobortis. Morbi et tellus augue. Phasellus euismod eleifend mauris, et laoreet tellus viverra ac. Donec sit amet suscipit tortor.\n\nSed tortor ex, tempor vel facilisis in, lobortis in magna. Nunc interdum, purus ut lacinia convallis, enim sem consequat purus, eget pharetra tellus tellus vitae massa. Pellentesque bibendum varius sapien, a fermentum est pharetra ut. Praesent libero sapien, iaculis nec justo ac, maximus vehicula metus. Fusce vitae ultrices est. Nulla mollis mauris non sem posuere, nec consequat ligula pharetra. Aenean at tincidunt enim, non pharetra risus. Donec at magna molestie, scelerisque tellus quis, venenatis mi. Suspendisse enim odio, volutpat a facilisis vitae, iaculis id nisl.\n\nCurabitur mattis id felis sit amet malesuada. Nam at massa et tellus scelerisque efficitur eu eu est. Suspendisse condimentum elit ut dignissim suscipit. Nam lacinia, felis sit amet accumsan porttitor, lectus justo consequat ex, ut tincidunt odio mauris sed arcu. Cras dignissim ipsum turpis, et tempus metus bibendum in. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Aliquam vehicula, nulla eu sodales maximus, lorem orci suscipit tortor, vitae rutrum sapien velit eu est.\n\nProin finibus in nulla at sollicitudin. In augue nisi, auctor in eros et, ullamcorper laoreet eros. Etiam mattis ligula eget magna tempor, vitae cursus nunc elementum. Aliquam pharetra turpis nec massa ultrices bibendum. Mauris sit amet vestibulum leo. Phasellus sem elit, malesuada lobortis ornare in, vestibulum molestie orci. Vestibulum commodo, neque nec aliquet viverra, purus turpis tempus ex, id maximus purus mauris ac ligula. Integer dignissim velit a nibh ullamcorper pharetra. Morbi lectus leo, tempor sed felis at, vulputate imperdiet dui. Duis placerat turpis tellus, a aliquam enim porttitor id.\n\nDonec ut aliquet augue, in dictum purus. In sed tellus nisi. Nunc at dolor sem. Nullam blandit et augue tempus tempor. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Aenean egestas dui fermentum sodales pulvinar. Integer suscipit egestas leo vitae porttitor. Praesent in euismod neque, vitae ultricies libero. Nullam eget nulla massa. Cras nec nunc nec sem pharetra scelerisque ac sodales mi. Donec in rhoncus metus, et imperdiet tortor. Pellentesque malesuada tortor non diam gravida, vitae pellentesque libero tempor.\n\nSuspendisse enim nulla, imperdiet at scelerisque pretium, bibendum vitae mi. In volutpat nibh nibh, at pretium libero condimentum ut. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Donec gravida, elit tempus iaculis cursus, mauris leo fringilla elit, et malesuada nibh diam quis mi. Quisque ut lacus justo. Suspendisse non molestie lorem, id finibus augue. Aenean vitae metus vehicula, semper lacus et, aliquet lacus. Vivamus eget molestie urna. Sed id ipsum non sem facilisis luctus vitae ut erat. Nullam vitae dui sit amet tellus semper sodales ac ac diam. Mauris at lacus suscipit nunc dictum accumsan eget sit amet libero. Maecenas suscipit elit non mi viverra, sit amet lacinia turpis vestibulum. Vestibulum est quam, tincidunt ac euismod vitae, sollicitudin vel urna.\n\n", Date: time.Now().Unix()})
	art = append(art, model.Article{Id: "4", Title: "Oh God, who recommended me that language?!", Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas rhoncus enim in velit vulputate mollis. Cras et ipsum vulputate elit dignissim suscipit id eu felis. Integer pulvinar facilisis interdum. Nulla luctus tincidunt velit id sodales. Sed in faucibus lacus, nec venenatis risus. Integer dignissim vestibulum arcu. Proin elementum dapibus euismod. Vivamus bibendum risus in arcu pretium viverra. Donec sollicitudin elementum ex laoreet congue. Donec in porta lorem. Sed scelerisque interdum lectus in congue. Cras lacinia velit in dolor feugiat lobortis. Morbi et tellus augue. Phasellus euismod eleifend mauris, et laoreet tellus viverra ac. Donec sit amet suscipit tortor.\n\nSed tortor ex, tempor vel facilisis in, lobortis in magna. Nunc interdum, purus ut lacinia convallis, enim sem consequat purus, eget pharetra tellus tellus vitae massa. Pellentesque bibendum varius sapien, a fermentum est pharetra ut. Praesent libero sapien, iaculis nec justo ac, maximus vehicula metus. Fusce vitae ultrices est. Nulla mollis mauris non sem posuere, nec consequat ligula pharetra. Aenean at tincidunt enim, non pharetra risus. Donec at magna molestie, scelerisque tellus quis, venenatis mi. Suspendisse enim odio, volutpat a facilisis vitae, iaculis id nisl.\n\nCurabitur mattis id felis sit amet malesuada. Nam at massa et tellus scelerisque efficitur eu eu est. Suspendisse condimentum elit ut dignissim suscipit. Nam lacinia, felis sit amet accumsan porttitor, lectus justo consequat ex, ut tincidunt odio mauris sed arcu. Cras dignissim ipsum turpis, et tempus metus bibendum in. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Aliquam vehicula, nulla eu sodales maximus, lorem orci suscipit tortor, vitae rutrum sapien velit eu est.\n\nProin finibus in nulla at sollicitudin. In augue nisi, auctor in eros et, ullamcorper laoreet eros. Etiam mattis ligula eget magna tempor, vitae cursus nunc elementum. Aliquam pharetra turpis nec massa ultrices bibendum. Mauris sit amet vestibulum leo. Phasellus sem elit, malesuada lobortis ornare in, vestibulum molestie orci. Vestibulum commodo, neque nec aliquet viverra, purus turpis tempus ex, id maximus purus mauris ac ligula. Integer dignissim velit a nibh ullamcorper pharetra. Morbi lectus leo, tempor sed felis at, vulputate imperdiet dui. Duis placerat turpis tellus, a aliquam enim porttitor id.\n\nDonec ut aliquet augue, in dictum purus. In sed tellus nisi. Nunc at dolor sem. Nullam blandit et augue tempus tempor. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Aenean egestas dui fermentum sodales pulvinar. Integer suscipit egestas leo vitae porttitor. Praesent in euismod neque, vitae ultricies libero. Nullam eget nulla massa. Cras nec nunc nec sem pharetra scelerisque ac sodales mi. Donec in rhoncus metus, et imperdiet tortor. Pellentesque malesuada tortor non diam gravida, vitae pellentesque libero tempor.\n\nSuspendisse enim nulla, imperdiet at scelerisque pretium, bibendum vitae mi. In volutpat nibh nibh, at pretium libero condimentum ut. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Donec gravida, elit tempus iaculis cursus, mauris leo fringilla elit, et malesuada nibh diam quis mi. Quisque ut lacus justo. Suspendisse non molestie lorem, id finibus augue. Aenean vitae metus vehicula, semper lacus et, aliquet lacus. Vivamus eget molestie urna. Sed id ipsum non sem facilisis luctus vitae ut erat. Nullam vitae dui sit amet tellus semper sodales ac ac diam. Mauris at lacus suscipit nunc dictum accumsan eget sit amet libero. Maecenas suscipit elit non mi viverra, sit amet lacinia turpis vestibulum. Vestibulum est quam, tincidunt ac euismod vitae, sollicitudin vel urna.\n\n", Date: time.Now().Unix()})

	if len(art) != 4 {
		return nil, errors.New("initialData: Cant load initial data")
	}
	return art, nil
}

func runApi() (string, error) {

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
		AllowedMethods: []string{"GET"}, // Allowing only get, just an example
	})

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

		// start server listen
		// with error handling
		log.Fatal(http.ListenAndServe(":8001", c.Handler(muxRouter)))
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
