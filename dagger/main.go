package main

import (
	"context"
)


type Wontexit struct{}

// Application specific build logic
func (m *Wontexit) Build(ctx context.Context, buildContext *Directory) *Container {
	return dag.Container().
		From("bash").
		WithFile("/wontexit.sh", buildContext.File("wontexit.sh")).
		WithEntrypoint([]string{"/wontexit.sh"})
}

// Take the built container and push it
func (m *Wontexit) BuildAndPush(ctx context.Context, registry, imageName, username string, password *Secret, buildContext *Directory) error {
	_, err := m.Build(ctx, buildContext).
		WithRegistryAuth(registry, username, password).
		Publish(ctx, registry+"/"+imageName)
	return err
}
