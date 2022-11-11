//go:build mage

package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

// file test.txt is removed
func WorksWithFile(ctx context.Context) {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	workdir := client.Host().Workdir()

	dir := client.Directory().WithDirectory(".", workdir, dagger.DirectoryWithDirectoryOpts{
		Exclude: []string{"test.txt"},
	})

	contents, _ := client.Container().
		From("alpine:latest").
		WithMountedDirectory("/host", dir).
		Exec(dagger.ContainerExecOpts{
			Args: []string{"ls", "/host", "-a"},
		}).Stdout().Contents(ctx)

	fmt.Println(contents)
}

// directory build is removed
func WorksWithWorkdir(ctx context.Context) {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	dir := client.Host().Workdir(dagger.HostWorkdirOpts{
		Exclude: []string{"build"},
	})

	contents, _ := client.Container().
		From("alpine:latest").
		WithMountedDirectory("/host", dir).
		Exec(dagger.ContainerExecOpts{
			Args: []string{"ls", "/host", "-a"},
		}).Stdout().Contents(ctx)

	fmt.Println(contents)
}

// directory build exists but is empty
func WorksWithStar(ctx context.Context) {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	workdir := client.Host().Workdir()

	dir := client.Directory().WithDirectory(".", workdir, dagger.DirectoryWithDirectoryOpts{
		Exclude: []string{"build/*"},
	})

	contents, _ := client.Container().
		From("alpine:latest").
		WithMountedDirectory("/host", dir).
		Exec(dagger.ContainerExecOpts{
			Args: []string{"ls", "/host/build", "-a"},
		}).Stdout().Contents(ctx)

	fmt.Println(contents)
}

// fails
func Fails(ctx context.Context) {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	workdir := client.Host().Workdir()

	dir := client.Directory().WithDirectory(".", workdir, dagger.DirectoryWithDirectoryOpts{
		Exclude: []string{"build"},
	})

	contents, _ := client.Container().
		From("alpine:latest").
		WithMountedDirectory("/host", dir).
		Exec(dagger.ContainerExecOpts{
			Args: []string{"ls", "/host", "-a"},
		}).Stdout().Contents(ctx)

	fmt.Println(contents)
}
