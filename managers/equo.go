package managers

import (
	"fmt"

	"github.com/lxc/distrobuilder/shared"
)

func enman_repo_caller(repo shared.DefinitionRepository) error {
	var err error
	var args []string

	if repo.Action == "" {
		err = fmt.Errorf("Invalid action!")
	} else {
		args = []string{
			repo.Action,
		}

		if repo.Action == "add" {
			if repo.Url != "" {
				args = append(args, repo.Url)
			} else {
				args = append(args, repo.Name)
			}
		} else if repo.Action == "remove" ||
			repo.Action == "enable" || repo.Action == "enable" {
			args = append(args, repo.Name)
		} else {
			err = fmt.Errorf("Invalid Action")
		}
	}

	if err == nil {
		err = shared.RunCommand("enman", args...)
	}

	return err
}

func equo_repo_caller(repo shared.DefinitionRepository) error {
	var args []string = []string{
		"repo",
	}

	if repo.Name == "" {
		return fmt.Errorf("Invalid repository name!")
	}

	if repo.Url == "" {
		return fmt.Errorf("Invalid url!")
	}

	if repo.Action == "" {
		return fmt.Errorf("Invalid action!")
	}

	args = append(args, repo.Action)

	if repo.Action == "add" {
		args = append(args, "--repo")
		args = append(args, repo.Url)
		args = append(args, "--pkg")
		args = append(args, repo.Url)
		args = append(args, repo.Name)

	} else if repo.Action == "remove" || repo.Action == "enable" ||
		repo.Action == "enable" {
		args = append(args, repo.Name)
	} else {
		return fmt.Errorf("Invalid Action!")
	}

	return shared.RunCommand("equo", args...)
}

// NewEquo creates a new Manager instance
func NewEquo() *Manager {
	return &Manager{
		command: "equo",
		flags: ManagerFlags{
			global: []string{},
			clean: []string{
				"cleanup",
			},
			install: []string{
				"install",
			},
			remove: []string{
				"remove",
			},
			refresh: []string{
				"update",
			},
			update: []string{
				"upgrade",
			},
		},
		RepoHandler: func(repoAction shared.DefinitionRepository) error {
			if repoAction.Type == "" || repoAction.Type == "equo" {
				return equo_repo_caller(repoAction)
			} else if repoAction.Type == "enman" {
				return enman_repo_caller(repoAction)
			}

			return fmt.Errorf("Invalid repository Type")
		},
	}
}
