# Fix: race condition in viper config handling

## Description
This PR fixes a race condition in the `configs` package where `viper.WatchConfig` updates the configuration concurrently with read operations.
It introduces a `sync.RWMutex` to protect access to the global `config` variable.

## Issue
Fixes #108

## Type of Change
- [x] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] Refactor (no functional changes, no api changes)
- [ ] Documentation
- [ ] Test

## Checklist
- [x] I have performed a self-review of my own code
- [x] I have added tests that prove my fix is effective or that my feature works
- [x] New and existing unit tests pass locally with my changes
- [x] I have commented my code, particularly in hard-to-understand areas

## Verification
Run the race detector test:
```bash
go test -v -race ./configs/...
```
Output:
```text
=== RUN   TestRace
--- PASS: TestRace (0.05s)
PASS
ok      github.com/xinliangnote/go-gin-api/configs      1.065s
```
