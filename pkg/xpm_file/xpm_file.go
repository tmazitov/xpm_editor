package xpm

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type XpmFile struct {
	FilePath      string
	colors        []*xpmColor
	image         []string
	header        []string
	footer        []string
	columns       int
	rows          int
	colorCount    int
	charsPerPixel int
}

func NewXpmFile(filePath string) *XpmFile {
	return &XpmFile{
		FilePath: filePath,
		colors:   nil,
		image:    nil,
	}
}

func (xpm *XpmFile) FillDetails(row string) error {
	var (
		row_parts []string
		err       error
	)

	row = strings.ReplaceAll(row, ",", "")
	row = strings.ReplaceAll(row, "\"", "")
	row_parts = strings.Split(row, " ")
	xpm.columns, err = strconv.Atoi(row_parts[0])
	if err != nil {
		return err
	}
	xpm.rows, err = strconv.Atoi(row_parts[1])
	if err != nil {
		return err
	}
	xpm.colorCount, err = strconv.Atoi(row_parts[2])
	if err != nil {
		return err
	}
	xpm.charsPerPixel, err = strconv.Atoi(row_parts[3])
	if err != nil {
		return err
	}
	return nil
}

func (xpm *XpmFile) Read() error {

	var (
		raw           []byte
		rawString     string
		rows          []string
		err           error
		row_is_color  bool
		row_is_header bool
		foundColors   bool
		foundHeader   bool
	)

	raw, err = os.ReadFile(xpm.FilePath)
	if err != nil {
		return err
	}

	rawString = string(raw)
	rows = strings.Split(rawString, "\n")
	for _, row := range rows {
		if (strings.Contains(row, "/*") && strings.Contains(row, "*/")) || row == "" {
			continue
		}
		row_is_color = isColor(row)
		row_is_header = isHeader(row)
		if !foundColors && row_is_color {
			foundColors = row_is_color
		}

		// Find the image details
		if !foundHeader && row_is_header {
			foundHeader = row_is_header
		} else if foundHeader && !row_is_header {
			if err = xpm.FillDetails(row); err != nil {
				return err
			}
			foundHeader = false
			continue
		}

		if !foundColors && !row_is_color {
			xpm.header = append(xpm.header, row)
		} else if foundColors && row_is_color {
			xpm.colors = append(xpm.colors, NewXpmColor(row))
		} else if row == "};" {
			xpm.footer = append(xpm.footer, "};")
		} else if foundColors && !row_is_color {
			row = strings.ReplaceAll(row, "\"", "")
			row = strings.ReplaceAll(row, ",", "")
			xpm.image = append(xpm.image, row)
		}
	}

	for _, imageRow := range xpm.image {
		fmt.Println(imageRow)
	}

	return nil
}

func (x *XpmFile) Write() error {
	f, err := os.Create(x.FilePath)
	if err != nil {
		return err
	}
	for _, headerRow := range x.header {
		_, err := f.WriteString(headerRow + "\n")
		if err != nil {
			return err
		}
	}
	_, err = f.WriteString(fmt.Sprintf("\"%d %d %d %d\",\n", x.columns, x.rows, x.colorCount, x.charsPerPixel))
	if err != nil {
		return err
	}
	for _, color := range x.colors {
		_, err := f.WriteString(color.ToString() + "\n")
		if err != nil {
			return err
		}
	}
	for _, imageRow := range x.image {
		_, err := f.WriteString("\"" + imageRow + "\",\n")
		if err != nil {
			return err
		}
	}
	for _, footerRow := range x.footer {
		_, err = f.WriteString(footerRow + "\n")
		if err != nil {
			return err
		}
	}

	f.Sync()
	return nil
}

func isHeader(row string) bool {
	return strings.Contains(row, "static char")
}
