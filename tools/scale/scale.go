package scale

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	xpm "github.com/tmazitov/xpm_editor.git/pkg/xpm_file"
)

type ScaleTool struct {
	xpmFiles []*xpm.XpmFile
}

func NewTool() *ScaleTool {
	return &ScaleTool{
		xpmFiles: []*xpm.XpmFile{},
	}
}

func (t *ScaleTool) Execute(filePaths []string) error {

	var (
		file       *xpm.XpmFile
		scaleCount int
	)
	// 1. Load xmp files
	for _, filePath := range filePaths {
		file = xpm.NewXpmFile(filePath)
		if err := file.Read(); err != nil {
			return err
		}
		t.xpmFiles = append(t.xpmFiles, file)
	}

	// 2. Scale count
	scaleCount = enterScaleCount()

	// 2. Scale files
	for _, file := range t.xpmFiles {
		fmt.Print(file.FilePath)
		if err := file.Scale(scaleCount); err != nil {
			return err
		}
		if err := file.Write(); err != nil {
			return err
		}
		fmt.Println(" done!")
	}
	return nil
}

func enterScaleCount() int {
	var result int

	fmt.Println("\nEnter the scale count:")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {

			input := scanner.Text()
			if isPositiveInteger(input) {
				result, _ = strconv.Atoi(input)
				break
			}

			fmt.Println("\nEnter the scale count:")
		} else {
			fmt.Println("Error reading input:", scanner.Err())
			break
		}
	}
	return result
}

func isPositiveInteger(s string) bool {
	num, err := strconv.Atoi(s)
	return err == nil && num > 0
}
