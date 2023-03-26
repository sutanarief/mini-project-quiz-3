package controllers

import (
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetHasilSegitiga(c *gin.Context) {
	var (
		result gin.H
	)

	queryParams := c.Request.URL.Query()
	hitung := strings.ToLower(queryParams["hitung"][0])
	alas, _ := strconv.Atoi(queryParams["alas"][0])
	tinggi, _ := strconv.Atoi(queryParams["tinggi"][0])

	if hitung == "luas" {
		result = gin.H{
			"result": "Luas dari segitiga sama sisi dengan alas:" + strconv.Itoa(alas) + " dan tinggi:" + strconv.Itoa(tinggi) + " adalah : " + strconv.Itoa(alas*tinggi/2),
		}
	} else {
		result = gin.H{
			"result": "Keliling dari segitiga sama sisi dengan alas:" + strconv.Itoa(alas) + " dan tinggi:" + strconv.Itoa(tinggi) + " adalah : " + strconv.Itoa(alas*3),
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetHasilPersegi(c *gin.Context) {
	var (
		result gin.H
	)
	queryParams := c.Request.URL.Query()
	hitung := strings.ToLower(queryParams["hitung"][0])
	sisi, _ := strconv.Atoi(queryParams["sisi"][0])

	if hitung == "luas" {
		result = gin.H{
			"result": "Luas dari dengan sisi:" + strconv.Itoa(sisi) + " adalah : " + strconv.Itoa(sisi*sisi),
		}
	} else {
		result = gin.H{
			"result": "Keliling dari dengan sisi:" + strconv.Itoa(sisi) + " adalah : " + strconv.Itoa(sisi*4),
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetHasilPersegiPanjang(c *gin.Context) {
	var (
		result gin.H
	)
	queryParams := c.Request.URL.Query()
	hitung := strings.ToLower(queryParams["hitung"][0])
	panjang, _ := strconv.Atoi(queryParams["panjang"][0])
	lebar, _ := strconv.Atoi(queryParams["lebar"][0])

	if hitung == "luas" {
		result = gin.H{
			"result": "Luas dari dengan panjang:" + strconv.Itoa(panjang) + " dan lebar:" + strconv.Itoa(lebar) + " adalah : " + strconv.Itoa(panjang*lebar),
		}
	} else {
		result = gin.H{
			"result": "Keliling dari dengan panjang:" + strconv.Itoa(panjang) + " dan lebar:" + strconv.Itoa(lebar) + " adalah : " + strconv.Itoa(2*(panjang+lebar)),
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetHasilLingkaran(c *gin.Context) {
	var (
		result gin.H
	)
	queryParams := c.Request.URL.Query()
	hitung := strings.ToLower(queryParams["hitung"][0])
	jariJari, _ := strconv.ParseFloat(queryParams["jariJari"][0], 64)

	if hitung == "luas" {
		result = gin.H{
			"result": "Luas dari dengan jari-jari:" + strconv.FormatFloat(jariJari, 'f', 0, 64) + " adalah : " + strconv.FormatFloat(math.Pi*math.Pow(jariJari, 2), 'f', 6, 64),
		}
	} else {
		result = gin.H{
			"result": "Keliling dari dengan jari-jari:" + strconv.FormatFloat(jariJari, 'f', 0, 64) + " adalah : " + strconv.FormatFloat(2*math.Pi*jariJari, 'f', 6, 64),
		}
	}

	c.JSON(http.StatusOK, result)
}

// func GetBangunDatar(c *gin.Context) {
// 	var (
// 		result gin.H
// 	)

// 	jenis := c.Param("jenis")
// 	queryParams := c.Request.URL.Query()

// 	switch jenis {
// 	case "segitiga-sama-sisi":
// 		hasilHitung := segitigaHandler(queryParams, jenis)
// 		result = gin.H{
// 			"result": hasilHitung,
// 		}
// 	case "persegi":
// 		hasilHitung := persegiHandler(queryParams, jenis)
// 		result = gin.H{
// 			"result": hasilHitung,
// 		}

// 	case "persegi-panjang":
// 		hasilHitung := persegiPanjangHandler(queryParams, jenis)
// 		result = gin.H{
// 			"result": hasilHitung,
// 		}

// 	case "lingkaran":
// 		hasilHitung := lingkaranHandler(queryParams, jenis)
// 		result = gin.H{
// 			"result": hasilHitung,
// 		}

// 	default:
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"Not Found": "Jenis Bangun datar tidak ditemukan",
// 		})
// 	}

// 	c.JSON(http.StatusOK, result)
// }

// func segitigaHandler(queryParams map[string][]string, jenis string) string {
// 	hitung := strings.ToLower(queryParams["hitung"][0])
// 	alas, _ := strconv.Atoi(queryParams["alas"][0])
// 	tinggi, _ := strconv.Atoi(queryParams["tinggi"][0])
// 	jenisBangunDatar := strings.Join(strings.Split(jenis, "-"), " ")

// 	if hitung == "luas" {
// 		return "Luas dari " + jenisBangunDatar + " dengan alas:" + strconv.Itoa(alas) + " dan tinggi:" + strconv.Itoa(tinggi) + " adalah : " + strconv.Itoa(alas*tinggi/2)
// 	} else {
// 		return "Keliling dari " + jenisBangunDatar + " dengan alas:" + strconv.Itoa(alas) + " dan tinggi:" + strconv.Itoa(tinggi) + " adalah : " + strconv.Itoa(alas*3)
// 	}
// }

// func persegiHandler(queryParams map[string][]string, jenis string) string {
// 	hitung := strings.ToLower(queryParams["hitung"][0])
// 	sisi, _ := strconv.Atoi(queryParams["sisi"][0])
// 	jenisBangunDatar := strings.Join(strings.Split(jenis, "-"), " ")

// 	if hitung == "luas" {
// 		return "Luas dari " + jenisBangunDatar + " dengan sisi:" + strconv.Itoa(sisi) + " adalah : " + strconv.Itoa(sisi*sisi)
// 	} else {
// 		return "Keliling dari " + jenisBangunDatar + " dengan sisi:" + strconv.Itoa(sisi) + " adalah : " + strconv.Itoa(sisi*4)
// 	}
// }

// func persegiPanjangHandler(queryParams map[string][]string, jenis string) string {
// 	hitung := strings.ToLower(queryParams["hitung"][0])
// 	panjang, _ := strconv.Atoi(queryParams["panjang"][0])
// 	lebar, _ := strconv.Atoi(queryParams["lebar"][0])
// 	jenisBangunDatar := strings.Join(strings.Split(jenis, "-"), " ")

// 	if hitung == "luas" {
// 		return "Luas dari " + jenisBangunDatar + " dengan panjang:" + strconv.Itoa(panjang) + " dan lebar:" + strconv.Itoa(lebar) + " adalah : " + strconv.Itoa(panjang*lebar)
// 	} else {
// 		return "Keliling dari " + jenisBangunDatar + " dengan panjang:" + strconv.Itoa(panjang) + " dan lebar:" + strconv.Itoa(lebar) + " adalah : " + strconv.Itoa(2*(panjang+lebar))
// 	}
// }

// func lingkaranHandler(queryParams map[string][]string, jenis string) string {
// 	hitung := strings.ToLower(queryParams["hitung"][0])
// 	jariJari, _ := strconv.ParseFloat(queryParams["jariJari"][0], 64)
// 	jenisBangunDatar := strings.Join(strings.Split(jenis, "-"), " ")

// 	if hitung == "luas" {
// 		return "Luas dari " + jenisBangunDatar + " dengan jari-jari:" + strconv.FormatFloat(jariJari, 'f', 0, 64) + " adalah : " + strconv.FormatFloat(math.Pi*math.Pow(jariJari, 2), 'f', 6, 64)
// 	} else {
// 		return "Keliling dari " + jenisBangunDatar + " dengan jari-jari:" + strconv.FormatFloat(jariJari, 'f', 0, 64) + " adalah : " + strconv.FormatFloat(2*math.Pi*jariJari, 'f', 6, 64)
// 	}
// }
