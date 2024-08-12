package console

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func RemoveFileStorage(dir string, dayTime int) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Jika bukan file, lewati
		if info.IsDir() {
			return nil
		}

		// Hitung durasi sejak file terakhir dimodifikasi
		age := time.Since(info.ModTime())

		// Jika file lebih tua dari hari yang ditentukan, hapus file tersebut
		if age > time.Duration(dayTime)*24*time.Hour {
			err := os.Remove(path)
			if err != nil {
				log.Printf("Gagal menghapus file %s: %v\n", path, err)
			} else {
				log.Printf("File dihapus: %s\n", path)
			}
		}

		return nil
	})

	if err != nil {
		log.Printf("Gagal melakukan pencarian di direktori: %v\n", err)
	}
}
