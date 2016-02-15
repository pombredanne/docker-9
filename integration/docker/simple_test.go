package dockerit

import (
	"strings"
	"testing"

	"github.com/vdemeester/libkermit/docker"
)

func TestCreateSimple(t *testing.T) {
	setupTest(t)

	container, err := docker.Create("busybox")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if container.ID == "" {
		t.Fatalf("expected a containerId, got nothing")
	}
	if strings.HasPrefix(container.Name, "kermit_") {
		t.Fatalf("expected name to start with 'kermit_', got %s", container.Name)
	}
}

func TestStartAndStop(t *testing.T) {
	setupTest(t)

	container, err := docker.Start("busybox")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if container.ID == "" {
		t.Fatalf("expected a containerId, got nothing")
	}
	if strings.HasPrefix(container.Name, "kermit_") {
		t.Fatalf("expected name to start with 'kermit_', got %s", container.Name)
	}
	if !container.State.Running {
		t.Fatalf("expected container to be running, but was in state %v", container.State)
	}

	err = docker.Stop(container.ID)
	if err != nil {
		t.Fatal(err)
	}

	container, err = docker.Inspect(container.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if container.State.Running {
		t.Fatalf("expected container to be running, but was in state %v", container.State)
	}
}
