package converter

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func StartCleanupJob() {
	go func() {
		for {
			files, err := os.ReadDir("downloads")
			if err != nil {
				time.Sleep(30 * time.Minute)
				continue
			}

			for _, file := range files {
				path := filepath.Join("downloads", file.Name())
				info, err := os.Stat(path)

				if err != nil {
					continue
				}
				if time.Since(info.ModTime()) > 15*time.Minute {
					err = os.Remove(path)
					if err == nil {
						log.Println("Deleted:", file.Name())
					}
				}
			}
			time.Sleep(30 * time.Minute)
		}
	}()
}
