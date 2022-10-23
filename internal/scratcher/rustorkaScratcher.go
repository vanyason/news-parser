package scratcher

import "github.com/vanyason/news-parser/internal/log"

type RustorkaScratcher struct {
}

func (s *RustorkaScratcher) Scratch() (res Result, err error) {
	log.Debug("Rustorka scratcher")
	return res, nil
}
