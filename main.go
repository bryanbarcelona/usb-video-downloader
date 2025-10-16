package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
)

const (
	urlsFile       = "urls.txt"
	logFile        = "download_log.txt"
	videoDir       = "download_videos"
	ytDlpBinary    = ".\\yt-dlp.exe"
	outputTemplate = "%(title)s.%(ext)s"
)

func main() {

	fmt.Print(`
	╔═══════════════════════════════════╗
	║        USB Video Downloader       ║
	║      Portable video tool          ║
	║                                   ║
	║   Usage: Add URLs to urls.txt     ║
	║          and run the program      ║
	╚═══════════════════════════════════╝
	`)

	fmt.Println("Updating yt-dlp...")
	updateCmd := exec.Command(ytDlpBinary, "-U")
	if err := updateCmd.Run(); err != nil {
		fmt.Printf("yt-dlp update failed (proceeding with current version): %v\n", err)
	}

	// Ensure video directory exists
	if err := os.MkdirAll(videoDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create video dir: %v\n", err)
		os.Exit(1)
	}

	// Read all URLs
	file, err := os.Open(urlsFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("urls.txt not found. Create it with URLs (one per line).")
			return
		}
		fmt.Fprintf(os.Stderr, "Failed to open urls.txt: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			urls = append(urls, line)
		}
	}

	if len(urls) == 0 {
		fmt.Println("No URLs found in urls.txt")
		return
	}

	logF, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open log file: %v\n", err)
		os.Exit(1)
	}
	defer logF.Close()

	successCount := 0
	bar := progressbar.Default(int64(len(urls)))

	fmt.Printf("\nStarting download of %d Videos...\n", len(urls))

	for _, url := range urls {
		fmt.Printf("\nUrl: %s\n", url)

		// Step 1: Ask yt-dlp for title and canonical URL
		titleCmd := exec.Command(ytDlpBinary, "--print", "title", "--print", "webpage_url", url)
		out, err := titleCmd.CombinedOutput()

		// Print ONLY the title (first line of output)
		linesForPrint := strings.Split(strings.TrimSpace(string(out)), "\n")
		if len(linesForPrint) > 0 {
			fmt.Printf("Video Title: %q\n", strings.TrimSpace(linesForPrint[0]))
		} else {
			fmt.Println("Video Title: <unknown>")
		}

		// fmt.Printf("Video Title: %q\n", string(out))
		if err != nil {
			fmt.Fprintf(os.Stderr, "yt-dlp failed for %s:\n%s\n", url, string(out))
			logF.WriteString(fmt.Sprintf("%s | [FAILED: metadata]\n", url))
			continue
		}

		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		if len(lines) < 2 {
			fmt.Fprintf(os.Stderr, "Bad output from yt-dlp for %s:\n%s\n", url, string(out))
			logF.WriteString(fmt.Sprintf("%s | [FAILED: metadata]\n", url))
			continue
		}
		title := strings.TrimSpace(lines[0])
		canonicalURL := strings.TrimSpace(lines[1])

		// Step 2: Download
		dlCmd := exec.Command(ytDlpBinary,
			"-f", `bv*[ext=mp4]+ba[ext=m4a]/b[ext=mp4] / b`,
			"--merge-output-format", "mp4",
			"-o", fmt.Sprintf("%s/%s", videoDir, outputTemplate),
			canonicalURL)

		timestamp := time.Now().Format("2006-01-02 15:04:05")

		// Step 3a: Log Failure
		if err := dlCmd.Run(); err != nil {
			logEntry := fmt.Sprintf("[%s] %s | FAILED | %s\n", timestamp, canonicalURL, title)
			logF.WriteString(logEntry)
			continue
		}

		// Step 3b: Log success
		logEntry := fmt.Sprintf("[%s] %s | OK | %s\n", timestamp, canonicalURL, title)
		logF.WriteString(logEntry)
		successCount++

		bar.Add(1)

	}

	// Step 4: Clear urls.txt ONLY after full run
	if err := os.WriteFile(urlsFile, nil, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to clear urls.txt: %v\n", err)
	} else {
		fmt.Printf("Processed %d URLs. urls.txt cleared.\n", successCount)
	}
}
