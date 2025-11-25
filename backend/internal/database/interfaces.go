package database

type Url interface {
	GetShortUrl(url string) (string, error)
	CreateShortUrl(url string, shortUrl string) error
	GetUrl(shortUrl string) (string, error)
}