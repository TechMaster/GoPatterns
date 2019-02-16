package Rect

import (
	"fmt"
)

type Rect struct {
	width  int
	height int
}

func NewRect(width int, height int) *Rect{
	var rect = new(Rect)
	rect.width = width
	rect.height = height
	return rect
}

func (r *Rect) Area() int {
	return r.width * r.height
}

func (r Rect) Perimeter() int {
	return (r.width + r.height) * 2
}

/*
hãy thay đổi func (r Rect) Scale (level int) {
sang 		 func (r *Rect) Scale (level int) {
để thấy sử khác biệt khi in fmt.Println(r.Info())

nếu method
*/
func (r Rect) Scale (level int) {
	fmt.Printf("%p\n", &r)
	r.width = r.width * level
	r.height = r.height * level
}

func (r Rect) Info() string {
	return fmt.Sprintf("width = %d, height = %d", r.width, r.height)
}