package dto

type UrlReq struct {
	Url string `json:"url" binding:"required"`
}

type UrlResp struct {
	ShortUrl string `json:"short_url"`
}