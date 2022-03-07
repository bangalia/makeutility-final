package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	urls := []string{
		"https://exploringjs.com/impatient-js/downloads/impatient-js-preview-book.pdf",
		"http://www.cheat-sheets.org/saved-copy/2084227-Mac-OS-X-Terminal-Commands-list.pdf",
		"https://www.dominican.edu/sites/default/files/2020-10/five-year-calendar_2020-2025.pdf",
		"https://education.github.com/git-cheat-sheet-education.pdf",
	}

	var wg sync.WaitGroup //WaitGroup waits for all launched goroutines to finish

	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			tokens := strings.Split(url, "/")
			fileName := tokens[len(tokens)-1]
			fmt.Println("Downloading", url, "to", fileName)
			//error expectations
			output, err := os.Create(fileName)
			if err != nil {
				log.Fatal("File creation error", fileName, "-", err)
			}
			defer output.Close()

			res, err := http.Get(url)
			if err != nil {
				log.Fatal("Site connection error: ", err)
			} else {
				defer res.Body.Close()
				_, err = io.Copy(output, res.Body)
				if err != nil {
					log.Fatal("Download failed", url, "-", err)
				} else {
					fmt.Println("Downloaded", fileName)
				}
			}
		}(url)
	}
	wg.Wait() // waits for full function to execute before printing statement
	fmt.Println("Download completed")
}
