# LGD – Linux Git Downloader

LGD (Linux Git Downloader) is a simple CLI tool for Linux that helps you download,
organize and manage GitHub repositories locally.

The main goal of this project is to provide a clean and repeatable way to work
with multiple Git projects using a single command-line interface.

This is **v1**, focused on core functionality.

---

## Features

- Download GitHub repositories using a single command
- Store all projects in one predictable location
- Enable and manage projects with metadata
- Simple CLI interface
- Written in Go

---

## Requirements

- Linux
- Git
- Go (to build LGD)
- Python 3 (optional, for Python projects)

---

## Installation

Clone the repository:

```bash
git clone https://github.com/ni3wyspany/lgd.git
cd lgd
```

---

## Build the binary:

```bash
go build -o lgd cmd/lgd/main.go
```

---

## Run LGD:

```bash
./lgd --help
```

---

## First-time setup

Before using LGD, initialize its directory structure:

```bash
./lgd --setup
```


This will create the required folders and configuration files in your home
directory.

---

## Usage
### Install a repository
```bash
./lgd install owner/repo
```

---

### Example:
```bash
./lgd install laramies/theHarvester
```

---

### Enable a project
```bash
./lgd enable owner/repoName
```

This step prepares the project for use and creates required metadata.

---

### List installed projects
```bash
./lgd list
```

---

### Remove a project
```bash
./lgd remove owner/repoName
```

---

## Help and version
```bash
./lgd --help
./lgd --version
```

---

## Directory structure

LGD uses the following structure inside the user's home directory:

```text
~/lgd/
├── bin/
├── projects/     # all downloaded repositories
├── cache/
├── meta/
├── cmd/
└── config.toml
```


All repositories are stored in a single projects directory to keep the
structure simple and predictable.

---

## Project status

LGD v1 provides the core functionality.
Future versions may include:

- More project metadata

- Global binary installation

- Configuration improvements

---

## License
MIT License
