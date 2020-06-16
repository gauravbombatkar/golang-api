package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	//"log"
	"WebApp/DataAccess1"
	"fmt"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	//Write Response

	var books = DataAccess1.GetBooks()

	//convert slice of book(GOLang object into JSON)

	json.NewEncoder(w).Encode(books)
}


func GetBookById(w http.ResponseWriter, r *http.Request) {
	//Retrive Request data passed via url

	//Reading/Writing Header Values
	w.Header().Set("Content-Type", "application/json")

	header1 := r.Header["Content-Type"]
	header2 := r.Header["Token"]

	fmt.Println(header1, header2)

	//Reading Query string values

	//param1 := r.URL.Query().Get("param1")

	params := mux.Vars(r)
	var id = params["id"]
	id1, _ := strconv.Atoi(id)
	book := DataAccess1.GetBookById(id1)
	
	if book == (DataAccess1.Book{}) {
		http.Error(w, "Wrong Book id+++++++++++++", http.StatusNotFound)
		
		return
	} else {
		json.NewEncoder(w).Encode(book)
	}

}

func PostBook(w http.ResponseWriter, r *http.Request) {
	//Write Response

	//Get data from Payload/Request body
	var book DataAccess1.Book
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(book)
	DataAccess1.InsertBook(book)
	w.WriteHeader(http.StatusCreated)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//Write Response

	params := mux.Vars(r)
	var id = params["id"]
	id1, _ := strconv.Atoi(id)
	fmt.Println("Id is ", id1)
	count := DataAccess1.DeleteBook(id1)

	if count == 0 {
		io.WriteString(w, "Wrong id is passed")
	} else {
		io.WriteString(w, "Delete Request is sent")
	}
}
func main() {

	//Call Mux Package
	router := mux.NewRouter()
	fmt.Println("Start Web Application on Port Number :- localhost:8080/")
	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", GetBookById).Methods("GET")
	router.HandleFunc("/books", PostBook).Methods("POST")
	router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}
