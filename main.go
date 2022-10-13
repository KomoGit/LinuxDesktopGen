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
//Runners
var wineRunner string = "wine"

var width float32 = 420
var height float32 = 420

var a = app.New()
var w = a.NewWindow("Hello World")

func main() {
	generateUI()
}

func generateUI() { //This method should be ran in another thread.
	w.Resize(fyne.NewSize(width, height))
	w.SetFixedSize(true)

	//Widgets
	appName := widget.NewEntry()
	appLocation := widget.NewEntry()

	//Descriptions.
	appName.SetPlaceHolder("Application Name")
	appLocation.SetPlaceHolder("Executable Location")

	content := container.NewVBox(appName, appLocation, widget.NewButton("Generate File", func() { generateFile(appName.Text) }))

	w.SetContent(content)
	w.ShowAndRun()
}
func generateFile(fileName string) {
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
}

// Maybe should take in a slice instead of each thing individually.
func writeToFile(f string) {

}
