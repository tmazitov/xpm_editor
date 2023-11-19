package xpm

func (x *XpmFile) Trim(trimCount int) error {
	x.image = x.image[trimCount : x.rows-trimCount]
	for index, imageRow := range x.image {
		x.image[index] = imageRow[trimCount : x.rows-trimCount]
	}
	x.rows = x.rows - trimCount*2
	x.columns = x.columns - trimCount*2
	return nil
}
