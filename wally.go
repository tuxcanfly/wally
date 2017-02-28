package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

var (
	url      = flag.String("url", "", "url of the wallpaper")
	path     = flag.String("path", "Pictures", "dir for wallpaper")
	force    = flag.Bool("force", false, "force download")
	filename = flag.String("filename", "wally.jpg", "filename for wallpaper")
	plugin   = flag.String("plugin", "bing", "plugin to use")
	list     = flag.Bool("list", false, "list plugins and exit")
)

func absfilepath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(usr.HomeDir, *path, *filename)
}

func init() {
	flag.Parse()

	pm := &pluginManager{}
	pm.register("bing", newBingPlugin())

	if *list {
		log.Println(pm.list())
		os.Exit(0)
	}

	fi, err := os.Stat(absfilepath())
	if err != nil {
		log.Fatal(err)
	}

	recent := fi.ModTime().After(time.Now().Add(-24 * time.Hour))
	if recent && !*force {
		os.Exit(0)
	}
	if *url != "" {
		return
	}

	p, err := pm.get(*plugin)
	if err != nil {
		log.Fatal(err)
	}
	*url, err = p.URL()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	resp, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	out, err := os.Create(absfilepath())
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	io.Copy(out, resp.Body)
}
