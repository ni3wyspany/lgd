package enable

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"github.com/ni3wyspany/lgd/meta"
	"github.com/pelletier/go-toml/v2"
)

func fileExists(path string) bool {
        _, err := os.Stat(path)
        return err == nil
}

func Run(repoName string, owner string) {
	//paths
	home, err := os.UserHomeDir()
	if err != nil {
                fmt.Println("Cannot determine home directory: ", err)
                os.Exit(1)
        }

	lgdRoot := filepath.Join(home, "lgd")
        projectsPath := filepath.Join(lgdRoot, "projects")
	projectPath := filepath.Join(projectsPath, repoName)

	//language detection
	if _, err := os.Stat(projectPath); err == nil {
                fmt.Println("Project is installed. Enabling and securing...")
		
		hasPython := false
		hasGo := false

		if fileExists(filepath.Join(projectPath, "pyproject.toml")) ||
		fileExists(filepath.Join(projectPath, "setup.py")) ||
		fileExists(filepath.Join(projectPath, "requirements.txt")) {
			hasPython = true
		}

		if fileExists(filepath.Join(projectPath, "go.mod")) {
			hasGo = true
		}

		language := "mixed"
		if hasPython && !hasGo {
			language = "python"
		} else if hasGo && !hasPython {
			language = "go"
		}

		//meta.toml

		var venvType string
		var venvPath string

		if language == "python" {
			venvDir := filepath.Join(projectPath, ".venv")

			venvType = "venv"
			venvPath = venvDir
			
			//env
			cmd := exec.Command("python3", "-m", "venv", venvDir)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				fmt.Println("Failed to create env")
				os.Exit(1)
			}
			fmt.Println("Succesfully created env")

		} else {
			venvType = "None"
			venvPath = "None"
		}

		//meta.toml

		metaPath := filepath.Join(projectPath, "meta.toml")
		m := meta.Meta{
			Project: meta.Project{
				Name: repoName,
				Language: language,
				Enabled: true,
			},
			Source: meta.Source{
				Type: "github",
				Owner: owner,
				Repo: repoName,
			},
			Env: meta.Env{
				Type: venvType,
				Path: venvPath,
			},
			Run: meta.Run{
				Cmd: "",
			},
		}

		data, err := toml.Marshal(m)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(metaPath, data, 0644)
		if err != nil {
			panic(err)
		}




        } else {
		fmt.Println("Did you install this repo?\nMake sure to install it with: lgd install <owner/repoName>")
		os.Exit(1)
	}
}

