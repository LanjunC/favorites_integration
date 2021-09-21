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

// GetPinboxFavors 获取pinbox所有收藏夹信息
func GetPinboxFavors(ctx *gin.Context) {
	req, _ := http.NewRequest("GET", model.GetPinboxFavorsApi, nil)
	req.Header.Add("authorization", utils.PinboxAuth)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		glog.Error(err)
		utils.CtxJson(ctx, utils.CodePinboxFavorsErr)
		return
	}
	if resp.StatusCode != http.StatusOK {
		glog.Errorf("the api may changed(or other reason), httpcode:%v", resp.StatusCode)
		utils.CtxJson(ctx, utils.CodePinboxFavorsErr)
		return
	}

	defer resp.Body.Close()
	respJson, err := ioutil.ReadAll(resp.Body)
	pinboxFavors := model.PinboxFavors{}
	_ = json.Unmarshal(respJson, &pinboxFavors)

	utils.CtxJsonOfData(ctx, utils.CodeSuccess, "pinbox_favors", pinboxFavors)
}

// GetPinboxFavorItemsByCid 通过cid获取该收藏夹所有items
func GetPinboxFavorItemsByCid(ctx *gin.Context) {
	cid := ctx.Param("cid")
	count := ctx.DefaultQuery("count", "10")
	if cidInt, err := strconv.Atoi(cid); err != nil || cidInt <= 0 { ////0是默认收藏夹，弃用
		glog.Errorf("Invalid \"cid\":%v", cid)
		utils.CtxJsonOfData(ctx, utils.CodePinboxInvalidParamErr, "cid", cid)
		return
	}
	if countInt, err := strconv.Atoi(count); err != nil || countInt <= 0 {
		glog.Errorf("Invalid \"count\":%v", count)
		utils.CtxJsonOfData(ctx, utils.CodePinboxInvalidParamErr, "count", count)
		return
	}

	req, _ := http.NewRequest("GET", fmt.Sprintf(model.GetPinboxFavorItemsByCidApi, cid, count), nil)
	req.Header.Add("authorization", utils.PinboxAuth)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		glog.Error(err)
		utils.CtxJson(ctx, utils.CodePinboxFavorItemsErr)
		return
	}
	//pinbox的api的错误处理做得很差，page不存在或对应cid不存在返回的http码都是404
	if resp.StatusCode != http.StatusOK {
		glog.Errorf("get items from pinbox`s api failed, httpCode:%v.The api may changed(or other reason) or the cid[%v] is not existing.", resp.StatusCode, cid)
		utils.CtxJsonOfData(ctx, utils.CodePinboxFavorItemsErr, "cid", cid)
		return
	}

	defer resp.Body.Close()
	respJson, err := ioutil.ReadAll(resp.Body)
	favorItems := model.PinboxFavorItems{}
	_ = json.Unmarshal(respJson, &favorItems)

	utils.CtxJsonOfData(ctx, utils.CodeSuccess, "pinbox_favoritems", favorItems)

}
