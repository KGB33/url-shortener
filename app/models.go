package app

type Url struct {
	Short string `json:"ShortUrl"` // "PrimaryKey"
	Dest  string `json:"DestUrl"`
}

// getUrl retrieves the url entry
// biased on the url.short entry.
func getUrl(short string, s *Server) (Url, error) {
	dest, err := s.DB.Get(ctx, short).Result()
	if err != nil {
		return Url{}, err
	}
	return Url{short, dest}, nil
}

// createUrl creates the url entry in
// the database
func createUrl(u Url, s *Server) (bool, error) {
	return s.DB.SetNX(ctx, u.Short, u.Dest, 0).Result()
}

// updateUrl updates the url entry in database
func updateUrl(u Url, s *Server) error {
	return s.DB.SetXX(ctx, u.Short, u.Dest, 0).Err()
}

// deleteUrl deletes the url in the database
// with the matching url.short
func deleteUrl(short string, s *Server) error {
	return s.DB.Del(ctx, short).Err()
}

// scanUrls returns all urls in the database
func scanUrls(s *Server) ([]Url, error) {
	iter := s.DB.Scan(ctx, 0, "", 0).Iterator()
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
