package routers

import (
	"github.com/gin-gonic/gin"
	"Tes/sirclo_berat_badan/config"
	"net/http"
	"Tes/sirclo_berat_badan/controllers/V1"
)

func RouterMain() http.Handler  {
	router := gin.New()
	db := config.SetupModels()

	router.Use(func(c * gin.Context){
		c.Set("db",db)
		c.Next()
	})
	v1TambahBerat := &V1.V1TambahBeratController{Status: 200}
	v1UpdateBerat := &V1.V1UpdateBeratController{Status: 200}
	v1HapusBerat := &V1.V1HapusBeratController{Status: 200}
	v1DetailBerat := &V1.V1DetailBeratController{Status: 200}
	v1IndexBerat := &V1.V1IndexBeratController{Status: 200}
	
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"data":"Welcome To Api Golang"})
	})

	router.POST("/tambah_berat", v1TambahBerat.TambahBerat)
	router.POST("/update_berat", v1UpdateBerat.UpdateBerat)
	router.POST("/hapus_berat", v1HapusBerat.HapusBerat)
	router.POST("/detail_berat", v1DetailBerat.DetailBerat)
	router.POST("/index_berat", v1IndexBerat.IndexBerat)
	
	
	return router
}
