# Jobsearch CLI

Jobsearch is a command line jobsearch tool.

## Usage

    jobsearch_cli -title 'Web Developer' -location London

    jobsearch_cli -titles 'Web Developer','Software Engineer' -location London

The default output is JSON, but you can use the `-csv` to output comma-separated values instead.
You can use `-csv-headers` to include the heads with the csv.

### Intended usage

The intended usage of this program is to save the output to a file.

    jobsearch_cli -title 'Web Developer' -location London > jobsearch.json

## Make

You can use `make` on windows or linux to run and build this project.

- See full list of make commands by running `make list` from the project root.

## Compile

Compile will build binaries for the declared platforms in the /bin directly.
If you want to build for a specific platform such as windows-arm-64 or linux-arm-64  
You can use the following:

```BASH
# Build for linux arm64
make compile.linux64

# Build for windows arm64
make compile.windows64

# Or to build for all supported platforms
make compile
```

## Run 

Run will run the application without building an executable.  
You can pass flags using the ARGS env.

```BASH
# Run app and get version
make run ARGS="-version"
```

## Test

Test will run all tests within the source.
You can also pass flags using the ARGS env. You may want to use this to pass flags to `go test`.  

```BASH
# Run tests and show logs in stdout
make test ARGS="-v"
```