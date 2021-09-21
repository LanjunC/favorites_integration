package v1

import (
	"codingcrea/favorites_integration/model"
	"codingcrea/favorites_integration/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"strconv"
)

// GetBiliFavors 调bili api获取我的所有收藏夹-created
func GetBiliFavors(ctx *gin.Context) {
	resp, err := http.Get(model.GetBiliFavors)
	if err != nil {
		glog.Error(err)
		utils.CtxJson(ctx, utils.CodeBiliFavorsErr)
		return
	}
	defer resp.Body.Close()
	respJson, err := ioutil.ReadAll(resp.Body)
	biliFavors := model.BiliFavors{}
	_ = json.Unmarshal(respJson, &biliFavors)

	if biliFavors.Code != 0 {
		glog.Errorf("bilibili api`s resp code is not 0, code:%v, message:%v", biliFavors.Code, biliFavors.Message)
		utils.CtxJson(ctx, utils.CodeBiliFavorsErr)
		return
	}
	biliFavors.Message = "" //Code和Message都设为0值，不传给前端（使用我们自己的code和message）
	utils.CtxJsonOfData(ctx, utils.CodeSuccess, "bili_favors", biliFavors)
}

// GetBiliFavorItemsByMid 调bili api通过media_id获取bili收藏夹所有items
func GetBiliFavorItemsByMid(ctx *gin.Context) {
	mediaId := ctx.Param("media_id")
	pn := ctx.DefaultQuery("pn", "1")
	ps := ctx.DefaultQuery("ps", "10")
	if mediaIdInt, err := strconv.Atoi(mediaId); err != nil || mediaIdInt <= 0 {
		glog.Error("Invalid \"media_id\":%v", mediaId)
		utils.CtxJsonOfData(ctx, utils.CodeZhihuInvalidParam, "media_id", mediaId)
		return
	}
	if pnInt, err := strconv.Atoi(pn); err != nil || pnInt < 0 { //pn为0或1都是第一页
		glog.Error("Invalid \"pn\":%v", pn)
		utils.CtxJsonOfData(ctx, utils.CodeZhihuInvalidParam, "pn", pn)
		return
	}
	if psInt, err := strconv.Atoi(ps); err != nil || psInt <= 0 {
		glog.Error("Invalid \"ps\":%v", ps)
		utils.CtxJsonOfData(ctx, utils.CodeZhihuInvalidParam, "ps", ps)
		return
	}

	//调bili api通过media_id获取bili收藏夹所有items
	resp, err := http.Get(fmt.Sprintf(model.GetBiliFavorItemsByMidApi, mediaId, pn, ps))
	if err != nil {
		glog.Error(err)
		utils.CtxJson(ctx, utils.CodeZhihuFavorItemsErr)
		return
	}
	defer resp.Body.Close()
	respJson, err := ioutil.ReadAll(resp.Body)
	favorItems := model.BiliFavorItems{}
	_ = json.Unmarshal(respJson, &favorItems)
	if favorItems.Code != 0 {
		glog.Errorf("bilibili api`s resp code is not 0, code:%v, message:%v", favorItems.Code, favorItems.Message)
		utils.CtxJson(ctx, utils.CodeBiliFavorsErr)
		return
	}
	favorItems.Message = ""
	if favorItems.Data == nil {
		glog.Error("not found, \"media_id\":%v", mediaId)
		utils.CtxJsonOfData(ctx, utils.CodeBiliNotFoundErr, "media_id", mediaId)
		return
	}

	utils.CtxJsonOfData(ctx, utils.CodeSuccess, "bili_favoritems", favorItems)
}
