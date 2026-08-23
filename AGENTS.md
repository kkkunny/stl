# AGENTS.md

Go generics-based standard library extension (stl = "standard template library"). Module `github.com/kkkunny/stl`, requires Go 1.23+. No CI, no lint config in repo.

## Commands

- `go test ./...` — full suite; tests live next to their code
- `go build ./...`, `go vet ./...` — verification

## Package naming gotcha

Most packages are named with a `stl` prefix that differs from the directory name — import by the prefixed name, not the dir:

- `stlcmp` (cmp), `stlos` (os), `stlhash` (hash), `stlslices` (container/slices), `stlheap` (container/heap), `stliter` (container/iter), `stlmaps` (container/maps), `stlerr` (error), `stllog` (log), `stlmath` (math), `stlbits` (math/bits), `stlreflect` (reflect), `stlruntime` (runtime), `stlstr` (str), `stlsync` (sync), `stltest` (test), `stltype` (types), `stlunsafe` (unsafe), `stlval` (value), `stlbasic` (basic)
- BUT container sub-packages keep plain names: `bimap`, `either`, `hashmap`, `linkedhashmap`, `linkedlist`, `optional`, `pqueue`, `queue`, `set`, `stack`, `treemap`, `tuple`, plus `clone`, `enum`, `lazy`

## Conventions

- Code comments are written in Chinese — follow that in new code.
- Tests use the custom `stltest` package (`AssertEq`, `AssertNotEq`), not testify. Some files also contain benchmarks.
- `generic/list` and `generic/pipe` are forks of Go stdlib code — keep the Go Authors copyright headers and minimize divergence.
- `internal/json` and `internal/reflect` are internal helpers used by container packages (bimap, hashmap, linkedhashmap) and clone/cmp/reflect.
- Arch-specific code is split via build tags (e.g. `os/size_386.go`, `os/size_amd64.go`).
- `example/` is a scratch main package, not part of the library.
- Default branch is `master`; iterator work happens on `feat/iter`.