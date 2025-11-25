package url

func (r *UrlRepoSQL) GetShortUrl(url string) (string, error){
	query := `
	SELECT short_url FROM urls
	WHERE url = $1`

	var urlShort string

	err := r.db.QueryRow(query, url).Scan(&urlShort)
	if err != nil {
		return "", err
	}

	return urlShort, nil
}