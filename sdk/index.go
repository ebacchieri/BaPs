package sdk

import (
	"github.com/gin-gonic/gin"
	"./gdconf"
)

func index(c *gin.Context) {
	c.JSON(200, gdconf.GetProdIndex())
}
