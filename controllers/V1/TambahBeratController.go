package V1

import(
	"github.com/gin-gonic/gin"
    "Tes/sirclo_berat_badan/models"
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
	
)

type V1TambahBeratController struct {
	Status int
}

type BeratData struct {
	Id int `json:"id", gorm:"primary_key", gorm:autoIncrement` // id
	Tanggal time.Time `json:"Tanggal"`
	Max int `json:"max"` 
	Min int `json:"min"` 
	Perbedaan int `json:"perbedaan"` 
}


type ResponseBerat struct {
	Tanggal string
	Max int
	Min int
	Perbedaan int
}


func (status *V1TambahBeratController) TambahBerat (c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	Tanggal := c.DefaultPostForm("params[tanggal]", "")
	MaxBerat, err := strconv.Atoi(c.PostForm("params[max_berat]"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Max Berat Harus di isi ",
		})
		return
	}
	MinBerat, err := strconv.Atoi(c.PostForm("params[min_berat]"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Min Berat Harus Di Isi",
		})
		return
	}

	TanggalConvert, err := time.Parse("2006-01-02",Tanggal);
	// TanggalPeriod := TanggalConvert.AddDate(0, 0, 29)
    if err !=nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error parsing tanggal",
		})
		return
    }
	BeratDb, _ := db.Table("tb_berats").Where("tanggal = ?", TanggalConvert).Rows()
	defer BeratDb.Close()
    result := make([]BeratData, 0)
    for BeratDb.Next() {
        var row BeratData
        result = append(result, row)
		
    }
	if len(result) > 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Data Sudah Ada",
		})
		return
	}else{

		inputBerat := models.Tb_Berat{
			Tanggal : TanggalConvert,
			Max : MaxBerat,
			Min : MinBerat,
			Perbedaan : (MaxBerat - MinBerat),
		}
		db.Create(&inputBerat)

		response := ResponseBerat {
			Tanggal : TanggalConvert.Format("2006-01-02"),
			Max : MaxBerat,
			Min : MinBerat,
			Perbedaan : (MaxBerat - MinBerat),
		}
		fmt.Println(response)
		c.JSON(200, gin.H{"status": 200,"message" : "Berhasil Disimpan","response":response})
		return
	}
	
	

	
}


