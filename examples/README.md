Examples (and validation) of usage for protobuf and steamlang packages.

They are written as tests so they are run with `go test ./...` automatically and are excluded from
package builds.

They are kept separate from the original packages because they also validate that imports work
correctly (since importing protobuf packages with duplicate messages fails, for example).
