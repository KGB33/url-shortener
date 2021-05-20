package app

type Url struct {
	Short string `json:"ShortUrl"` // "PrimaryKey"
	Dest  string `json:"DestUrl"`
}

// getUrl retrieves the url entry
// biased on the url.short entry.
func getUrl(short string, s *server) (Url, error) {
	dest, err := s.db.Get(ctx, short).Result()
	if err != nil {
		return Url{}, err
	}
	return Url{short, dest}, nil
}

// postUrl updates or creates the url entry in
// the database
func postUrl(u Url, s *server) error {
	return s.db.Set(ctx, u.Short, u.Dest, 0).Err()
}

// deleteUrl deletes the url in the database
// with the matching url.short
func deleteUrl(short string, s *server) error {
	return s.db.Del(ctx, short).Err()
}

// scanUrls returns all urls in the database
func scanUrls(s *server) ([]Url, error) {
	iter := s.db.Scan(ctx, 0, "", 0).Iterator()
	var urls []Url
	for iter.Next(ctx) {
		nextUrl, err := getUrl(iter.Val(), s)
		if err != nil {
			return urls, err
		}
		urls = append(urls, nextUrl)
	}
	return urls, iter.Err()
}
