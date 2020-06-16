package DataAccess1

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	Id     int
	Name   string
	Isbn   string
	Pages  int
	author string
}

func connect() *sql.DB {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sample")

	if err != nil {

		log.Fatal(err)

	}

	return db

}

func DeleteBook(id int) (count int64) {
	db := connect()
	defer db.Close()
	fmt.Println("Delete book connected", id)

	delForm, err := db.Prepare("DELETE FROM Books WHERE ID=?")
	if err != nil {
		panic(err.Error())
	}
	//	delForm.Exec(id)
	res, _ := delForm.Exec(id)

	values, _ := res.RowsAffected()
	fmt.Println(values)
	return values

}

func UpdateBook(book Book) {

	var db = connect()
	defer db.Close()
	insForm, err := db.Prepare("UPDATE Books SET name=?, isbn=? WHERE id=?")
	if err != nil {
		log.Print(err)
	}
	_, err1 := insForm.Exec(book.Name, book.Isbn, book.Id)

	if err1 != nil {
		log.Print(err1)
	}

}
func InsertBook(book Book) {

	db := connect()
	defer db.Close()

	stment, err1 := db.Prepare("Insert into books(name,isbn,author,pages)values(?,?,?,?)")
	if err1 != nil {
		fmt.Println(err1)
	}
	res, err := stment.Exec(book.Name, book.Isbn, book.author, book.Pages)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

}

func GetBookById(id int) Book {
	db := connect()
	defer db.Close()
	var book Book

	row := db.QueryRow("select * from books where ID= ?;", id)

	err := row.Scan(&book.Id, &book.Name, &book.Isbn, &book.author, &book.Pages)

	if err != nil {
		fmt.Println(err)
	}

	return book

}

func GetBooks() []Book {
	var books []Book
	var book Book
	db := connect()
	defer db.Close()
	rows, err := db.Query("select Id,Name,Isbn,author,Pages from Books")
	if err != nil {
		fmt.Println(err)
	} else {

		for rows.Next() {
			err := rows.Scan(&book.Id, &book.Name, &book.Isbn, &book.author, &book.Pages)
			if err != nil {
				fmt.Println(err)
			}
			books = append(books, book)
		}
	}
	return books
}
