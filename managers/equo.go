package managers

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
	}
}
