package controllers

import (
	"database/sql"
	"net/http"
	"time"

	"quiz/database"
	"quiz/models"

	"github.com/gin-gonic/gin"
)

func GetAllKategori(c *gin.Context) {
	rows, err := database.DB.Query(`SELECT id, name, created_at, created_by,
		COALESCE(modified_at, NOW()), COALESCE(modified_by, '') FROM kategori ORDER BY id`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Kategori
	for rows.Next() {
		var k models.Kategori
		if err := rows.Scan(&k.ID, &k.Name, &k.CreatedAt, &k.CreatedBy, &k.ModifiedAt, &k.ModifiedBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		list = append(list, k)
	}

	c.JSON(http.StatusOK, list)
}

func GetKategoriByID(c *gin.Context) {
	id := c.Param("id")
	var k models.Kategori

	query := `SELECT id, name, created_at, created_by, 
	          COALESCE(modified_at, NOW()), COALESCE(modified_by, '') FROM kategori WHERE id=$1`
	err := database.DB.QueryRow(query, id).Scan(&k.ID, &k.Name, &k.CreatedAt, &k.CreatedBy, &k.ModifiedAt, &k.ModifiedBy)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan!!"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, k)
}

func CreateKategori(c *gin.Context) {
	var kategori models.Kategori
	if err := c.ShouldBindJSON(&kategori); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	username := c.GetString("username")
	query := `INSERT INTO kategori (name, created_by, created_at)
	          VALUES ($1, $2, NOW())
	          RETURNING id, name, created_at, created_by, NOW(), ''`

	err := database.DB.QueryRow(query, kategori.Name, username).Scan(
		&kategori.ID,
		&kategori.Name,
		&kategori.CreatedAt,
		&kategori.CreatedBy,
		&kategori.ModifiedAt,
		&kategori.ModifiedBy,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, kategori)
}

func UpdateKategori(c *gin.Context) {
	id := c.Param("id")
	username := c.GetString("username")

	var input models.Kategori
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	query := `UPDATE kategori 
			  SET name=$1, modified_by=$2, modified_at=$3 
			  WHERE id=$4
			  RETURNING id, name, created_at, created_by, modified_at, modified_by`

	var updated models.Kategori
	err := database.DB.QueryRow(query, input.Name, username, time.Now(), id).Scan(
		&updated.ID, &updated.Name, &updated.CreatedAt, &updated.CreatedBy, &updated.ModifiedAt, &updated.ModifiedBy)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan!!"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func DeleteKategori(c *gin.Context) {
	id := c.Param("id")
	result, err := database.DB.Exec("DELETE FROM kategori WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan!!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kategori deleted"})
}
