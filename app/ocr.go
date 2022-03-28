package main

import (
    "log"
    "net/http"
	"github.com/otiai10/gosseract/v2"
	"io/ioutil"
	"strings"
)

// TextFromUrl recibes an url to download an image, converts it to bytes[] and
// send it to gosseract.SetImageFromBytes to extraxT the text.
// You can select the language for the extract: "spa":español, "en":english
// Then, some minor corrections are made before returning the ocr text string
func TextFromURL(url string, language string)string{
		response, e := http.Get(url)
		if e != nil {
			log.Fatal(e)
		}
		defer response.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(response.Body)

		client := gosseract.NewClient()
		defer client.Close()
		client.SetImageFromBytes(bodyBytes)
		client.SetLanguage(language)
		text, _ := client.Text()
		text = strings.Replace(text,"º", "o",  -1)
		text = strings.Replace(text,"©", "@",  -1)
		text = strings.Replace(text,"\n", " ",  -1)
		return text	
}

// func main(){
// 	log.Print(TextFromURL("https://pbs.twimg.com/media/EH-Pvo9WwAEKFwc?format=jpg&name=small", "spa"))
// }