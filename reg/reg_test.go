package reg

import (
	"fmt"
	"os"
	"testing"
)

const filename = "index.hbs"

func TestReg(t *testing.T) {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	result := FindMatchesInFile(file)
	fmt.Println(len(result))
	for _, r := range result {
		fmt.Println(r)
	}
}
