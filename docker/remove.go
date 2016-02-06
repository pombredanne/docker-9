package docker

import (
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
)

// Remove removes the container
func Remove(containerID string) error {
	client, err := client.NewEnvClient()
	if err != nil {
		return err
	}
	return remove(client, containerID)
}

func remove(cli client.APIClient, containerID string) error {
	if err := cli.ContainerRemove(types.ContainerRemoveOptions{
		ContainerID: containerID,
		Force:       true,
	}); err != nil {
		return err
	}
	return nil
}