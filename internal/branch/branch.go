package branch

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

func get(ctx context.Context) ([]string, error) {
	output, err := exec.CommandContext(ctx, "git", "for-each-ref", "refs/heads", "--format=%(refname:short)").Output()
	if err != nil {
		return nil, err
	}
	// TrimRight to prevent an empty branch name in the final slice of branch names
	return strings.Split(strings.TrimRight(string(output), " \n"), "\n"), nil
}

func GetDifferent(ctx context.Context, from string) ([]string, error) {
	brnchs, err := get(ctx)
	if err != nil {
		return nil, err
	}
	log.Printf("%d branches: %v", len(brnchs), brnchs)

	var same []string
	for _, brnch := range brnchs {
		if brnch == from {
			continue // don't include branch we are diffing against
		}
		err := exec.CommandContext(ctx, "git", "diff", from+".."+brnch, "--exit-code").Run()
		if err != nil {
			exitErr, ok := err.(*exec.ExitError)
			if !ok {
				return nil, fmt.Errorf("diffing branch %s: %v", brnch, err)
			}
			if exitErr.ExitCode() > 1 {
				return nil, fmt.Errorf("diffing branch %s: output was %s - %v", brnch, exitErr.Stderr, exitErr)
			}
			log.Printf("Branch %s is different", brnch)
			continue
		}
		same = append(same, brnch)
	}
	log.Printf("%d branches same as %s: %v", len(same), from, same)
	return same, nil
}
