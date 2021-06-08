package V1

import(
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	"time"
	
)

type V1DetailBeratController struct {
	Status int
}
type ResponseDetailBerat struct {
	Tanggal string
	Max int
	Min int
	Perbedaan int
}

func (status *V1DetailBeratController) DetailBerat (c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	Tanggal := c.DefaultPostForm("params[tanggal]", "")
	
	TanggalConvert, err := time.Parse("2006-01-02",Tanggal);
    if err !=nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error parsing tanggal",
		})
		return
    }
	BeratDb, _ := db.Table("tb_berats").Select("tanggal, max, min, perbedaan").Where("tanggal = ?", TanggalConvert).Rows()
	defer BeratDb.Close()
    
	resultBeratBadan := ResponseDetailBerat{}
	result := []ResponseDetailBerat{}
    for BeratDb.Next() {
        var max, min, perbedaan int
		var tanggal time.Time
		BeratDb.Scan(&tanggal, &max, &min, &perbedaan)
		resultBeratBadan.Tanggal = tanggal.Format("2006-01-02")
        resultBeratBadan.Max = max
        resultBeratBadan.Min = min
		resultBeratBadan.Perbedaan = perbedaan
        result = append(result, resultBeratBadan)
    }

	if len(result) > 0 {
		fmt.Println(result)
		c.JSON(200, gin.H{"status": 200,"response":result})
		return
	}else{
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Data Tidak di Temukan",
		})
		return
	}
}


