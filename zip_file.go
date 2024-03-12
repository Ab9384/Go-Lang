package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func createZipFile() {
	// location of the files to be zipped
	files := []string{}
	imageDir := "./images"
	filepath.Walk(imageDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	fmt.Println(len(files))
	for _, file := range files {
		fmt.Println(file)
	}

	zipFile, err := os.Create("images.zip")
	if err != nil {
		fmt.Println(err)
	}
	defer zipFile.Close()

	// create a new zip archive
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// add files to the archive
	for _, file := range files {
		// open the file
		fileToZip, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
		}
		defer fileToZip.Close()

		// get the file info
		info, err := fileToZip.Stat()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(info.Name())

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			fmt.Println(err)
		}
		// set the compression method to deflate
		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			fmt.Println(err)
		}
		_, err = io.Copy(writer, fileToZip)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Zipped File:", zipFile.Name())

}

// read content from zip file
func readContentFromZipFile() {
	// open a zip archive for reading
	zipReader, err := zip.OpenReader("images.zip")
	if err != nil {
		fmt.Println(err)
	}
	defer zipReader.Close()
	zipDir := "./unzipped"
	os.Mkdir(zipDir, os.ModePerm)

	// iterate through the files in the archive
	for _, file := range zipReader.Reader.File {
		zippedFile, err := file.Open()
		if err != nil {
			fmt.Println(err)
		}
		defer zippedFile.Close()

		// specify the file path
		extractedFilePath := filepath.Join(zipDir, file.Name)

		// check for the file path
		if file.FileInfo().IsDir() {
			os.MkdirAll(extractedFilePath, os.ModePerm)
		} else {
			// open the file to write
			extractedFile, err := os.OpenFile(extractedFilePath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}
			defer extractedFile.Close()

			_, err = io.Copy(extractedFile, zippedFile)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
