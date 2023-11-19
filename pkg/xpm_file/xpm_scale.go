package xpm

import "strings"

func scaleRow(row string, scaleCount int) (string, error) {

	var (
		counter int
		result  strings.Builder
		err     error
	)

	for _, rowPixel := range row {
		counter = 0
		for counter < scaleCount {
			if _, err = result.WriteRune(rowPixel); err != nil {
				result.Reset()
				return "", err
			}
			counter++
		}
	}

	return result.String(), nil
}

func (x *XpmFile) Scale(scaleCount int) error {

	var (
		originImage []string
		convString  string
		convCounter int
		err         error
	)

	originImage = x.image
	x.image = []string{}
	for _, imageRow := range originImage {
		convCounter = 0
		convString, err = scaleRow(imageRow, scaleCount)
		if err != nil {
			return err
		}
		for convCounter < scaleCount {
			x.image = append(x.image, convString)
			convCounter++
		}
	}
	x.columns = x.columns * scaleCount
	x.rows = x.rows * scaleCount
	return nil
}
