package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	suffixes [5]string
)

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func HumanFileSize(size float64) string {
	fmt.Println(size)
	suffixes[0] = "B"
	suffixes[1] = "KB"
	suffixes[2] = "MB"
	suffixes[3] = "GB"
	suffixes[4] = "TB"

	base := math.Log(size) / math.Log(1024) //file.size = 1901
	fmt.Println("base :", base)             //1.0892542816648552
	getSize := Round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	fmt.Println("getSize :", getSize)
	fmt.Println(int(math.Floor(base)))
	getSuffix := suffixes[int(math.Floor(base))]
	fmt.Println("getSuffix :", getSuffix)
	return strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix)
}

func main() {

	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB // 8 x 2^20 // 2 ^ (10+10)
	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		// Source
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}

		fs := float64(file.Size)
		fmt.Println("float64 file size:: ", fs)
		hrfs := HumanFileSize(fs)
		fmt.Println("Human Readable file size::", hrfs)

		filename := file.Filename
		fmt.Println("file name is ::", filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}

		c.String(http.StatusOK, "File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email)
	})
	//server running on port
	//hello server
	router.Run(":8080")

}
