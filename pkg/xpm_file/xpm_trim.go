package xpm

import "fmt"

func (x *XpmFile) Trim(trimCount int) error {
	x.image = x.image[trimCount : x.rows-trimCount]
	for index, imageRow := range x.image {
		x.image[index] = imageRow[trimCount : x.rows-trimCount]
	}
	x.rows = x.rows - trimCount*2
	x.columns = x.columns - trimCount*2
	return nil
}

func isStringSingleChar(line string) bool {
	for _, char := range line {
		if rune(line[0]) != char {
			return false
		}
	}
	return true
}

func calcTrimLine(line string) (int, int) {
	var (
		left          int  = 0
		right         int  = 0
		body_is_found bool = false
	)

	for _, char := range line {

		if char != rune(line[0]) && !body_is_found {
			body_is_found = true
		}
		if !body_is_found && char == rune(line[0]) {
			left++
		}
		if body_is_found && char == rune(line[0]) {
			right++
		}
	}

	return left, right
}

func (x *XpmFile) TrimAuto() error {

	var (
		image_single_char bool
		trim_head         int  = 0
		body_is_found     bool = false
		trim_foot         int  = 0
		trim_right        int  = 2147483647
		trim_left         int  = 2147483647
	)

	for _, imageRow := range x.image {
		image_single_char = isStringSingleChar(imageRow)
		if !image_single_char && !body_is_found {
			body_is_found = true
		}
		if image_single_char && !body_is_found {
			trim_head++
		} else if image_single_char && body_is_found {
			trim_foot++
		} else if !image_single_char && body_is_found {
			left, right := calcTrimLine(imageRow)
			if right < trim_right {
				trim_right = right
			}
			if left < trim_left {
				trim_left = left
			}
		}
	}
	fmt.Println(trim_head, trim_foot)
	x.image = x.image[trim_head : x.rows-trim_foot-1]
	for index, imageRow := range x.image {
		x.image[index] = imageRow[trim_left : x.columns-trim_right]
	}
	x.rows = x.rows - trim_head - trim_foot
	x.columns = x.columns - trim_left - trim_right
	return nil
}
