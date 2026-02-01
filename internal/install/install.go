package install

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func Run(owner string, repoName string) {
	
	//build repo url
	url := "https://github.com/" + owner + "/" + repoName + ".git"

	//check home dir
	home, err := os.UserHomeDir()
        if err != nil {
                fmt.Println("Cannot determine home directory: ", err)
                os.Exit(1)
        }

	//paths
	lgdRoot := filepath.Join(home, "lgd")
	projectsPath := filepath.Join(lgdRoot, "projects")
	projectPath := filepath.Join(projectsPath, repoName)
	
	//is installed
	if _, err := os.Stat(projectPath); err == nil {
		fmt.Println("Project already installed: ", owner+"/"+repoName)
		fmt.Println("If you didn't enable this repo yet use command: lgd enable <owner/repo>")
		os.Exit(1)
	}

	//clone repo
	cmd := exec.Command("git", "clone", url, projectPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to clone repo: ", err)
		os.Exit(1)
	}

	fmt.Printf("Installed %s/%s at %s\n", owner, repoName, projectPath)
	fmt.Printf("Run command: lgd enable %s/%s  to sort and enable this repo\n", owner, repoName)
}

