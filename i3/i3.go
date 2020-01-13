package i3

import (
  "fmt"
	"errors"
  "go.i3wm.org/i3/v4"
)

func MoveWorkspace(workspace, display string) error {
  command := fmt.Sprintf("workspace %s; move workspace to %s", workspace, display)
  _, err := i3.RunCommand(command)
  return err
}

func SetCurrentWorkspace(workspaceNum int64) error {
  command := fmt.Sprintf("workspace %d", workspaceNum)
  _, err := i3.RunCommand(command)
  return err
}

func GetCurrentWorkspaceNumber() (int64, error) {
  ws, err := i3.GetWorkspaces()
  if err != nil {
    return -1, err
  }

  for _, w := range ws {
    if w.Focused {
      return w.Num, nil
    }
  }

  return -1, errors.New("Cant find current workspace")
}

func getFocused(ws []i3.Workspace) i3.Workspace {
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
	output := getFocused(ws).Output

	for _, w := range ws {
		hm[w.Num] += 1

		if w.Output == output {
			hm[w.Num] += 1
		}
	}

	return hm
}

func Next() {
	var i int64
  ws, _ := i3.GetWorkspaces()
	hm := createWorkspaceHashmap(ws)
	pos := getFocused(ws).Num
	length := int64(len(hm))

	for i = pos + 1; i < length; i++ {
		if hm[i] != 1 {
			fmt.Println(hm, i)
			SetCurrentWorkspace(i)
			break
		}
	}
}

func Prev() {
	var i int64
  ws, _ := i3.GetWorkspaces()
	hm := createWorkspaceHashmap(ws)
	pos := getFocused(ws).Num

	for i = pos - 1; i >= 0; i-- {
		if hm[i] != 1 {
			fmt.Println(hm, i)
			SetCurrentWorkspace(i)
			break
		}
	}
}
