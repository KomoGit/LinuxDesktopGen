package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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
# File Selection - For Icon / Exec

Some of these options can be optional.
*/
/*TODO :
# Ensure that apart from optionals, nothing else should be empty. If so, program should throw errors. (Done)

*/
var wineRunner string = "wine " //Standart wine runner.

var width float32 = 420
var height float32 = 420

var appLocation string

var a = app.New()
var w = a.NewWindow("LinuxDesktopGen - v.0.1")

func main() {
	generateUI()
}

func generateUI() { //This method should be ran in main.
	w.Resize(fyne.NewSize(width, height))
	w.SetFixedSize(true)
	//Widgets
	//Entries
	appName := widget.NewEntry()
	//Buttons
	GenerateFileButton := widget.NewButton("Generate File", func() { go generateFile(appName.Text, appLocation) })

	//FATAL : The application crashes when closing out of file Dialog.
	openFile := widget.NewButton("Open Executables", func() {
		file_Dialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				appLocation = r.URI().String()
				log.Println(appLocation)
			}, w)
		file_Dialog.Show()
	})

	ExitButton := widget.NewButton("Exit", func() { os.Exit(0) })
	//Descriptions.
	appName.SetPlaceHolder("Application Name")

	content := container.NewVBox(appName, openFile, GenerateFileButton, ExitButton)

	w.SetContent(content)
	w.ShowAndRun()
}

// This should be ran in another thread.
func generateFile(fileName string, appLocation string) {
	if fileName == "" {
		log.Println("Warning, filename cannot be empty!")
	} else {
		file, err := os.Create(fileName + ".desktop")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		_, err2 := file.WriteString(
			"[Desktop Entry]\nName=" + fileName + "\n")

		if err2 != nil {
			log.Fatal(err2)
		}
		log.Println("File created successfully")
		writeExec(*file, appLocation)
	}
}

// Maybe should take in a slice instead of each thing individually. Perhaps I should move all write functions to a single method.
func writeExec(file os.File, pathToExec string) {
	if pathToExec == "" {
		log.Panic("Warning, Executable Path Cannot be empty!")
	} else {
		_, err2 := file.WriteString(
			"Exec= " + wineRunner + pathToExec + "\nType=Application") //Move the application types into different function.

		if err2 != nil {
			log.Fatal(err2)
		}
	}
}

func writeIcon(file os.File, pathToIco string) {
	_, err2 := file.WriteString(
		"Icon= " + pathToIco) //Move the application types into different function.

	if err2 != nil {
		log.Fatal(err2)
	}
}
