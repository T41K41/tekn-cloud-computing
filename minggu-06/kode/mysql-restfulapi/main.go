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
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	type Mahasiswa struct {
		ID   int
		NIM  string
		Nama string
	}

	router := gin.Default()

	router.GET("/mahasiswa/:id", func(c *gin.Context) {
		var (
			mahasiswa Mahasiswa
			result    gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("SELECT id, nim, nama FROM mahasiswa WHERE id = ?;", id)
		err = row.Scan(&mahasiswa.ID, &mahasiswa.NIM, &mahasiswa.Nama)
		if err != nil {
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

	router.GET("/mahasiswa", func(c *gin.Context) {
		var (
			mahasiswa Mahasiswa
			mahasiswaList []Mahasiswa
		)
		rows, err := db.Query("SELECT id, nim, nama FROM mahasiswa;")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&mahasiswa.ID, &mahasiswa.NIM, &mahasiswa.Nama)
			mahasiswaList = append(mahasiswaList, mahasiswa)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": mahasiswaList,
			"count":  len(mahasiswaList),
		})
	})

	router.POST("/mahasiswa", func(c *gin.Context) {
		var buffer bytes.Buffer
		nim := c.PostForm("nim")
		nama := c.PostForm("nama")
		stmt, err := db.Prepare("INSERT INTO mahasiswa (nim, nama) VALUES (?,?);")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(nim, nama)

		if err != nil {
			fmt.Print(err.Error())
		}

		buffer.WriteString(nim)
		buffer.WriteString(" ")
		buffer.WriteString(nama)
		defer stmt.Close()
		namaString := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%s successfully created", namaString),
		})
	})

	router.PUT("/mahasiswa", func(c *gin.Context) {
		var buffer bytes.Buffer
		id := c.Query("id")
		nim := c.PostForm("nim")
		nama := c.PostForm("nama")
		stmt, err := db.Prepare("UPDATE mahasiswa SET nim=?, nama=? WHERE id=?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(nim, nama, id)
		if err != nil {
			fmt.Print(err.Error())
		}

		buffer.WriteString(nim)
		buffer.WriteString(" ")
		buffer.WriteString(nama)
		defer stmt.Close()
		namaString := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully updated to %s", namaString),
		})
	})

	router.DELETE("/mahasiswa", func(c *gin.Context) {
		id := c.Query("id")
		stmt, err := db.Prepare("DELETE FROM mahasiswa WHERE id=?;")
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
