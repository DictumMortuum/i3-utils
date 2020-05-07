package i3

import (
	"go.i3wm.org/i3/v4"
)

func GetFocused(ws []i3.Workspace) i3.Workspace {
	var ret i3.Workspace

	for _, w := range ws {
		if w.Focused {
			ret = w
			break
		}
	}

	return ret
}

func getMaxWorkspace(ws []i3.Workspace) int64 {
	max := ws[0].Num

	for _, w := range ws {
		if w.Num > max {
			max = w.Num
		}
	}

	return max
}

func createWorkspaceHashmap(ws []i3.Workspace) [100]int {
	hm := [100]int{}
	output := GetFocused(ws).Output

	for _, w := range ws {
		hm[w.Num] = 1

		if w.Output == output {
			hm[w.Num] = 2
		}
	}

	return hm
}
