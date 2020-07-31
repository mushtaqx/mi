package utils_test

import (
	"fmt"
	"testing"

	"github.com/mushtaqx/migo/utils"
)

func TestFileExist(t *testing.T) {
	file := utils.FileExist("../data/migrations", "migration")
	if file == nil {
		t.Fatal("File not found...")
		return
	}

	fmt.Printf("File %s, exists!\n", file.Name())
}
