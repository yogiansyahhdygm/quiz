package controllers

import (
	"database/sql"
	"net/http"
	"time"

	"quiz/database"
	"quiz/models"

	"github.com/gin-gonic/gin"
)

// GET /api/buku
func GetAllBuku(c *gin.Context) {
	rows, err := database.DB.Query(`
		SELECT id, title, description, image_url, release_year, price, total_page, thickness,
		       category_id, created_at, created_by,
		       COALESCE(modified_at, NOW()), COALESCE(modified_by, '')
		FROM buku ORDER BY id`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Buku
	for rows.Next() {
		var b models.Buku
		if err := rows.Scan(
			&b.ID, &b.Title, &b.Description, &b.ImageURL, &b.ReleaseYear,
			&b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID,
			&b.CreatedAt, &b.CreatedBy, &b.ModifiedAt, &b.ModifiedBy,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		list = append(list, b)
	}

	c.JSON(http.StatusOK, list)
}

// GET /api/buku/:id
func GetBukuByID(c *gin.Context) {
	id := c.Param("id")
	var b models.Buku

	query := `SELECT id, title, description, image_url, release_year, price, total_page, thickness,
	          category_id, created_at, created_by,
	          COALESCE(modified_at, NOW()), COALESCE(modified_by, '')
	          FROM buku WHERE id=$1`

	err := database.DB.QueryRow(query, id).Scan(
		&b.ID, &b.Title, &b.Description, &b.ImageURL, &b.ReleaseYear,
		&b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID,
		&b.CreatedAt, &b.CreatedBy, &b.ModifiedAt, &b.ModifiedBy,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan!!"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, b)
}

// POST /api/buku
func CreateBuku(c *gin.Context) {
	var b models.Buku
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	username := c.GetString("username")

	// Validasi release_year
	if b.ReleaseYear < 1980 || b.ReleaseYear > 2025 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Release year harus antara 1980 and 2025"})
		return
	}

	// Konversi otomatis thickness
	if b.TotalPage > 100 {
		b.Thickness = "tebal"
	} else {
		b.Thickness = "tipis"
	}

	query := `
		INSERT INTO buku (title, description, image_url, release_year, price, total_page, thickness, category_id, created_by, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,NOW())
		RETURNING id, title, description, image_url, release_year, price, total_page, thickness,
		          category_id, created_at, created_by, NOW(), ''`

	err := database.DB.QueryRow(query,
		b.Title, b.Description, b.ImageURL, b.ReleaseYear, b.Price,
		b.TotalPage, b.Thickness, b.CategoryID, username,
	).Scan(
		&b.ID, &b.Title, &b.Description, &b.ImageURL, &b.ReleaseYear,
		&b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID,
		&b.CreatedAt, &b.CreatedBy, &b.ModifiedAt, &b.ModifiedBy,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, b)
}

// PUT /api/buku/:id
func UpdateBuku(c *gin.Context) {
	id := c.Param("id")
	username := c.GetString("username")

	var input models.Buku
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validasi release_year
	if input.ReleaseYear < 1980 || input.ReleaseYear > 2025 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Release year must be between 1980 and 2025"})
		return
	}

	// Konversi otomatis thickness
	if input.TotalPage > 100 {
		input.Thickness = "tebal"
	} else {
		input.Thickness = "tipis"
	}

	query := `
		UPDATE buku
		SET title=$1, description=$2, image_url=$3, release_year=$4, price=$5,
		    total_page=$6, thickness=$7, category_id=$8,
		    modified_by=$9, modified_at=$10
		WHERE id=$11
		RETURNING id, title, description, image_url, release_year, price, total_page, thickness,
		          category_id, created_at, created_by, modified_at, modified_by`

	var updated models.Buku
	err := database.DB.QueryRow(query,
		input.Title, input.Description, input.ImageURL, input.ReleaseYear, input.Price,
		input.TotalPage, input.Thickness, input.CategoryID,
		username, time.Now(), id,
	).Scan(
		&updated.ID, &updated.Title, &updated.Description, &updated.ImageURL, &updated.ReleaseYear,
		&updated.Price, &updated.TotalPage, &updated.Thickness, &updated.CategoryID,
		&updated.CreatedAt, &updated.CreatedBy, &updated.ModifiedAt, &updated.ModifiedBy,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan!!"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

// DELETE /api/buku/:id
func DeleteBuku(c *gin.Context) {
	id := c.Param("id")
	result, err := database.DB.Exec("DELETE FROM buku WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan!!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Buku telah terhapus"})
}

// GET buku by kategori
func GetBukuByKategori(c *gin.Context) {
	categoryID := c.Param("id")

	query := `SELECT 
		b.id, b.title, b.description, b.image_url, 
		b.release_year, b.price, b.total_page, b.thickness, 
		b.category_id, COALESCE(k.name, '') AS category_name,
		b.created_at, b.created_by, 
		COALESCE(b.modified_at, NOW()) AS modified_at,
		COALESCE(b.modified_by, '') AS modified_by
	FROM buku b
	JOIN kategori k ON b.category_id = k.id
	WHERE b.category_id = $1
	ORDER BY b.id`

	rows, err := database.DB.Query(query, categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	type BukuWithCategory struct {
		models.Buku
		CategoryName string `json:"category_name"`
	}

	var list []BukuWithCategory

	for rows.Next() {
		var b BukuWithCategory
		err := rows.Scan(
			&b.ID,
			&b.Title,
			&b.Description,
			&b.ImageURL,
			&b.ReleaseYear,
			&b.Price,
			&b.TotalPage,
			&b.Thickness,
			&b.CategoryID,
			&b.CategoryName,
			&b.CreatedAt,
			&b.CreatedBy,
			&b.ModifiedAt,
			&b.ModifiedBy,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		list = append(list, b)
	}

	if len(list) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Tidak ada buku untuk kategori ini"})
		return
	}

	c.JSON(http.StatusOK, list)
}
