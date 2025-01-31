// package main

// import (
// 	"archive/zip"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"time"
// )

// func backupFiles(sourceDir, backupName string) {
// 	outFile, _ := os.Create(backupName)
// 	defer outFile.Close()
// 	zipWriter := zip.NewWriter(outFile)
// 	defer zipWriter.Close()

// 	files, _ := os.ReadDir(sourceDir)
// 	for _, file := range files {
// 		srcFile, _ := os.Open(filepath.Join(sourceDir, file.Name()))
// 		defer srcFile.Close()
// 		wr, _ := zipWriter.Create(file.Name())
// 		srcFile.WriteTo(wr)
// 	}
// 	fmt.Println("Backup completed:", backupName)
// }

// func main() {
// 	for {
// 		backupFiles("./files", fmt.Sprintf("backup-%d.zip", time.Now().Unix()))
// 		time.Sleep(24 * time.Hour) // Adjust interval as needed
// 	}
// }

// The code is functional but lacks proper error handling, which is crucial for robustness. Additionally, using defer inside a loop can lead to resource leaks because the deferred calls will not execute until the function returns. Here is an improved version with error handling and proper resource management.

package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func backupFiles(sourceDir, backupName string) error {
	outFile, err := os.Create(backupName)
	if err != nil {
		return fmt.Errorf("error creating backup file: %v", err)
	}
	defer outFile.Close()

	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()

	files, err := os.ReadDir(sourceDir)
	if err != nil {
		return fmt.Errorf("error reading source directory: %v", err)
	}

	for _, file := range files {
		srcFile, err := os.Open(filepath.Join(sourceDir, file.Name()))
		if err != nil {
			return fmt.Errorf("error opening source file: %v", err)
		}
		defer srcFile.Close()

		wr, err := zipWriter.Create(file.Name())
		if err != nil {
			return fmt.Errorf("error creating zip entry: %v", err)
		}

		if _, err := io.Copy(wr, srcFile); err != nil {
			return fmt.Errorf("error writing to zip: %v", err)
		}
	}

	fmt.Println("Backup completed:", backupName)
	return nil
}

func main() {
	for {
		err := backupFiles("./files", fmt.Sprintf("backup-%d.zip", time.Now().Unix()))
		if err != nil {
			fmt.Printf("Backup failed: %v\n", err)
		}
		time.Sleep(24 * time.Hour) // Adjust interval as needed
	}
}
