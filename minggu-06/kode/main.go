package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:<password>@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	type mahasiswa struct {
		Id         	int
		nim			string
		nama  		string
	}
	router := gin.Default()

	// GET a mahasiswa detail
	router.GET("/mahasiswa/:id", func(c *gin.Context) {
		var (
			mahasiswa Mahasiswa
			result gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("select id, nim, nama from mahasiswa where id = ?;", id)
		err = row.Scan(&mahasiswa.Id, &mahasiswa.nim, &mahasiswa.Nama)
		if err != nil {
			// If no results send null
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": mahasiswa,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})

	// GET all mahasiswa
	router.GET("/mahasiswa", func(c *gin.Context) {
		var (
			pmahasiswa Mahasiswa
			mahasiswa []Mahasiswa
		)
		rows, err := db.Query("select id, nim, nama from mahasiswa;")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&mahasiswa.Id, &mahasiswa.nim, &mahasiswa.nama)
			mahasiswa = append(mahasiswa, mahasiswa)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": mahasiswa,
			"count":  len(mahasiswa),
		})
	})

	// POST new mahasiswa details
	router.POST("/mahasiswa", func(c *gin.Context) {
		var buffer bytes.Buffer
		nim := c.PostForm("nim")
		nama := c.PostForm("nama")
		stmt, err := db.Prepare("insert into mahasiswa (nim, nama) values(?,?);")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(nim, nama)

		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(nim)
		buffer.WriteString(" ")
		buffer.WriteString(nama)
		defer stmt.Close()
		nama := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf(" %s successfully created", nama),
		})
	})

	// PUT - update a mahasiswa details
	router.PUT("/mahasiswa", func(c *gin.Context) {
		var buffer bytes.Buffer
		id := c.Query("id")
		nim := c.PostForm("nim")
		nama := c.PostForm("nama")
		stmt, err := db.Prepare("update mahasiswa set nim= ?, nama= ? where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(nim, nama, id)
		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(nim)
		buffer.WriteString(" ")
		buffer.WriteString(nama)
		defer stmt.Close()
		nama := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully updated to %s", nama),
		})
	})

	// Delete resources
	router.DELETE("/mahasiswa", func(c *gin.Context) {
		id := c.Query("id")
		stmt, err := db.Prepare("delete from mahasiswa where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(id)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted user: %s", id),
		})
	})
	router.Run(":3000")
}