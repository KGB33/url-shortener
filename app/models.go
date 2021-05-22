package app

import "errors"

type Url struct {
	Short string `json:"ShortUrl"` // "PrimaryKey"
	Dest  string `json:"DestUrl"`
}

// Get retrieves the url entry
// biased on the url.short entry.
func (u *Url) Get(short string, s *Server) error {
	dest, err := s.DB.Get(ctx, short).Result()
	if err != nil {
		return err
	}
	u.Short = short
	u.Dest = dest
	return nil
}

// Create creates the url entry in
// the database
func (u *Url) Create(s *Server) error {
	success, err := s.DB.SetNX(ctx, u.Short, u.Dest, 0).Result()
	if !success {
		return errors.New("Unable to insert the URL into the database. This is likely due to a duplicate ShortUrl.")
	}
	return err
}

// Update updates the url entry in database
func (u *Url) Update(s *Server) error {
	return s.DB.SetXX(ctx, u.Short, u.Dest, 0).Err()
}

// Delete deletes the url in the database
// with the matching url.short
func (u *Url) Delete(s *Server) error {
	return s.DB.Del(ctx, u.Short).Err()
}

// scanUrls returns all urls in the database
func scanUrls(s *Server) ([]Url, error) {
	iter := s.DB.Scan(ctx, 0, "", 0).Iterator()
	var urls []Url
	for iter.Next(ctx) {
		nextUrl := Url{}
		err := nextUrl.Get(iter.Val(), s)
		if err != nil {
			return urls, err
		}
		urls = append(urls, nextUrl)
	}
	return urls, iter.Err()
}
