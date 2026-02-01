package disable

import (
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

func Run(repoName string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	repoPath := filepath.Join(home, "lgd", "projects", repoName)

	m, err := readMeta(repoPath)
	if err != nil {
		return err
	}

	if m.Project.Language == "python" && m.Env.Type == "venv" {
		_ = os.RemoveAll(m.Env.Path)
	}

	return os.Remove(filepath.Join(repoPath, "meta.toml"))
}

