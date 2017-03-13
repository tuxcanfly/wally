wally
======

wally is a tiny command to download a wallpaper

it can be added to autostart to get new wallpapers

    wally -url <url>

by default, the command fetches a wallpaper to $HOME/Pictures/wally.jpg
it only downloads if the wallpaper is a day (24h) old

in addition to url, it can be configured to use plugins from different sources

plugins
=======

bing (default): fetches bing picture of the day

usage
=====

`wally --help`:

	Usage of wally:
	-filename string
			filename for wallpaper (default "wally.jpg")
	-force
			force download
	-list
			list plugins and exit
	-path string
			dir for wallpaper (default "Pictures")
	-plugin string
			plugin to use (default "bing")
	-url string
			url of the wallpaper
