package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vearne/tinyurl/config"
	"github.com/vearne/tinyurl/dao"
	"github.com/vearne/tinyurl/libs"
	zlog "github.com/vearne/tinyurl/log"
	"github.com/vearne/tinyurl/resource"
	"go.uber.org/zap"
	"net/http"
)

type URLParam struct {
	URL string `form:"url" json:"url" binding:"required"`
}

type URLResp struct {
	TinyURL string `json:"tinyurl"`
}

func UrlChange(c *gin.Context) {
	var param URLParam
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "E001", "msg": "参数错误或者缺失"})
		return
	}

	var res URLResp
	value := dao.CreateTinyURL(param.URL)
	res.TinyURL = config.GetOpts().Domain + libs.UintToBase62(value)
	c.JSON(http.StatusOK, res)
}

func UrlGet(c *gin.Context) {
	var url string
	sid := c.Param("sid")

	url, err := getUrl(sid)
	if err != nil {
		zlog.Error("can't parse tinyurl", zap.String("sid", sid))
		url = "/"
	}
	c.Redirect(http.StatusFound, url)
}

func getUrl(sid string) (string, error) {
	value, err := libs.Base62ToUint(sid)
	if err != nil {
		return "", err
	}
	// 尝试从cache中获取
	res := resource.FixedCache.Get(value)
	if res != nil {
		return res.(string), nil
	}
	// 尝试从数据库中获取
	url := dao.GetURL(value)
	if len(url) > 0 {
		resource.FixedCache.Set(value, url, -1)
		return url, nil
	}

	return "", errors.New("can't find url")
}
