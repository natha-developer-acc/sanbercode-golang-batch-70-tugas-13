package main

import (
	"net/http"
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

	// POST /bioskop
	r.POST("/bioskop", func(c *gin.Context) {
		var input Bioskop

		// bind JSON
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// validasi
		if input.Nama == "" || input.Lokasi == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
			return
		}

		// simpan ke DB
		bioskop := Bioskop{
			Nama:   input.Nama,
			Lokasi: input.Lokasi,
			Rating: input.Rating,
		}

		if err := database.DB.Create(&bioskop).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil ditambahkan!", "data": bioskop})
	})

	// GET /bioskop untuk lihat data
	r.GET("/bioskop", func(c *gin.Context) {
		var bioskops []Bioskop
		if err := database.DB.Find(&bioskops).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": bioskops})
	})

	r.Run(":8080")
}
