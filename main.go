package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"sanbercode-golang-batch-70-tugas-13/database"
)

type Bioskop struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Nama   string  `json:"nama"`
	Lokasi string  `json:"lokasi"`
	Rating float64 `json:"rating"`
}

func main() {
	// Koneksi ke DB
	database.ConnectDB()

	// Auto migrate model ke DB
	database.DB.AutoMigrate(&Bioskop{})

	r := gin.Default()

	// ============= CREATE (Tugas 13) =============
	r.POST("/bioskop", func(c *gin.Context) {
		var input Bioskop
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if input.Nama == "" || input.Lokasi == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
			return
		}
		bioskop := Bioskop{Nama: input.Nama, Lokasi: input.Lokasi, Rating: input.Rating}
		if err := database.DB.Create(&bioskop).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil ditambahkan!", "data": bioskop})
	})

	// ============= READ (Tugas 14) =============
	// GET semua bioskop
	r.GET("/bioskop", func(c *gin.Context) {
		var bioskops []Bioskop
		if err := database.DB.Find(&bioskops).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": bioskops})
	})

	// GET detail bioskop by ID
	r.GET("/bioskop/:id", func(c *gin.Context) {
		id := c.Param("id")
		var bioskop Bioskop
		if err := database.DB.First(&bioskop, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": bioskop})
	})

	// ============= UPDATE (Tugas 14) =============
	r.PUT("/bioskop/:id", func(c *gin.Context) {
		id := c.Param("id")
		var bioskop Bioskop

		// cek data ada atau tidak
		if err := database.DB.First(&bioskop, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
			return
		}

		// bind JSON
		var input Bioskop
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if input.Nama == "" || input.Lokasi == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
			return
		}

		// update data
		bioskop.Nama = input.Nama
		bioskop.Lokasi = input.Lokasi
		bioskop.Rating = input.Rating

		if err := database.DB.Save(&bioskop).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil diperbarui!", "data": bioskop})
	})

	// ============= DELETE (Tugas 14) =============
	r.DELETE("/bioskop/:id", func(c *gin.Context) {
		id := c.Param("id")

		// pastikan ID valid angka
		if _, err := strconv.Atoi(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
			return
		}

		// hapus data
		if err := database.DB.Delete(&Bioskop{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil dihapus"})
	})

	// jalankan server
	r.Run(":8080")
}
