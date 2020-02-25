package appProvider

import (
	"github.com/gin-gonic/gin"
	"github.com/koushamad/goro-cache/app/cache"
)

func Boot(egn *gin.Engine) {
	cache.Boot()
}

func Kill()  {

}