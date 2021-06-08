package V1

import(
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	"time"
	
)

type V1IndexBeratController struct {
	Status int
}

type ResponAverageBerat struct {
	Max int
	Min int
	Perbedaan int
}

type ResponIndex struct {
	BeratBadan []ResponseDetailBerat
	RataRataBerat ResponAverageBerat
}

func (status *V1IndexBeratController) IndexBerat (c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	
	BeratDb, _ := db.Table("tb_berats").Select("tanggal, max, min, perbedaan").Order("tanggal desc").Rows()
	defer BeratDb.Close()
    
	resultBeratBadan := ResponseDetailBerat{}
	result := []ResponseDetailBerat{}
	var maxTotal, minTotal, jumlahData int
    for BeratDb.Next() {
        var max, min, perbedaan int
		var tanggal time.Time
		BeratDb.Scan(&tanggal, &max, &min, &perbedaan)
		resultBeratBadan.Tanggal = tanggal.Format("2006-01-02")
        resultBeratBadan.Max = max
        resultBeratBadan.Min = min
		resultBeratBadan.Perbedaan = perbedaan
		maxTotal += max
		minTotal += min
		jumlahData +=1
        result = append(result, resultBeratBadan)
    }
	if len(result) > 0 {
		fmt.Println(result)
		averageBerat := ResponAverageBerat{
			Max :maxTotal/jumlahData,
			Min :minTotal/jumlahData, 
			Perbedaan :((maxTotal/jumlahData)-(minTotal/jumlahData)),
		}
		responIndex := ResponIndex{
			BeratBadan :result,
			RataRataBerat: averageBerat,
		}

		c.JSON(200, gin.H{"status": 200,"response":responIndex})
		return
	}else{
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Data Tidak di Temukan",
		})
		return
	}
}


