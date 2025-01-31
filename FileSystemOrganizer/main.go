// package main

// import (
// 	"os"
// 	"path/filepath"
// )

// func main() {
// 	files, _ := os.ReadDir(".")
// 	for _, file := range files {
// 		if !file.IsDir() {
// 			ext := filepath.Ext(file.Name())
// 			os.Mkdir(ext, 0755)
// 			os.Rename(file.Name(), filepath.Join(ext, file.Name()))
// 		}
// 	}
// }

// The code is functional but lacks error handling and could be improved for better readability and robustness. Here is an improved version with error handling:

package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			ext := filepath.Ext(file.Name())
			if ext == "" {
				ext = "no_extension"
			}
			err := os.MkdirAll(ext, 0755)
			if err != nil {
				log.Printf("Failed to create directory %s: %v", ext, err)
				continue
			}
			err = os.Rename(file.Name(), filepath.Join(ext, file.Name()))
			if err != nil {
				log.Printf("Failed to move file %s: %v", file.Name(), err)
			}
		}
	}
}
