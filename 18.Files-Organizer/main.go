package main

import(
	"fmt"
	"os"
	"path/filepath"
)


func readFiles(dirPath string) ([]os.DirEntry, error) {
	path, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	return path, nil
}


func getCategory(fileName string) string {
	extension := filepath.Ext(fileName)
	switch extension {
		case ".jpg", ".jpeg", ".png", ".webp", ".gif":
			return "Images"
		case ".mp4", ".mov", ".mkv":
			return "Videos"
		case ".txt", ".pdf", ".doc", ".docx":
			return "Documents"
		default:
			return "Others"
	}
}


func createDirectory(dirPath string, category string) error {
	categoryPath := filepath.Join(dirPath, category)
	return os.MkdirAll(categoryPath, os.ModePerm)
}


func moveFile(filePath string, destDir string) error {
	destPath := filepath.Join(destDir, filepath.Base(filePath))
	return os.Rename(filePath, destPath)
}


func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./file-organizer <path/to/directory>")
		return
	}

	dirPath := os.Args[1]
	fmt.Println("Organizing files in directory:", dirPath)

	files, err := readFiles(dirPath)
	if err != nil {
		fmt.Println("Error reading directory: ", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			category := getCategory(file.Name())
			categoryPath := filepath.Join(dirPath, category)
			err := createDirectory(dirPath, category)
			if err != nil {
				fmt.Println("Error creating folder: ", err)
				continue
			}
			srcPath := filepath.Join(dirPath, file.Name())
			err2 := moveFile(srcPath, categoryPath)
			if err2 != nil {
				fmt.Println("Error moving the file: ", err2)
			} else {
				fmt.Println("Moved: " + file.Name() + " -> " + category)
			}
		}

	}
}