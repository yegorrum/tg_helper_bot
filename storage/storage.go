package storage

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	e "tg_bot/lib"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

var ErrNoSAvePages = errors.New("no save pages")

type Page struct {
	URL      string
	UserName string
}

func (p Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("cant calc hash", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("cant calc hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
