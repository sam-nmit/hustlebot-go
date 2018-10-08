package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	binDirectory               = "bin"
	applicationSourceDirectory = "src"
	pluginDirectory            = "plugins"
)

type buildOS struct {
	GOOS, GOARCH string
}

var buildOptions = map[string]buildOS{
	"linux":   {GOOS: "linux", GOARCH: "arm"},
	"linux64": {GOOS: "linux", GOARCH: "rm64"},
}

var (
	target = flag.String("target", "linux", "target compile platform")
	goos   = flag.String("goos", "", "goos")
	goarch = flag.String("goarch", "", "goarch")
)

func main() {
	flag.Parse()

	var b buildOS

	if val, ok := buildOptions[*target]; ok {
		b = val
	} else {
		b = buildOS{
			GOOS:   *goos,
			GOARCH: *goarch,
		}
	}

	outputDir := filepath.Join(binDirectory, concat(b.GOOS, ".", b.GOARCH))
	outputPluginDir := filepath.Join(outputDir, "plugins")

	os.MkdirAll(outputDir, os.ModePerm)
	os.Mkdir(outputPluginDir, os.ModePerm)

	env := []string{
		concat("GOOS=", b.GOOS),
		concat("GOARCH=", b.GOARCH),
		concat("GOPATH=", os.Getenv("GOPATH")),
		concat("PATH=", os.Getenv("PATH")),
	}

	log.Println("Compiling executable...")

	if out, err := buildSource(
		filepath.Join(applicationSourceDirectory, "main.go"),
		filepath.Join(outputDir, "hustlebot"),
		"default", env); err != nil {

		log.Println(out)
		log.Fatalln(err)
	}

	log.Println("Compiling plugins...")

	pluginFolders, _ := ioutil.ReadDir(pluginDirectory)
	for _, folder := range pluginFolders {
		if !folder.IsDir() {
			continue
		}
		log.Println(">", folder.Name())
		if out, err := buildSource(
			filepath.Join(pluginDirectory, folder.Name(), "plugin.go"),
			filepath.Join(outputDir, "plugins", folder.Name()),
			"plugin", env); err != nil {

			log.Println(out)
			log.Fatalln(err)
		}
	}
	log.Println("done.")
}

func buildSource(sourcefile, outfile, buildmode string, env []string) (string, error) {
	cmd := exec.Command("go", "build", "-o", outfile, "-buildmode", buildmode, sourcefile)
	cmd.Env = env

	out, err := cmd.CombinedOutput()
	return string(out), err
}

func concat(src ...string) string {
	return strings.Join(src, "")
}
