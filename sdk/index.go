package sdk

import (
	"github.com/gin-gonic/gin"
	"github.com/gucooing/BaPs/gdconf"
)

func index(c *gin.Context) {
	c.JSON(200, gdconf.GetProdIndex())
}
