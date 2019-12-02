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
