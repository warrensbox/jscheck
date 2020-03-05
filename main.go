package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/warrensbox/jscheck/lib"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

/*
* Version 0.1.0
* Compatible with Mac OS X AND LINUX ONLY
 */

/*** OPERATION WORKFLOW ***/

var version = "0.1.0\n"

var (
	debugFlag   *bool
	versionFlag *bool
	directory   *string
)

func init() {

	const (
		cmdDesc         = "Install github app binaries on your local machine. Ex: hubapp install mmmorris1975/aws-runas"
		versionFlagDesc = "Displays the version of hubapp"
		actionArgDesc   = "Provide action needed. Ex: install, update, or uninstall"
		giturlArgDesc   = "Provide giturl in user/repo format. Ex: mmmorris1975/aws-runas"
	)

	versionFlag = kingpin.Flag("version", versionFlagDesc).Short('v').Bool()
	directory = kingpin.Flag("directory", versionFlagDesc).Short('d').String()

}

func main() {

	validJSON := true
	dir := "."

	kingpin.CommandLine.Interspersed(false)
	kingpin.Parse()

	//Map of allowed extensions for variable
	//Currently it only accepts 5 extension
	var allowedExt = make(map[string]bool)
	allowedExt[".tfvars"] = true
	allowedExt[".tfvars.json"] = true
	allowedExt[".auto.tfvars"] = true
	allowedExt[".auto.tfvars.json"] = true
	allowedExt[".json"] = true
	allowedExt[".properties"] = true

	if *versionFlag {
		fmt.Printf("Version : %s\n", version)
		os.Exit(0)
	}

	if *directory != "" {
		dir = *directory
	}

	//Walking through current directory
	errWalking := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println("Error walking through directory")
				return err
			}

			//Get path of sub driectory and files
			filePath, errPath := os.Stat(path)
			mode := filePath.Mode()

			if errPath != nil {
				fmt.Println("Error getting path!")
			}

			//if if it is a file; continue
			if mode.IsRegular() {
				fileExt := filepath.Ext(path)         //get file extension
				if _, ok := allowedExt[fileExt]; ok { //check if it's an acceptable extension
					fileContent, _ := ioutil.ReadFile(path)
					fileStr := string(fileContent) //read file content and convert content into string

					if len(fileStr) > 0 { //if file content > 0; continue

						//get the first character in file
						openChar := fileStr[0:3]
						openChar = lib.RemoveNewLine(openChar)

						//get the last character in file
						closeChar := fileStr[len(fileContent)-2:]
						closeChar = lib.RemoveNewLine(closeChar)

						//if the file starts with { and }, it assume that it is a json file; continue
						if openChar == "{" && closeChar == "}" {
							//check if it's a valid json file
							valid := lib.IsJSON(fileStr)
							if !valid {
								fmt.Println("ERROR in JSON:")
								fmt.Println(path)
								validJSON = false
							}
						}
					}

				}
			}

			return nil
		})

	if errWalking != nil {
		log.Println(errWalking)
	}

	if !validJSON {
		os.Exit(1)
	}

	os.Exit(0)
}
