package c

import (
	"github.com/danielborowski/cb-code-runner/cmd"
	"github.com/danielborowski/cb-code-runner/util"
	"path/filepath"
)

func Run(files []string, stdin string) (string, string, error) {
	workDir := filepath.Dir(files[0])
	binName := "a.out"

	sourceFiles := util.FilterByExtension(files, "c")
	args := append([]string{"clang", "-o", binName, "-lm"}, sourceFiles...)
	stdout, stderr, err := cmd.Run(workDir, args...)
	if err != nil || stderr != "" {
		return stdout, stderr, err
	}

	binPath := filepath.Join(workDir, binName)
	return cmd.RunStdin(workDir, stdin, binPath)
}