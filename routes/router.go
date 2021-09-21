package routes

import (
	v1 "codingcrea/favorites_integration/api/v1"
	"codingcrea/favorites_integration/utils"
)
import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	// gin模式为debug时可打印更多信息 debug/release
	gin.SetMode(utils.AppMode)
	e := gin.Default()

	groupV1 := e.Group("api/v1")
	{
		groupV1.GET("test", v1.Test)
		groupV1Zhihu := groupV1.Group("zhihu")
		{
			groupV1Zhihu.GET("favors", v1.GetZhihuFavors)
			groupV1Zhihu.GET("favor-items/:cid", v1.GetZhihuFavorItemsByCid)
		}

		groupV1Bili := groupV1.Group("bili")
		{
			groupV1Bili.GET("favors", v1.GetBiliFavors)
			groupV1Bili.GET("favor-items/:media_id", v1.GetBiliFavorItemsByMid)
		}

		groupV1Pinbox := groupV1.Group("pinbox")
		{
			groupV1Pinbox.GET("favors", v1.GetPinboxFavors)
			groupV1Pinbox.GET("favor-items/:cid", v1.GetPinboxFavorItemsByCid)
		}
	}
	return e
}
