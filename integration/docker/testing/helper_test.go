package dockerit

import (
	"testing"

	"golang.org/x/net/context"

	dockerclient "github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	d "github.com/vdemeester/libkermit/docker"
	docker "github.com/vdemeester/libkermit/docker/testing"
)

func setupTest(t *testing.T) *docker.Project {
	return cleanContainers(t)
}

func cleanContainers(t *testing.T) *docker.Project {
	client, err := dockerclient.NewEnvClient()
	if err != nil {
		t.Fatal(err)
	}

	filterArgs := filters.NewArgs()
	if filterArgs, err = filters.ParseFlag(d.KermitLabelFilter, filterArgs); err != nil {
		t.Fatal(err)
	}

	containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{
		All:    true,
		Filter: filterArgs,
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, container := range containers {
		t.Logf("cleaning container %s…", container.ID)
		if err := client.ContainerRemove(context.Background(), types.ContainerRemoveOptions{
			ContainerID: container.ID,
			Force:       true,
		}); err != nil {
			t.Errorf("Error while removing container %s : %v\n", container.ID, err)
		}
	}

	return docker.NewProject(client)
}