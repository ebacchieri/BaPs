package sdk

import (
	"github.com/gin-gonic/gin"
	"github.com/ebacchieri/BaPs/gdconf"
)

func index(c *gin.Context) {
	c.JSON(200, gdconf.GetProdIndex())
}
