# Welcome to the contributing guide <!-- omit in toc -->

Thank you for investing your time in contributing to our project!

## New contributor guide

To get an overview of the project, read the [README](README.md).

## Linting

Linting commands to run from the devcontainer/pipeline.

Pylint

```bash
golangci-lint run
```

## Testing

### Unit testing

```bash
go test -v ./header_setter -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Integration testing

```bash
npx sls invoke -f setter -s dev -p fixtures/s3_event.json
```