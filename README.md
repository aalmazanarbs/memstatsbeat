# Memstatsbeat

Welcome to Memstatsbeat.

Ensure that this folder is at the following location:
`${GOPATH}/github.com/aalmazanarbs/memstatsbeat`

## Getting Started with Memstatsbeat

### Requirements

* [Golang](https://golang.org/dl/) 1.8

### Dependencies

The project need to install the following dependencies to run.

```
go get golang.org/x/sys/unix
go get github.com/shirou/gopsutil
```

### Init Project
To get running with Memstatsbeat run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


### Build

To build the binary for Memstatsbeat run the command below. This will generate a binary
in the same directory with the name memstatsbeat.

```
make
```


### Run

To run Memstatsbeat with debugging output enabled, run:

```
./memstatsbeat -c memstatsbeat.yml -e -d "*"
```


### Test

To test Memstatsbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `etc/fields.yml`.
To generate etc/memstatsbeat.template.json and etc/memstatsbeat.asciidoc

```
make update
```


### Cleanup

To clean  Memstatsbeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Memstatsbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/aalmazanarbs/memstatsbeat
cd ${GOPATH}/github.com/aalmazanarbs/memstatsbeat
git clone https://github.com/aalmazanarbs/memstatsbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make package
```

This will fetch and create all images required for the build process. The hole process to finish can take several minutes.
