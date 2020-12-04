package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/thatisuday/commando"
)

var (
	defaultGravatarSize = 160
	gravatarTypes       = []string{"identicon", "robohash", "monsterid", "wavatar", "retro"}
)

func main() {
	commando.
		SetExecutableName("get-random-gravatar").
		SetVersion("v1.0.0").
		SetDescription("A command-line tool to download and save a random Gravatar image on current folder.")

	commando.
		Register(nil).
		AddFlag("size,s", "size of the gravatar in pixels (minimum: 1) (maximum: 2048)", commando.Int, defaultGravatarSize).
		AddFlag("type,t", fmt.Sprint("type of the gravatar (available: ", strings.Join(gravatarTypes[:], ", "), ")"), commando.String, gravatarTypes[0]).
		AddFlag("verbose,V", "display log information", commando.Bool, false).
		SetAction(defaultActionHandler)

	commando.Parse(nil)
}

func defaultActionHandler(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	hash := generateRandomMD5()
	fileName := fmt.Sprint(hash, ".jpg")
	size := int(math.Max(1, math.Min(float64(flags["size"].Value.(int)), 2048)))
	gravatarType := flags["type"].Value.(string)
	fileContent := getGravatarImage(hash, size, gravatarType)
	file := createFileWithContent(fileName, fileContent)
	path, _ := filepath.Abs(file.Name())
	fmt.Println("Gravatar saved to", path)
}

func createFileWithContent(fileName string, content io.ReadCloser) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, content)
	if err != nil {
		log.Fatal(err)
	}
	defer content.Close()

	return file
}

func getGravatarImage(hash string, size int, gravatarType string) io.ReadCloser {
	if !isGravatarTypeValid(gravatarType) {
		fmt.Println(
			fmt.Sprintf(
				"Error: Gravatar type '%v' is invalid. Please use one of the following: %v.",
				gravatarType, strings.Join(gravatarTypes[:], ", "),
			),
		)
		os.Exit(0)
	}

	gravatarURL := fmt.Sprint(
		"https://www.gravatar.com/avatar/", hash,
		"?size=", size,
		"&default=", gravatarType,
		"&forcedefault=y",
	)

	response, err := http.Get(gravatarURL)
	if err != nil {
		log.Fatal(err)
	}

	return response.Body
}

func isGravatarTypeValid(gravatarType string) bool {
	for _, v := range gravatarTypes {
		if v == gravatarType {
			return true
		}
	}

	return false
}

func generateRandomMD5() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
