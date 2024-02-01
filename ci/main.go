package main

import (
	"context"
	"fmt"

	"dagger.io/dagger"
	"dagger.io/dagger/dag"
)

func main() {
	ctx := context.Background()
    // singleFile(ctx)
    // containerTwoFiles(ctx)
    dirTwoFiles(ctx)
}

func singleFile(ctx context.Context) {
    license := dag.Host().File("LICENSE")
	dag.Container().From("golang:1.19").
        WithFile("/", license).
        WithExec([]string{"cat", "/LICENSE"}).
        Stdout(ctx)

}

func containerTwoFiles(ctx context.Context) {
    fmt.Println(" ****************** running containerTwoFiles *******************")
    license := dag.Host().File("LICENSE")
    tsconfig := dag.Host().File("tsconfig.json")
    files := []*dagger.File{license, tsconfig}
	dag.Container().From("golang:1.19").
        WithFiles("/", files).
        WithExec([]string{"ls", "/"}).
        WithExec([]string{"cat", "/LICENSE"}).
        Stdout(ctx)
}

func dirTwoFiles(ctx context.Context) {
    fmt.Println(" ****************** running dirTwoFiles *******************")
    license := dag.Host().File("LICENSE")
    tsconfig := dag.Host().File("tsconfig.json")
    files := []*dagger.File{license, tsconfig}
    dir := dag.Directory().WithFiles("myfiles", files)
	dag.Container().From("golang:1.19").
        WithDirectory("/", dir).
        WithExec([]string{"ls", "/myfiles"}).
        WithExec([]string{"cat", "/myfiles/tsconfig.json"}).
        Stdout(ctx)
}
