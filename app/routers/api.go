package routers

import (
	"fmt"
	"mirae-va/app/handlers"
	"mirae-va/config"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	healthHandler := handlers.NewHealthHandler()

	vaTrxhandler := handlers.NewVirtualAccountTransactionHandler()

	api := r.Group(fmt.Sprintf("/%s", config.SERVICE_ROOT_PATH))

	rootPrefix := api.Group("/v1")

	rootPrefix.GET("health", healthHandler.Health)

	rootPrefix.POST("generate", vaTrxhandler.BillGenerate)
	rootPrefix.POST("inquiry", vaTrxhandler.Inquiry)
	rootPrefix.POST("payment", vaTrxhandler.Payment)
	rootPrefix.POST("report", vaTrxhandler.Report)

}
