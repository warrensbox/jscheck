package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

/*
* Version 0.1.0
* Compatible with Mac OS X AND LINUX ONLY
 */

/*** OPERATION WORKFLOW ***/

func main() {

	f, err := os.Open("generals.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	contentType, err := GetFileContentType(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("Content Type: " + contentType)

	// buf, _ := ioutil.ReadFile("version")

	// kind, _ := filetype.Match(buf)
	// if kind == filetype.Unknown {
	// 	fmt.Println("Unknown file type")
	// 	return
	// }

	// fmt.Printf("File type: %s. MIME: %s\n", kind.Extension, kind.MIME.Value)

	os.Exit(0)
	validJson := true

	//Map of allowed extensions for variable
	//Currently it only accepts 4 extension
	var allowedExt = make(map[string]bool)
	allowedExt[".tfvars"] = true
	allowedExt[".tfvars.json"] = true
	allowedExt[".auto.tfvars"] = true
	allowedExt[".auto.tfvars.json"] = true

	//Walking through current directory
	errWalking := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println("Error walking through directory")
				return err
			}

			//Get path of sub driectory and files
			filePath, errPath := os.Stat(path)
			mode := filePath.Mode()

			if errPath != nil {
				fmt.Println("Error getting path")
			}

			//if if it is a file; continue
			if mode.IsRegular() {
				fileExt := filepath.Ext(path)         //get file extension
				if _, ok := allowedExt[fileExt]; ok { //check if it's an acceptable extension
					fileContent, _ := ioutil.ReadFile(path)
					fileStr := string(fileContent) //read file content and convert content into string

					if len(fileStr) > 0 { //if file content > 0; continue

						//get the first character in file
						openChar := fileStr[0:1]
						openChar = strings.TrimPrefix(openChar, "\n") //remove new line from beginning
						openChar = strings.TrimSuffix(openChar, "\n") //remove new line from end of line

						//get the last character in file
						closeChar := fileStr[len(fileContent)-2:]
						closeChar = strings.TrimPrefix(closeChar, "\n") //remove new line from beginning
						closeChar = strings.TrimSuffix(closeChar, "\n") //remove new line from end of line

						//if the file starts with { and }, it assume that it is a json file; continue
						if openChar == "{" && closeChar == "}" {
							//check if it's a valid json file
							valid := isJSON(fileStr)
							if !valid {
								fmt.Println("ERROR in JSON:")
								fmt.Println(path)
								validJson = false
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

	if !validJson {
		os.Exit(1)
	}

	os.Exit(0)
}

//isJSON : check if valid json
func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
