package url

func (r *UrlRepoSQL) GetUrl(shortUrl string) (string, error) {
	query := `
	SELECT url FROM urls
	WHERE short_url = $1`

	var url string

	err := r.db.QueryRow(query, shortUrl).Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}