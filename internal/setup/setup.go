package setup

import (
	"fmt"
	"os"
	"path/filepath"
)

func Run() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Cannot determine home directory: ", err)
		os.Exit(1)
	}

	lgdRoot := filepath.Join(home, "lgd")

	folders := []string{
		"bin",
		"projects",
		"cache",
	}

	for _, f := range folders {
		path := filepath.Join(lgdRoot, f)

		info, err := os.Stat(path)
		if err == nil && info.IsDir() {
			fmt.Printf("Folder %s already exists\n", path)
			continue
		}

		// mkdir -p
		if err := os.MkdirAll(path, 0755); err != nil {
			fmt.Println("Error creating folder: ", path, err)
			os.Exit(1)
		}
	}

	configPath := filepath.Join(lgdRoot, "config.toml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		defaultConfig := `[paths]
root = "~/lgd"
bin = "~/lgd/bin"
projects = "~/lgd/projects"

[defaults]
python = python3
`
		if err := os.WriteFile(configPath, []byte(defaultConfig), 0644); err != nil {
			fmt.Println("Error writing config.toml: ", err)
			os.Exit(1)
		} else {
			fmt.Println("Config file already exists: ", configPath)
		}
	}

	fmt.Println("LGD setup complete at ", lgdRoot)
}

