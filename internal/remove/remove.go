package remove

import (
	"fmt"
	"os"
	"path/filepath"
)

func Run(repoName string) error {
	//paths
        home, err := os.UserHomeDir()
        if err != nil {
                fmt.Println("Cannot determine home directory: ", err)
                os.Exit(1)
        }

        lgdRoot := filepath.Join(home, "lgd")
        projectsPath := filepath.Join(lgdRoot, "projects")
        repoPath := filepath.Join(projectsPath, repoName)

	return os.RemoveAll(repoPath)
}

