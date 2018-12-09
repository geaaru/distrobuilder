package main

import (
	"fmt"

	"github.com/lxc/distrobuilder/managers"
	"github.com/lxc/distrobuilder/shared"
)

func managePackages(def shared.DefinitionPackages, actions []shared.DefinitionAction) error {
	var err error

	manager := managers.Get(def.Manager)
	if manager == nil {
		return fmt.Errorf("Couldn't get manager")
	}

	// Handle repositories actions
	if def.Repositories != nil && len(def.Repositories) > 0 {
		if manager.RepoHandler == nil {
			return fmt.Errorf("No repository handler present!")
		}

		for _, repo := range def.Repositories {
			err = manager.RepoHandler(repo)
			if err != nil {
				return fmt.Errorf(
					fmt.Sprintf("Error for repository %s: %s", repo.Name, err))
			}
		}
	}

	err = manager.Refresh()
	if err != nil {
		return err
	}

	if def.Update {
		err = manager.Update()
		if err != nil {
			return err
		}

		// Run post update hook
		for _, action := range actions {
			err = shared.RunScript(action.Action)
			if err != nil {
				return fmt.Errorf("Failed to run post-update: %s", err)
			}
		}
	}

	err = manager.Install(def.Install)
	if err != nil {
		return err
	}

	err = manager.Remove(def.Remove)
	if err != nil {
		return err
	}

	err = manager.Clean()
	if err != nil {
		return err
	}

	return nil
}
