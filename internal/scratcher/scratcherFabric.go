package scratcher

import (
	"github.com/vanyason/news-parser/internal/config"
)

func CreateScratchers(conf config.Config) (scratchers []Scratcher, err error) {
	scratchers = []Scratcher{
		&RustorkaScratcher{},
	}

	return scratchers, nil
}
