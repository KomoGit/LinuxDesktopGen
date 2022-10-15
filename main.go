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

/*Purpose of this tool:
Try to create a .Desktop generator that will easily create executables for Linux Applications.
The application should be controlled with a GUI so you will need to import fyne.
UI will contain:
# Dropdown - To select runner. Wine,native,etc.
# Checkbox(es) - To whether run the app as terminal / Notification. ETC
# User Input - For name,description,etc.
# File Selection - For Icon / Exec (Done)

Some of these options can be optional.*/

/*
TODO :
# Ensure that apart from optionals, nothing else should be empty. If so, program should throw errors. (Done)
*/
var (
	wineRunner           = "wine "
	protonRunner         = "proton-call -r"
	width        float32 = 420
	height       float32 = 420
	appLocation  string
	icoLocation  string
	a            = app.New()
	w            = a.NewWindow("LinuxDesktopGen - v.0.1")
)

func main() {
	generateUI()
}

func generateUI() {
	w.Resize(fyne.NewSize(width, height))
	w.SetFixedSize(true)
	//Widgets
	//Entries
	appName := widget.NewEntry()
	//Buttons
	GenerateFileButton := widget.NewButton("Generate File", func() { go generateFile(appName.Text, appLocation, icoLocation) })

	//FATAL : The application crashes when closing out of file Dialog.
	openFile := widget.NewButton("Select Executable", func() {
		file_Dialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				appLocation = r.URI().String()
			}, w)
		file_Dialog.Show()
	})
	/*Ico File doesn't work. Need to add filter so only image files are being shown.*/
	openIcoFile := widget.NewButton("Select Icon", func() {
		file_Dialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				icoLocation = r.URI().String()
			}, w)
		file_Dialog.Show()
	})

	ExitButton := widget.NewButton("Exit", func() { os.Exit(0) })
	content := container.NewVBox(appName, openFile, openIcoFile, GenerateFileButton, ExitButton)
	w.SetContent(content)
	w.ShowAndRun()
}

func generateFile(fileName string, appLocation string, icoLocation string) {
	if fileName == "" {
		log.Println("Warning, filename cannot be empty!")
		return
	} else {
		file, err := os.Create(fileName + ".desktop")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		if _, err2 := file.WriteString(
			"[Desktop Entry]"); err2 != nil {
			log.Fatal(err2)
		}
		writeExec(*file, appLocation)
		writeIcon(*file, icoLocation)
	}

	file, err := os.Create(fileName + ".desktop")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	//Prefer this over longer error handles.
	if _, err2 := file.WriteString(
		"[Desktop Entry]\nName=" + fileName + "\n"); err2 != nil {
		log.Fatal(err2)
	}
	log.Println("File created successfully")
	writeExec(*file, appLocation)
	writeIcon(*file, icoLocation)
}

// Maybe should take in a slice instead of each thing individually. Perhaps I should move all write functions to a single method.
func writeExec(file os.File, pathToExec string) {
	writeType(file)
	if pathToExec == "" {
		log.Panic("Warning, Executable Path Cannot be empty!")
		return
	}
	if _, err := file.WriteString("Exec= " + wineRunner + pathToExec); err != nil {
		log.Fatal(err)
	}
}

func writeIcon(file os.File, pathToIco string) {
	if _, err := file.WriteString("\nIcon=" + pathToIco); err != nil {
		log.Fatal(err)
	}
}

func writeType(file os.File) {
	if _, err := file.WriteString("\nType=Application\n"); err != nil {
		log.Fatal(err)
	}
}
