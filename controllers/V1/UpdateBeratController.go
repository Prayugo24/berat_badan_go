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

type V1UpdateBeratController struct {
	Status int
}

type ResponseIdBerat struct {
	Id   int 
}



func (status *V1UpdateBeratController) UpdateBerat (c *gin.Context){
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
		perbedaan := (MaxBerat - MinBerat)
		
		db.Model(&beratkModels).Where("id = ?", resultIdBeratBadan.Id).Updates(map[string]interface{}{"max": MaxBerat, "min": MinBerat, "perbedaan": perbedaan})

		response := ResponseBerat {
			Tanggal : TanggalConvert.Format("2006-01-02"),
			Max : MaxBerat,
			Min : MinBerat,
			Perbedaan : perbedaan,
		}
		fmt.Println(response)
		c.JSON(200, gin.H{"status": 200,"message" : "Berhasil Di ubah","response":response})
		return
	}else{
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Data Tidak Ditemukan",
		})
		return
	}
	
	

	
}


