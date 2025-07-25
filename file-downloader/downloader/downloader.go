package downloader

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync" // Import the sync package for WaitGroup
	"time" // Already imported for unique filename, but useful for observing delays
)

var urlsToDownload []string
var urlin_raw string
var urlin string
var wg sync.WaitGroup // Declare a WaitGroup

// downloadFile downloads a single file from a given URL to a specified destination.
// It now also takes a WaitGroup pointer to signal completion.
func downloadGo(url string, destinationDir string, wg *sync.WaitGroup) {
	// IMPORTANT: Signal that this goroutine is done when the function exits.
	// This ensures wg.Wait() in main eventually unblocks.
	defer wg.Done()

	// 1. Extract filename from URL
	filename := filepath.Base(url)
	if filename == "." || filename == "/" || strings.TrimSpace(filename) == "" {
		filename = "downloaded_file_" + fmt.Sprintf("%d", time.Now().UnixNano())
	}
	filePath := filepath.Join(destinationDir, filename)

	fmt.Printf("[%s] Starting download: %s to %s\n", time.Now().Format("15:04:05"), url, filePath)

	// 2. Create the output file
	out, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("ERROR: Failed to create file %s for URL %s: %v\n", filePath, url, err)
		return // Exit goroutine on error
	}
	defer out.Close()

	// 3. Make HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("ERROR: Failed to fetch URL %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	// 4. Check HTTP response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("ERROR: Bad status for URL %s: %s\n", url, resp.Status)
		return
	}

	// 5. Copy data from response body to file
	bytesWritten, err := io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("ERROR: Failed to write data to file %s for URL %s: %v\n", filePath, url, err)
		return
	}

	fmt.Printf("[%s] Finished download: %s (%d bytes written)\n", time.Now().Format("15:04:05"), url, bytesWritten)
}

func urlInput() {
	urlin_raw = ""
	urlin = ""

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input the download URL: ")
	urlin_raw, err := reader.ReadString('\n')
	for err != nil {
		fmt.Println("Error:", err)
		fmt.Print("Enter task: ")
		urlin_raw, err = reader.ReadString('\n')
	}
	urlin = strings.TrimSpace(urlin_raw)
}

func Project() {
	fmt.Println("--- Parallel File Downloader (Goroutines & WaitGroup) ---")

	downloadDir := "downloads"
	os.MkdirAll(downloadDir, 0755) // Ensure downloads directory exists

	for urlin != "back" && urlin != "process" {
		urlInput()

		switch urlin {
		case "back":
			fmt.Println("Exiting...")
		case "process":
			for _, url := range urlsToDownload {
				wg.Add(1) // Increment the WaitGroup counter for each goroutine we are about to launch

				// Launch the downloadFile function as a goroutine
				// Pass the WaitGroup pointer so the goroutine can signal completion
				go downloadGo(url, downloadDir, &wg)
			}
		default:
			urlsToDownload = append(urlsToDownload, urlin)

		}
	}

	fmt.Println("\nMain goroutine: All download goroutines launched. Waiting for them to finish...")

	// Block the main goroutine until all launched goroutines have called wg.Done()
	wg.Wait()

	fmt.Println("\nMain goroutine: All downloads completed. Program finished.")
}
