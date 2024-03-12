package main

import (
	"fmt"
	"os"

	"github.com/jung-kurt/gofpdf"
)

// write content to .txt file
func writeContentToFile() {
	// store the content in a file
	content := "This is a sample text that we need to store in a txt file using golang"
	fmt.Println(content)
	file, err := os.Create("./sample.pdf")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	len, err := file.WriteString(content)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len, "bytes written successfully")
}

// write content to .pdf file
func writeContentToPDF() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		fmt.Println(err)
	}
}

// read content from directory and store in pdf
func readContentFromDirectory() {
	dir := "./images"
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(files))
	pdf := gofpdf.New("P", "mm", "A4", "")

	// store images in pdf
	for _, file := range files {
		// check if file is image
		if file.IsDir() {
			continue
		}
		// add page if no space  left
		pdf.AddPage()
		pdf.Image(dir+"/"+file.Name(), 10, 10, 200, 0, false, "", 0, "")
		fmt.Println(file.Name())
	}
	err = pdf.OutputFileAndClose("images.pdf")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("PDF created successfully")

}

// get properties of file
func getProperties() {
	dir := "./images"
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(files))
	for _, file := range files {
		f, err := os.OpenFile(dir+"/"+file.Name(), os.O_RDONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		fi, err := f.Stat()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("File Name:", fi.Name())
		fmt.Println("File Size:", fi.Size())
		fmt.Println("File Mode:", fi.Mode())
		fmt.Println("File ModTime:", fi.ModTime())
		fmt.Println("File IsDir:", fi.IsDir())
		fmt.Println("File Sys:", fi.Sys())
	}
}

// read all file from a directory and rename them
func readContentAndRename() {
	dir := "D:/Learning/content/Reels"
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(files))
	renamedDir := "D:/Learning/content/reels"
	// check if directory exists
	_, err = os.Stat(renamedDir)
	if os.IsNotExist(err) {
		// create directory
		err = os.Mkdir(renamedDir, 0755)
	}
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(files); i++ {
		err = os.Rename(dir+"/"+files[i].Name(), renamedDir+"/reel_"+fmt.Sprint(i+1)+".mp4")
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Files renamed successfully")

}
