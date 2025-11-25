package handler

import (
	"errors"
	"net/http"
	"url-shortener/internal/dto"
	apperrors "url-shortener/internal/error"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateShortUrl(c *gin.Context) {
	var req dto.UrlReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl, err := h.Service.CreateUrlShort(&req)
	if err != nil {
		if errors.Is(err, apperrors.ErrInvalidUrl) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": apperrors.ErrInvalidUrl.Error(),
			})
			return
		}
		if errors.Is(err, apperrors.ErrUnExistsUrl) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": apperrors.ErrUnExistsUrl.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, dto.UrlResp{
		ShortUrl: shortUrl,
	})
}

func (h *Handler) GetFullUrl(c *gin.Context) {
	shortUrl := c.Param("short_url")

	url, err := h.Service.GetUrl(shortUrl)
	if err != nil {
		if errors.Is(err, apperrors.ErrUrlNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": apperrors.ErrUrlNotFound.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	
	c.Redirect(http.StatusFound, url)
}
