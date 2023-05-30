package main

import (
  "database/sql"
  "fmt"

  _"github.com/go-sql-driver/mysql"
)

type mahasiswa struct {
  ID string `json:"id"`
  NIM  string `json:"nim"`
  Nama string `json:"nama"`
}

func main() {
  fmt.Println("Go MySQL Tutorial")

  // Open up our database connection.
  // I've set up a database on my local machine using phpmyadmin.
  // The database is called testDb
  db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/mysql")

  // if there is an error opening the connection, handle it
  if err != nil {
    panic(err.Error())
  }

  // defer the close till after the main function has finished
  // executing
  defer db.Close()

  // Execute the query
  results, err := db.Query("SELECT nim, nama FROM mahasiswa")
  if err != nil {
    panic(err.Error()) // proper error handling instead of panic in your app
  }

  for results.Next() {
    var mahasiswa mahasiswa
    // for each row, scan the result into our tag composite object
    err = results.Scan(&mahasiswa.NIM, &mahasiswa.Nama)
    if err != nil {
      panic(err.Error()) // proper error handling instead of panic in your app
    }
    // and then print out the mahasiswa's attribute
    fmt.Println(mahasiswa.NIM)
    fmt.Println(mahasiswa.Nama)
  }

}