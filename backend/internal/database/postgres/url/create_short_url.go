package url

func (r *UrlRepoSQL) CreateShortUrl(url string, shortUrl string) error {
	query := `
	INSERT INTO urls (short_url, url)
	VALUES($1, $2)`

	_, err := r.db.Exec(query, shortUrl, url)
	if err != nil {
		return err
	}

	return nil
}