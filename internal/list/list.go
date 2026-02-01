package list

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
	"github.com/ni3wyspany/lgd/meta"
)

func readMeta(repoPath string) (*meta.Meta, error) {
	metaPath := filepath.Join(repoPath, "meta.toml")

	data, err := os.ReadFile(metaPath)
	if err != nil {
		return nil, err
	}

	var m meta.Meta
	if err := toml.Unmarshal(data, &m); err != nil {
		return nil, err
	}

	return &m, nil
}

func Run() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	projectsDir := filepath.Join(home, "lgd", "projects")

	entries, err := os.ReadDir(projectsDir)
	if err != nil {
		return err
	}

	fmt.Println("Enabled repos:")

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		repoPath := filepath.Join(projectsDir, entry.Name())
		metaPath := filepath.Join(repoPath, "meta.toml")

		if _, err := os.Stat(metaPath); err == nil {
			m, err := readMeta(repoPath)
			if err != nil {
				continue
			}

			if m.Project.Enabled {
				fmt.Printf(
					"- %s (%s)\n",
					m.Project.Name,
					m.Project.Language,
				)
			}
		}
	}

	return nil
}

