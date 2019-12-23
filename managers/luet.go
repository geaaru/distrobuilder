package managers

import (
	"github.com/lxc/distrobuilder/shared"
)

func luetRepoCaller(repo shared.DefinitionPackagesRepository) error {
	// TODO
	return nil
}

// NewLuet create a new Manager instance
func NewLuet() *Manager {
	return &Manager{
		commands: ManagerCommands{
			install: "luet",
			refresh: "luet",
			remove:  "luet",
			update:  "luet",
		},
		flags: ManagerFlags{
			global: []string{},
			install: []string{
				"install",
			},
			remove: []string{
				"uninstall",
			},
			update: []string{
				"upgrade",
			},
		},
		RepoHandler: func(repoAction shared.DefinitionPackagesRepository) error {
			return luetRepoCaller(repoAction)
		},
	}
}
