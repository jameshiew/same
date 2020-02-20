package git

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

func getBranchNames(ctx context.Context) ([]string, error) {
	output, err := exec.CommandContext(ctx, "git", "for-each-ref", "refs/heads", "--format=%(refname:short)").CombinedOutput()
	if err != nil {
		if len(output) > 0 {
			log.Debugf("Output for `git for-each-ref` error: %s", output)
		}
		return nil, fmt.Errorf("executing git for-each-ref: %v", err)
	}
	// TrimRight to prevent an empty git name in the final slice of git names
	return strings.Split(strings.TrimRight(string(output), " \n"), "\n"), nil
}

// GetSameBranchNames returns a list of branch names for this repository which have no diff with branch `from`
func GetSameBranchNames(ctx context.Context, from string) ([]string, error) {
	brnchs, err := getBranchNames(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting branch names: %w", err)
	}
	log.Infof("%d branches: %v", len(brnchs), brnchs)

	var same []string
	for _, brnch := range brnchs {
		if brnch == from {
			continue // don't include git we are diffing against
		}
		err := exec.CommandContext(ctx, "git", "diff", from+".."+brnch, "--exit-code").Run()
		if err != nil {
			exitErr, ok := err.(*exec.ExitError)
			if !ok {
				return nil, fmt.Errorf("diffing git %s: %w", brnch, err)
			}
			if exitErr.ExitCode() > 1 {
				return nil, fmt.Errorf("diffing git %s: output was %s - %w", brnch, exitErr.Stderr, exitErr)
			}
			log.Infof("Branch %s is different", brnch)
			continue
		}
		same = append(same, brnch)
	}
	log.Infof("%d branches same as %s: %v", len(same), from, same)
	return same, nil
}
