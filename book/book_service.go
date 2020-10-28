package book

import (
	"encoding/json"
	"fmt"
	"github.com/KiranGosavi/gopractice/cors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//SetupRoutes to setup the API routes
func SetupRoutes(basePath string)  {
	fmt.Println(basePath)
	singleBookHandler :=http.HandlerFunc(singleBookHandler)
	booksHandler :=http.HandlerFunc(allBooksHandler)
	http.Handle("/api/books", cors.MiddlewareHandler(booksHandler))
	http.Handle("/api/books/", cors.MiddlewareHandler(singleBookHandler))

}
func singleBookHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("In singleHandler")
	urlPathSegment :=strings.Split(r.URL.Path, "books/")

	bookID,err :=strconv.Atoi(urlPathSegment[len(urlPathSegment)-1])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	singleBook :=getBook(bookID)
	if singleBook == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method{

	case http.MethodGet:
			bookJSON, err :=json.Marshal(singleBook)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			_,err = w.Write(bookJSON)
			if err != nil {
				log.Fatal(err)
			}
	case http.MethodPut:
			var books Book
			err =json.NewDecoder(r.Body).Decode(&books)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			_, err :=addOrUpdateBook(books)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
	case http.MethodDelete:
			removeBook(bookID)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
func allBooksHandler(w http.ResponseWriter,r *http.Request) {
	//fmt.Println("In books handler")
	switch r.Method {
	case http.MethodGet:
		bookList :=getAllBooks()
		bookJSON,err := json.Marshal(bookList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_,err =w.Write(bookJSON)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		var books Book
		err :=json.NewDecoder(r.Body).Decode(&books)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_,err =addOrUpdateBook(books)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
