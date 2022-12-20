package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"ghost-theme-translator/reg"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()

	fmt.Println(
		"Ghost theme translator\n",
		"We will find every hard-encoded string in your theme and ask you to translate them.\n",
		"Please enter the translation for each string.\n",
		"Press enter to skip a string.\n",
		"Press Ctrl+C to exit.",
	)

	fmt.Println("Working directory:", root)
	err := os.Chdir(root)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scanning files...")
	for _, dir := range dirs {
		files, err := os.ReadDir(dir)
		if err != nil {
			panic(err)
		}

		for _, file := range files {
			if file.IsDir() ||
				!file.Type().IsRegular() ||
				!strings.HasSuffix(file.Name(), extension) {
				continue
			}
			path := filepath.Join(dir, file.Name())
			walkDirCallback(path)
		}
	}

	writeTranslationFile()
}

func walkDirCallback(path string) {
	fmt.Println("Current file: ", path)
	file, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
		return
	}

	result := reg.FindMatchesInFile(file)

	text := string(file)
	for _, match := range result {
		if _, ok := translation[match]; ok { // already translated
			text = strings.ReplaceAll(text, match, fmt.Sprintf(`{{t "%s"}}`, match))
			continue
		}
		fmt.Printf("%s: ", match)
		input := readString()
		if input == "" {
			continue
		}
		translation[match] = input
		text = strings.ReplaceAll(text, match, fmt.Sprintf(`{{t "%s"}}`, match))
	}

	err = os.WriteFile(path, []byte(text), 0644)
	if err != nil {
		log.Println(err)
		return
	}
}

func writeTranslationFile() {
	fmt.Print("translation completed, please enter your language code: ")
	lang := readString()

	fmt.Printf("translation en and %s will be saved and overrided, continue? (y/n): ", lang)
	input := readString()
	if input != "y" {
		fmt.Println("canceled")
		return
	}

	en := map[string]string{}
	for k := range translation {
		en[k] = k
	}

	writeJsonToFile("en.json", en)
	writeJsonToFile(fmt.Sprintf("%s.json", lang), translation)
}

func writeJsonToFile(filename string, data map[string]string) {
	const dirName = "locales"
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		panic(err)
		return
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(filepath.Join(dirName, filename), bytes, 0644)
	if err != nil {
		panic(err)
	}
}
