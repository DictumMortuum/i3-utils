package i3

import (
  "fmt"
  "go.i3wm.org/i3"
)

func MoveWorkspace(workspace, display string) error {
  command := fmt.Sprintf("workspace %s; move workspace to %s", workspace, display)
  _, err := i3.RunCommand(command)

  if err != nil {
    return err
  }

  return nil
}
