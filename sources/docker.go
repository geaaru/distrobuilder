package sources

import (
	"path/filepath"

	"github.com/lxc/distrobuilder/shared"
	dcapi "github.com/mudler/docker-companion/api"
)

type DockerHTTP struct{}

// NewDockerHTTP create a new DockerHTTP instance for
// use a docker image as rootfs
func NewDockerHTTP() *DockerHTTP {
	return &DockerHTTP{}
}

// Run downloads and unpack a docker image
func (d *DockerHTTP) Run(definition shared.Definition, rootfsDir string) error {
	var absRootfsDir string
	var err error

	absRootfsDir, err = filepath.Abs(rootfsDir)
	if err == nil {
		// NOTE: For now we use only docker official server but we can
		//       add a new parameter on DefinitionSource struct.
		err = dcapi.DownloadImage(definition.Source.URL, absRootfsDir, "")
	}

	return err
}
