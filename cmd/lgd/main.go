package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ni3wyspany/lgd/internal/setup"
	"github.com/ni3wyspany/lgd/internal/install"
	"github.com/ni3wyspany/lgd/internal/enable"
	"github.com/ni3wyspany/lgd/internal/list"
	"github.com/ni3wyspany/lgd/internal/disable"
	"github.com/ni3wyspany/lgd/internal/remove"
)

func main () {
	if len(os.Args) < 2 {
		fmt.Println("Usage: lgd <command>")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	//setup
	case "--setup":
		fmt.Println("Setup command called")
		setup.Run()

	//install
	case "install":
		//lgd install
		args := os.Args[2:]
		if len(args) < 1 {
			fmt.Println("Usage lgd install <owner/repo>")
			os.Exit(1)
		}
		repoArg := args[0]
		parts := strings.Split(repoArg, "/")
		if len(parts) != 2 {
			fmt.Println("Invalid repo format. Use owner/repo")
			os.Exit(1)
		}
		owner := parts[0]
		repoName := parts[1]

		//download repo
		fmt.Println("Installing " , repoArg)
		install.Run(owner, repoName)

	//enable
	case "enable":
		fmt.Println("Enable command called")
		args := os.Args[2:]
		if len(args) < 1 {
			fmt.Println("Usage: lgd enable <owner/repo>")
			os.Exit(1)
		}
		repoArg := args[0]
		parts := strings.Split(repoArg, "/")
		if len(parts) != 2 {
			fmt.Println("Invalid repo format. Use owner/repo")
			os.Exit(1)
		}
		owner := parts[0]
		repoName := parts[1]
		enable.Run(repoName, owner)
	case "list":
		list.Run()
	case "disable":
		args := os.Args[2:]
		if len(args) < 1 {
			fmt.Println("Usage: lgd disable <owner/repo>")
			os.Exit(1)
		}
		repoArg := args [0]
		parts := strings.Split(repoArg, "/")
		if len(parts) != 2 {
			fmt.Println("Invalid repo format. Use owner/repo")
			os.Exit(1)
		}
		repoName := parts[1]
		disable.Run(repoName)
	case "remove":
		args := os.Args[2:]
                if len(args) < 1 {
                        fmt.Println("Usage: lgd disable <owner/repo>")
                        os.Exit(1)
                }
                repoArg := args [0]
                parts := strings.Split(repoArg, "/")
                if len(parts) != 2 {
                        fmt.Println("Invalid repo format. Use owner/repo")
                        os.Exit(1)
                }
                repoName := parts[1]
                remove.Run(repoName)
	case "--help":
		fmt.Printf(`lgd --setup              # prepares LGD environment
lgd install <owner/repo> # installs a repository
lgd enable <owner/repo>  # enables a project
lgd disable <owner/repo> # disabler a project
lgd remove <owner/repo>  # removes a project
lgd list                 # lists installed projects
lgd help                 # shows this help
lgd version              # shows LGD version
`)
	case "-h":        
                fmt.Printf(`lgd --setup              # prepares LGD environment
lgd install <owner/repo> # installs a repository
lgd enable <owner/repo>  # enables a project
lgd disable <owner/repo> # disabler a project
lgd remove <owner/repo>  # removes a project
lgd list                 # lists installed projects
lgd help                 # shows this help
lgd version              # shows LGD version
`)
	case "--version":
		fmt.Println("LGD Linux Git Downloader version - 1.0")

	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		os.Exit(1)
	}
}

