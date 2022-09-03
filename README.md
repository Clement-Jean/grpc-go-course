# gRPC Go

[![Go Unit Tests](https://github.com/Clement-Jean/grpc-go-course/actions/workflows/tests.yml/badge.svg)](https://github.com/Clement-Jean/grpc-go-course/actions/workflows/tests.yml) [![Lint protobuf](https://github.com/Clement-Jean/grpc-go-course/actions/workflows/lint.yml/badge.svg)](https://github.com/Clement-Jean/grpc-go-course/actions/workflows/lint.yml) ![cross-platform](https://img.shields.io/badge/Platform-windows%20%7C%20macos%20%7C%20linux-brightgreen) ![Udemy](.github/badges/udemy.svg)

## COUPON: `START_SEP_22`

## Notes

### `Windows`

- I recommend you use powershell (try to update: [see](https://github.com/PowerShell/PowerShell/releases)) for following this course, you might have unexepected behavior if you use Git bash or other (especially with OpenSSL)
- I recommend you use [Chocolatey](https://chocolatey.org/) as package installer (see [Install](https://chocolatey.org/install))


### Build

#### `Linux/MacOS`

```shell
make all
```
***all is a Makefile rule** - check the other [rules](#makefile)

#### `Windows - Chocolatey`
```shell
choco install make
make all
```
***all is a Makefile rule** - check the other [rules](#makefile)

#### `Windows - Without Chocolatey`

```shell
protoc -Igreet/proto --go_opt=module=github.com/Clement-Jean/grpc-go-course --go_out=. --go-grpc_opt=module=github.com/Clement-Jean/grpc-go-course --go-grpc_out=. greet/proto/*.proto

protoc -Icalculator/proto --go_opt=module=github.com/Clement-Jean/grpc-go-course --go_out=. --go-grpc_opt=module=github.com/Clement-Jean/grpc-go-course --go-grpc_out=. calculator/proto/*.proto

protoc -Iblog/proto --go_opt=module=github.com/Clement-Jean/grpc-go-course --go_out=. --go-grpc_opt=module=github.com/Clement-Jean/grpc-go-course --go-grpc_out=. blog/proto/*.proto

go build -o bin/greet/server.exe ./greet/server
go build -o bin/greet/client.exe ./greet/client

go build -o bin/calculator/server.exe ./calculator/server
go build -o bin/calculator/client.exe ./calculator/client

go build -o bin/blog/server.exe ./blog/server
go build -o bin/blog/client.exe ./blog/client
```

<a name="makefile"></a>
## Makefile

For more information about what are the rules defined in the Makefile, please type:

```shell
make help
```

## Reporting a bug

As I need to know a little bit more information about your environment to help you, when filling an issue, please provide the output of:

```shell
make about
```
