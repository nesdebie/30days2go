package main

import (
    "archive/zip"
    "fmt"
    "io"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: " + os.Args[0] + " <file1> <file2> ... <fileN>")
        return
    }
    files := os.Args[1:]
    for _, file := range files {
        fileToCheck, err := os.Open(file)
        if err != nil {
            fmt.Println("Error opening file" + file + " : " + err.Error())
            return
        }
        fileToCheck.Close()
    }
    zipFile, err := os.Create("compressed.zip")
    if err != nil {
        fmt.Println("Error creating zip file: " + err.Error())
        return
    }
    // Ensure the zip file is closed properly at the end of the program only
    defer zipFile.Close()

    zipWriter := zip.NewWriter(zipFile)
    defer zipWriter.Close()

    for _, file := range files {
        fileToZip, err := os.Open(file)
        if err != nil {
            fmt.Println("Error opening file " + file + ": " + err.Error())
            return
        }

        fileInfo, err := fileToZip.Stat()
        if err != nil {
            fmt.Println("Error getting file info for " + file + ": " + err.Error())
            return
        }
        header, err := zip.FileInfoHeader(fileInfo)
        if err != nil {
            fmt.Println("Error creating zip header for " + file + ": " + err.Error())
            return
        }

        header.Name = file

        writer, err := zipWriter.CreateHeader(header)
        if err != nil {
            fmt.Println("Error creating zip writer for " + file + ": " + err.Error())
            return
        }

        _, err = io.Copy(writer, fileToZip)
        fileToZip.Close()
        if err != nil {
            fmt.Println("Error copying file " + file + " to zip: " + err.Error())
            return
        }
    }
    fmt.Println("Zip archive created successfully!")
}