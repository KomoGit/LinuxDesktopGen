package main

import (
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

/*
Purpose of this tool:
Try to create a .Desktop generator that will easily create executables for Linux Applications.
The application should be controlled with a GUI so you will need to import fyne.
UI will contain:
# Dropdown - To select runner. Wine,native,etc.
# Checkbox(es) - To whether run the app as terminal / Notification. ETC
# User Input - For name,description,etc.
# File Selection - For Icon.

Some of these options can be optional.
*/

var width float32 = 420
var height float32 = 420

func main() {
	generateUI()
}

func generateUI() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.NewSize(width, height))
	w.SetFixedSize(true)

	content := widget.NewButton("Generate File", func() {
		generateFile("Test")
	})
	w.SetContent(content)
	w.ShowAndRun()
}
func generateFile(fileName string) {
	file, err := os.Create(fileName + ".desktop")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File created successfully")
	defer file.Close()
}
