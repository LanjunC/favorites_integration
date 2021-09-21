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
	"strings"
)

// GetZhihuFavors 获取知乎所有收藏夹信息
func GetZhihuFavors(ctx *gin.Context) {
	limit := ctx.DefaultQuery("limit", "10")
	offset := ctx.DefaultQuery("offset", "0")
	if limitInt, err := strconv.Atoi(limit); err != nil || limitInt <= 0 {
		glog.Errorf("Invalid \"limit\":%v", limit)
		utils.CtxJsonOfData(ctx, utils.CodeZhihuInvalidParamErr, "limit", limit)
		return
	}
	if offsetInt, err := strconv.Atoi(offset); err != nil || offsetInt < 0 {
		glog.Errorf("Invalid \"offset\":%v", offset)
		utils.CtxJsonOfData(ctx, utils.CodeZhihuInvalidParamErr, "offset", offset)
		return
	}

	//调zhihu api 获取知乎所有收藏夹信息
	resp, err := http.Get(fmt.Sprintf(model.GetZhihuFavorsApi+"?limit=%v&offset=%v", limit, offset))
	if err != nil {
		glog.Error(err)
		utils.CtxJson(ctx, utils.CodeZhihuFavorsErr)
		return
	}
	if resp.StatusCode != http.StatusOK {
		glog.Errorf("the api may changed(or other reason), httpcode:%v", resp.StatusCode)
		utils.CtxJson(ctx, utils.CodeZhihuFavorsErr)
		return
	}

	defer resp.Body.Close()
	respJson, err := ioutil.ReadAll(resp.Body)
	zhihuFavors := model.ZhihuFavors{}
	_ = json.Unmarshal(respJson, &zhihuFavors)
	//glog.Info(resp)

	//处理信息后发送给前端
	IndexAnyOfNextParam := strings.IndexAny(zhihuFavors.Paging.Next, "?")
	zhihuFavors.Paging.Next = string([]byte(zhihuFavors.Paging.Next)[IndexAnyOfNextParam+1:])
	IndexAnyOfPreviousParam := strings.IndexAny(zhihuFavors.Paging.Previous, "?")
	zhihuFavors.Paging.Previous = string([]byte(zhihuFavors.Paging.Previous)[IndexAnyOfPreviousParam+1:])

	utils.CtxJsonOfData(ctx, utils.CodeSuccess, "zhihu_favors", zhihuFavors)
}

// GetZhihuFavorItemsByCid 通过cid获取该收藏夹所有items
func GetZhihuFavorItemsByCid(ctx *gin.Context) {
	limit := ctx.DefaultQuery("limit", "10")
	offset := ctx.DefaultQuery("offset", "0")
	cid := ctx.Param("cid")
	if limitInt, err := strconv.Atoi(limit); err != nil || limitInt <= 0 {
		glog.Errorf("Invalid \"limit\":%v", limit)
		utils.CtxJsonOfData(ctx, utils.CodeZhihuInvalidParamErr, "limit", limit)
		return
	}
	if offsetInt, err := strconv.Atoi(offset); err != nil || offsetInt < 0 {
		glog.Errorf("Invalid \"offset\":%v", offset)
		utils.CtxJsonOfData(ctx, utils.CodeZhihuInvalidParamErr, "offset", offset)
		return
	}
	if cidInt, err := strconv.Atoi(cid); err != nil || cidInt <= 0 {
		glog.Errorf("Invalid \"cid\":%v", cid)
		utils.CtxJsonOfData(ctx, utils.CodeZhihuInvalidParamErr, "cid", cid)
		return
	}

	//调zhihu api 通过cid获取该收藏夹所有items
	resp, err := http.Get(fmt.Sprintf(model.GetZhihuFavorItemsByCidApi+"?limit=%v&offset=%v", cid, limit, offset))
	if err != nil {
		glog.Error(err)
		utils.CtxJson(ctx, utils.CodeZhihuFavorItemsErr)
		return
	}
	if resp.StatusCode != http.StatusOK {
		glog.Errorf("the api may changed(or other reason), httpcode:%v", resp.StatusCode)
		utils.CtxJson(ctx, utils.CodeZhihuFavorItemsErr)
		return
	}

	defer resp.Body.Close()
	respJson, err := ioutil.ReadAll(resp.Body)
	favorItems := model.ZhihuFavorItems{}
	_ = json.Unmarshal(respJson, &favorItems)
	if favorItems.Data == nil {
		glog.Errorf("not found, \"cid\":%v", cid)
		utils.CtxJsonOfData(ctx, utils.CodeZhihuNotFoundErr, "cid", cid)
		return
	}

	//处理信息后发送给前端
	IndexAnyOfNextParam := strings.IndexAny(favorItems.Paging.Next, "?")
	favorItems.Paging.Next = string([]byte(favorItems.Paging.Next)[IndexAnyOfNextParam+1:])
	IndexAnyOfPreviousParam := strings.IndexAny(favorItems.Paging.Previous, "?")
	favorItems.Paging.Previous = string([]byte(favorItems.Paging.Previous)[IndexAnyOfPreviousParam+1:])

	utils.CtxJsonOfData(ctx, utils.CodeSuccess, "zhihu_favoritems", favorItems)
}
