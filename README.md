# bing-wallpaper

Fetches the daily Bing wallpaper and sets it as your desktop background. Written in Go.

## Usage

Build it with `go build`, then:

| Command | What it does |
|---|---|
| `bing-wallpaper` | Download today's wallpaper and set it |
| `bing-wallpaper -d 3` | Use the wallpaper from 3 days ago (0 to 7) |
| `bing-wallpaper download <dir>` | Only download the image to `<dir>` |
| `bing-wallpaper --daemon` | Run a tray icon where you can pick a wallpaper from the last week |
| `bing-wallpaper install` | Move the binary to `~/.local/bin` and set up a service (Linux and macOS) |

## Configuration

The first run creates `~/.bing-wallpaper.yaml`. Set `auto_update: false` to stop the wallpaper changing automatically.

This project is no longer maintained.
