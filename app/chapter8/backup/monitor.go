package backup

import (
	"fmt"
	"path/filepath"
	"time"
)

type Monitor struct {
	Paths       map[string]string
	Archiver    Archiver
	Destination string
}

func (m *Monitor) Now() (int, error) {
	var counter int
	for path, lastHash := range m.Paths {
		newHash, err := DirHash(path)
		if err != nil {
			return 0, nil
		}
		if newHash != lastHash {
			if err := m.act(path); err != nil {
				return counter, err
			}
		}
	}
	return counter, nil
}

func (m *Monitor) act(path string) error {
	dirname := filepath.Base(path)
	filename := fmt.Sprintf("%d.zip", time.Now().UnixNano())
	return m.Archiver.Archive(path, filepath.Join(m.Destination, dirname, filename))
}
