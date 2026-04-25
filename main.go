package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	connStr := "postgres://postgres:mmdh1334G!@localhost:5432/postgres?sslmode=disable"

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/albums", func(c *gin.Context) {
		rows, err := db.Query(`
			SELECT id, title, artist, price
			FROM albums
			ORDER BY id
		`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var albums []Album
		for rows.Next() {
			var a Album
			if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			albums = append(albums, a)
		}

		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, albums)
	})

	r.POST("/albums", func(c *gin.Context) {
		var newAlbum Album

		if err := c.ShouldBindJSON(&newAlbum); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := db.Exec(`
			INSERT INTO albums (id, title, artist, price)
			VALUES ($1, $2, $3, $4)
		`, newAlbum.ID, newAlbum.Title, newAlbum.Artist, newAlbum.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, newAlbum)
	})

	r.Run(":8080")
}
