package V1

import(
	"github.com/gin-gonic/gin"
    "Tes/sirclo_berat_badan/models"
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	"time"
	
)

type V1HapusBeratController struct {
	Status int
}

type ResponseHapusBerat struct {
	Tanggal string
}



func (status *V1HapusBeratController) HapusBerat (c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	Tanggal := c.DefaultPostForm("params[tanggal]", "")
	
	TanggalConvert, err := time.Parse("2006-01-02",Tanggal);
    if err !=nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error parsing tanggal",
		})
		return
    }
	BeratDb, _ := db.Table("tb_berats").Where("tanggal = ?", TanggalConvert).Rows()
	defer BeratDb.Close()
    result := make([]BeratData, 0)
	var resultIdBeratBadan ResponseIdBerat
    for BeratDb.Next() {
		db.ScanRows(BeratDb, &resultIdBeratBadan)
        var row BeratData
        result = append(result, row)
    }

	if len(result) > 0 {
		var beratkModels models.Tb_Berat
		db.Where("id = ?", resultIdBeratBadan.Id).Delete(&beratkModels)
		response := ResponseHapusBerat {
			Tanggal : TanggalConvert.Format("2006-01-02"),
		}
		fmt.Println(response)
		c.JSON(200, gin.H{"status": 200,"message" : "Berhasil Di hapus","response":response})
		return
	}else{
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Data Tidak di Temukan, atau data sudah tidak ada",
		})
		return
	}
}


