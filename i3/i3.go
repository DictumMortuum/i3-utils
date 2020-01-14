package i3

import (
	"go.i3wm.org/i3/v4"
)

func Next(fn func(int64) error) {
	var i int64
	ws, _ := i3.GetWorkspaces()
	hm := createWorkspaceHashmap(ws)
	pos := GetFocused(ws).Num
	length := int64(len(hm))

	for i = pos + 1; i < length; i++ {
		if hm[i] != 1 {
			fn(i)
			break
		}
	}
}

func Prev(fn func(int64) error) {
	var i int64
	ws, _ := i3.GetWorkspaces()
	hm := createWorkspaceHashmap(ws)
	pos := GetFocused(ws).Num

	for i = pos - 1; i >= 0; i-- {
		if hm[i] != 1 {
			fn(i)
			break
		}
	}
}
