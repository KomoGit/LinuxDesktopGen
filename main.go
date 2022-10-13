package main

import (
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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
/*TODO :
# Ensure that apart from optionals, nothing else should be empty. If so, program should throw errors.

*/
var wineRunner string = "wine " //Standart wine runner.

var width float32 = 420
var height float32 = 420

var a = app.New()
var w = a.NewWindow("LinuxDesktopGen - v.0.1")

func main() {
	generateUI()
}

func generateUI() { //This method should be ran in another thread.
	w.Resize(fyne.NewSize(width, height))
	w.SetFixedSize(true)

	//Widgets
	//Entries
	appName := widget.NewEntry()
	appLocation := widget.NewEntry()
	//Buttons
	GenerateFileButton := widget.NewButton("Generate File", func() { generateFile(appName.Text, appLocation.Text) })
	ExitButton := widget.NewButton("Exit", func() { os.Exit(0) })
	//Descriptions.
	appName.SetPlaceHolder("Application Name")
	appLocation.SetPlaceHolder("Executable Location")

	content := container.NewVBox(appName, appLocation, GenerateFileButton, ExitButton)

	w.SetContent(content)
	w.ShowAndRun()
}

func generateFile(fileName string, appLocation string) {
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
	fmt.Println("File created successfully")
	writeExec(*file, appLocation)
}

// Maybe should take in a slice instead of each thing individually.
func writeExec(file os.File, pathToExec string) {
	_, err2 := file.WriteString(
		"Exec= " + wineRunner + pathToExec + "\nType=Application")

	if err2 != nil {
		log.Fatal(err2)
	}
}
