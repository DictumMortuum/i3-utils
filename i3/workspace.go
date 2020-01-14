package i3

import (
	"fmt"
	"go.i3wm.org/i3/v4"
)

func MoveWorkspace(workspace, display string) error {
	command := fmt.Sprintf("workspace %s; move workspace to %s", workspace, display)
	_, err := i3.RunCommand(command)
	return err
}

func MoveContainer(workspaceNum int64) error {
	command := fmt.Sprintf("move container to workspace %d", workspaceNum)
	_, err := i3.RunCommand(command)
	return err
}

func SetCurrentWorkspace(workspaceNum int64) error {
	command := fmt.Sprintf("workspace %d", workspaceNum)
	_, err := i3.RunCommand(command)
	return err
}
