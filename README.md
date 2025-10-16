# 📥 USB Video Downloader

> **A self-contained, portable video downloader — runs directly from your USB stick with zero dependencies.**

---

## 🧠 Why This Exists

This is a passion project born during my college years. I was constantly downloading educational videos from platforms like YouTube for offline study, but tired of manual downloads and the hassle of managing files across different computers.

Originally written in Python, I later realized that not every campus PC had Python installed — making the script unusable on shared machines. So I rewrote it in **Go** to produce a single, standalone executable that runs anywhere, **without requiring any interpreter or installation**.

Now I keep everything — the downloader, `yt-dlp`, and my download queue — right on my USB stick. Plug in, run, and collect videos offline. Perfect for students, researchers, or lifelong learners.

---

## 🚀 How It Works

1. **Auto-updates `yt-dlp`** on every run (to handle ever-changing video sites).
2. Reads video URLs from `urls.txt` (one per line).
3. Fetches video **title** and **canonical URL** using `yt-dlp`.
4. Downloads videos into a folder called `download_videos/` (auto-created if missing).
    - Every file is saved as **MP4** (H.264/AAC) for maximum compatibility—no post-processing needed. 
5. Logs every attempt (success or failure) to `download_log.txt` in **append mode** — so you can retry failed downloads later with other tools if needed.
6. **Clears `urls.txt` only after a full run** — so you never lose your queue if the program crashes mid-download.

---

## 📁 Installation

This section covers how to set up the USB Video Downloader on your USB drive or preferred location.

1. [Download the ZIP folder](https://github.com/bryanbarcelona/usb-video-downloader/releases/download/v1.0.0/usb-video-downloader.zip).
2. Extract the ZIP to your USB drive or another chosen location.
3. Optionally, rename the extracted "usb-video-downloader" folder to something like `Staging`, `VideoCollection` or whatever you like.
4. Ensure the folder contains exactly these three files:

```
your-folder/
├── usb-dl.exe          ← The downloader executable
├── yt-dlp.exe          ← Official yt-dlp Windows binary
└── urls.txt            ← File for listing video URLs (one per line)
```

> 💡 **Tip**: The folder name is flexible—choose any name that suits your needs, e.g., `Staging` or `VideoCollector`.

The following will be **auto-created** on first run:
- `download_videos/` — Directory where downloaded videos are saved
- `download_log.txt` — Log file for tracking download attempts

---

## 📝 How to Use

This section explains how to use the USB Video Downloader after installation.

1. Open `urls.txt` in a text editor (e.g., Notepad).
2. Add one video URL per line, for example:

```
https://www.youtube.com/watch?v=dQw4w9WgXcQ
https://www.youtube.com/watch?v=jNQXAC9IVRw
```

3. Save `urls.txt`.
4. Double-click `usb-dl.exe` to start downloading.
5. Check the `download_videos/` folder for your downloaded videos.
6. Move videos from `download_videos/` to your preferred organized folders.
7. Review `download_log.txt` for a record of download attempts.

> ⚠️ **Important**: The program clears `urls.txt` only after all URLs are successfully processed. If it crashes, your URL list remains intact for the next run.

---

## 🌐 Supported Sites

This tool supports **1000+ sites** — including YouTube, Vimeo, Twitch, Dailymotion, and many more!

Since it uses [`yt-dlp`](https://github.com/yt-dlp/yt-dlp) under the hood (via Go’s `os/exec`), it inherits all of `yt-dlp`’s capabilities and frequent updates.

👉 Full list of supported sites: [yt-dlp Supported Sites](https://github.com/yt-dlp/yt-dlp/blob/master/supportedsites.md)

> We use `yt-dlp` as a subprocess (not a native Go library) because video sites change constantly — and `yt-dlp` is actively maintained by a large community. This ensures long-term reliability without reinventing the wheel.

## 
