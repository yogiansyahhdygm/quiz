package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBuku(c *gin.Context) {
	username := c.GetString("username")
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil akses data buku",
		"user":    username,
	})
}

func GetKategori(c *gin.Context) {
	username := c.GetString("username")
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil akses data kategori",
		"user":    username,
	})
}
