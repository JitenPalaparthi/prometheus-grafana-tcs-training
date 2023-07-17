package main

import (
	"flag"
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	_ "github.com/prometheus/client_golang/prometheus/push"
)

var (
	root                string
	dirCount, fileCount uint32
	fileCounter         = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "prom_push_app_fileCount",
		Help: "Number of files that the given path has",
	})

	dirCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "prom_push_app_dirCount",
		Help: "Number of directories that the given path has",
	})
)

func main() {

	flag.StringVar(&root, "path", ".", "--path=./")
	flag.Parse()

	fmt.Println(root)
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			dirCount++
		} else {
			fileCount++
		}
		return nil
	})
	fileCounter.Add(float64(fileCount))
	dirCounter.Add(float64(dirCount))

	fmt.Println("Total number of directories:", dirCount, "\nTotal number of files:", fileCount)
	if err := push.New("http://localhost:9091", "prom-push-app-fc").Collector(fileCounter).Push(); err != nil {
		fmt.Println(err)
	}
	if err := push.New("http://localhost:9091", "prom-push-app-dc").Collector(dirCounter).Push(); err != nil {
		fmt.Println(err)
	}
}
