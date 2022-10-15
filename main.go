package main

import (
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
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
	w            = a.NewWindow("LinuxDesktopGen - v.1.0")
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
	appComment := widget.NewEntry()
	//Buttons
	GenerateFileButton := widget.NewButton("Generate File", func() { go generateFile(appName.Text, appLocation, icoLocation, appComment.Text) })
	//Placeholders

	//FATAL : The application crashes when closing out of file Dialog.
	openFile := widget.NewButton("Select Executable", func() {
		file_Dialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				appLocation = r.URI().String()
			}, w)
		file_Dialog.Show()
	})
	/*Ico File doesn't work. Need to add filter so only image files are being shown. (FIXED)*/
	openIcoFile := widget.NewButton("Select Icon", func() {
		file_Dialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				icoLocation = r.URI().String()
			}, w)
		file_Dialog.SetFilter(storage.NewExtensionFileFilter([]string{".png"})) //Need to make sure more than 1 single extension is allowed.
		file_Dialog.Show()
	})

	ExitButton := widget.NewButton("Exit", func() { os.Exit(0) })
	content := container.NewVBox(appName, appComment, openFile, openIcoFile, GenerateFileButton, ExitButton)
	w.SetContent(content)
	w.ShowAndRun()
}

func generateFile(fileName string, appLocation string, icoLocation string, appComment string) {
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
		writeComment(*file, appComment)
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
		"[Desktop Entry]\nName=" + fileName); err2 != nil {
		log.Fatal(err2)
	}
	log.Println("File created successfully")
	writeType(*file)
	writeExec(*file, appLocation)
	writeIcon(*file, icoLocation)
	writeComment(*file, appComment)
}

func writeType(file os.File) {
	if _, err := file.WriteString("\nType=Application"); err != nil {
		log.Fatal(err)
	}
}

func writeExec(file os.File, pathToExec string) {
	if pathToExec == "" {
		log.Panic("Warning, Executable Path Cannot be empty!")
		return
	}
	if _, err := file.WriteString("\nExec= " + wineRunner + pathToExec); err != nil {
		log.Fatal(err)
	}
}

func writeIcon(file os.File, pathToIco string) {
	res := strings.Split(pathToIco, "file://")
	fltrPath := string(res[len(res)-1]) //Filtered path turned into string (From array)
	if _, err := file.WriteString("\nIcon=" + fltrPath); err != nil {
		log.Fatal(err)
	}
}

func writeComment(file os.File, comment string) {
	if _, err := file.WriteString("\nComment=" + comment); err != nil {
		log.Fatal(err)
	}
}
