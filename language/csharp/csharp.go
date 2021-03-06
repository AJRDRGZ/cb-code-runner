package csharp

import (
	"github.com/coderbyte-org/cb-code-runner/cmd"
	"github.com/coderbyte-org/cb-code-runner/util"
	"path/filepath"
)

func Run(files []string, stdin string) (string, string, error, string) {
	workDir := filepath.Dir(files[0])
	binName := "a.exe"

	sourceFiles := util.FilterByExtension(files, "cs")
	args := append([]string{"mcs", "-out:" + binName}, sourceFiles...)
	stdout, stderr, err, duration := cmd.Run(workDir, args...)
	if err != nil || stderr != "" {
		return stdout, stderr, err, duration
	}

	binPath := filepath.Join(workDir, binName)
	return cmd.RunStdin(workDir, stdin, "mono", binPath)
}