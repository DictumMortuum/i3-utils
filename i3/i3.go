package i3

import (
	"go.i3wm.org/i3/v4"
)

func EmptyWorkspace(i int) bool {
	return i == 0
}

func AnyScreen(i int) bool {
	return i != 1
}

func SameScreen(i int) bool {
	return i == 2
}

func Next(fn func(int64) error, skip func(int) bool) {
	var i int64
	ws, _ := i3.GetWorkspaces()
	hm := createWorkspaceHashmap(ws)
	pos := GetFocused(ws).Num
	length := int64(len(hm))

	for i = pos + 1; true; i++ {
		j := i % length

		if skip(hm[j]) {
			fn(j)
			break
		}
	}
}

func Prev(fn func(int64) error, skip func(int) bool) {
	var i int64
	ws, _ := i3.GetWorkspaces()
	hm := createWorkspaceHashmap(ws)
	pos := GetFocused(ws).Num
	length := int64(len(hm))

	for i = pos - 1; true; i-- {
		if i < 0 {
			i = length - 1
		}

		j := i % length

		if skip(hm[j]) {
			fn(j)
			break
		}
	}
}
