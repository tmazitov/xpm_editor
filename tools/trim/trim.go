package trim

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	xpm "github.com/tmazitov/xpm_editor.git/pkg/xpm_file"
)

type TrimTool struct{}

func (t *TrimTool) Execute(filePaths []string) error {

	var (
		file      *xpm.XpmFile
		files     []*xpm.XpmFile
		trimCount int
	)
	// 1. Load xmp files
	for _, filePath := range filePaths {
		file = xpm.NewXpmFile(filePath)
		if err := file.Read(); err != nil {
			return err
		}
		files = append(files, file)
	}

	// 2. Trim count
	trimCount = enterTrimCount()

	// 3. Trim files
	for _, file := range files {
		fmt.Print(file.FilePath)
		if trimCount > 0 {
			if err := file.Trim(trimCount); err != nil {
				return err
			}
		} else if trimCount == -1 {
			if err := file.TrimAuto(); err != nil {
				return err
			}
		}
		if err := file.Write(); err != nil {
			return err
		}
		fmt.Println(" done!")
	}
	return nil
}

func enterTrimCount() int {
	var result int

	fmt.Println("\nEnter the trim count:")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {

			input := scanner.Text()
			if input == "a" {
				result = -1
				break
			}
			if isPositiveInteger(input) {
				result, _ = strconv.Atoi(input)
				break
			}

			fmt.Println("\nEnter the trim count:")
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
