# yml 文件一定要注意缩进，每一级缩进两个空格，不可多，不可少。就很皮

# Axtx_beats

Welcome to Axtx_beats.

Ensure that this folder is at the following location:
`${GOPATH}/github.com/caibaoying/axtx_beats`

## Getting Started with Axtx_beats


### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Init Project
To get running with Axtx_beats and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Axtx_beats in the git repository, run the following commands:

```
git remote set-url origin https://github.com/caibaoying/axtx_beats
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Axtx_beats run the command below. This will generate a binary
in the same directory with the name axtx_beats.

```
make
```


### Run

To run Axtx_beats with debugging output enabled, run:

```
./axtx_beats -c axtx_beats.yml -e -d "*"
```


### Test

To test Axtx_beats, run the following command:

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
To generate etc/axtx_beats.template.json and etc/axtx_beats.asciidoc

```
make update
```


### Cleanup

To clean  Axtx_beats source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Axtx_beats from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/caibaoying/axtx_beats
cd ${GOPATH}/github.com/caibaoying/axtx_beats
git clone https://github.com/caibaoying/axtx_beats
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make package
```

This will fetch and create all images required for the build process. The hole process to finish can take several minutes.
