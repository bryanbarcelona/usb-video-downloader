# 📥 USB Video Downloader

> **A self-contained, portable video downloader for educational content — runs directly from your USB stick with zero dependencies.**

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
5. Logs every attempt (success or failure) to `download_log.txt` in **append mode** — so you can retry failed downloads later with other tools if needed.
6. **Clears `urls.txt` only after a full run** — so you never lose your queue if the program crashes mid-download.

> ✅ No internet? No problem — once `yt-dlp.exe` is updated, it can often still download if you've used it on that site before.  
> 🔒 All data stays on your USB — nothing is uploaded or shared.

---

## 📁 Required Files (Place Next to the Executable)

Your USB folder should contain **exactly these 3 files**:

```
your-usb-folder/
├── usb-dl.exe          ← The downloader
├── yt-dlp.exe          ← Official yt-dlp Windows binary
└── urls.txt            ← Your list of video URLs (one per line)
```


> 💡 **Tip**: You can name the folder anything you like — e.g., `Staging`, `VideoCollector`, etc.

The following will be **auto-created** on first run:

- `download_videos/` — where videos are saved
- `download_log.txt` — persistent log of all download attempts

---

## 📝 How to Use

1. **Download the latest `dist/` folder** from this repo (or [grab the ZIP]() if available).
2. Copy the 3 files (`usb-dl.exe`, `yt-dlp.exe`, `urls.txt`) to your USB stick.
3. Open `urls.txt` in any text editor and paste **one video URL per line**, like:

```
https://www.youtube.com/watch?v=dQw4w9WgXcQ

https://www.youtube.com/watch?v=jNQXAC9IVRw
```
  
4. Double-click `usb-dl.exe` and watch the magic happen!
5. After completion, move videos from `download_videos/` to your organized folders.

> ⚠️ **Important**: The program **empties `urls.txt` only after all URLs are processed** — so if it crashes, your list is safe for the next run.

---

## 🌐 Supported Sites

This tool supports **1000+ sites** — including YouTube, Vimeo, Twitch, Dailymotion, and many more!

Since it uses [`yt-dlp`](https://github.com/yt-dlp/yt-dlp) under the hood (via Go’s `os/exec`), it inherits all of `yt-dlp`’s capabilities and frequent updates.

👉 Full list of supported sites: [yt-dlp Supported Sites](https://github.com/yt-dlp/yt-dlp/blob/master/supportedsites.md)

> We use `yt-dlp` as a subprocess (not a native Go library) because video sites change constantly — and `yt-dlp` is actively maintained by a large community. This ensures long-term reliability without reinventing the wheel.
