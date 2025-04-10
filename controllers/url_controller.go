package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rishavkumar7/url-shortner/constants"
	"github.com/rishavkumar7/url-shortner/database"
	"github.com/rishavkumar7/url-shortner/helper"
	"github.com/rishavkumar7/url-shortner/types"
)

func ShortTheUrl(c *gin.Context) {
	var shortUrlBody types.ShortUrlBody
	err := c.BindJSON(&shortUrlBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": constants.BindError,
			"error":   true,
		})
		return
	}

	code := helper.GenerateRandomString(constants.ShortUrlCodeLength)
	record, _ := database.Mgr.GetUrlFromCode(code, constants.UrlCollection)

	if record.UrlCode != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": constants.ExistError,
			"error":   true,
		})
		return
	}

	var url types.UrlDb
	url.CreatedAt = time.Now().Unix()
	url.ExpiredAt = time.Now().Add(time.Hour * 24 * 30).Unix()
	url.LongUrl = shortUrlBody.LongUrl
	url.UrlCode = code
	url.ShortUrl = constants.BaseUrl + os.Getenv("API_VERSION") + "/url/" + code

	resp, err := database.Mgr.Insert(url, constants.UrlCollection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      resp,
		"short_url": url.ShortUrl,
		"error":     false,
	})
}

func RedirectToLongUrl(c *gin.Context) {
	code := c.Param("code")

	record, _ := database.Mgr.GetUrlFromCode(code, constants.UrlCollection)
	if record.UrlCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": constants.NotFoundError,
			"error":   true,
		})
		return
	}

	c.Redirect(http.StatusPermanentRedirect, record.LongUrl)
}
