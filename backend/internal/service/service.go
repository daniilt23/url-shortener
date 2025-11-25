package service

import (
	"database/sql"
	"errors"
	"net/http"
	urls "net/url"
	"url-shortener/internal/database/postgres/url"
	"url-shortener/internal/dto"
	apperrors "url-shortener/internal/error"

	"github.com/thanhpk/randstr"
)

type Service struct {
	UrlRepoSQL *url.UrlRepoSQL
}

func NewService(urtRepoSQL *url.UrlRepoSQL) *Service {
	return &Service{
		UrlRepoSQL: urtRepoSQL,
	}
}

func (s *Service) CreateUrlShort(req *dto.UrlReq) (string, error) {
	_, err := urls.ParseRequestURI(req.Url)
	if err != nil {
		return "", apperrors.ErrInvalidUrl
	}

	resp, err := http.Get(req.Url)
	if err != nil || !(resp.StatusCode >= 200 && resp.StatusCode < 400) {
		return "", apperrors.ErrUnExistsUrl
	}

	shortUrl, err := s.UrlRepoSQL.GetShortUrl(req.Url)
	if err == nil {
		return shortUrl, nil
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return "", err
	}

	shortUrl = s.generateShortUrl()

	err = s.UrlRepoSQL.CreateShortUrl(req.Url, shortUrl)
	if err != nil {
		return "", err
	}
	
	return shortUrl, nil
}

func (s *Service) GetUrl(shortUrl string) (string, error) {
	url, err := s.UrlRepoSQL.GetUrl(shortUrl)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", apperrors.ErrUrlNotFound
		}
		return "", err
	}

	return url, nil
}

func (s *Service) generateShortUrl() string {
	return randstr.String(6)
}
