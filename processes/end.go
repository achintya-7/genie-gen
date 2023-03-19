package processes

import (
	"fmt"
	"os/exec"
)

func End() error {
	cmd := exec.Command("go", "mod", "tidy")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}
