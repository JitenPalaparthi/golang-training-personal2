
- To run go application

```
go run main.go
```

- go run -> It compiles, assembles and links creates a binary in a temp location and that binary is executed

- Go build

```
go build hello.go
```

- Go build to stripe down the binary

```
go build -ldflags="-w -s" -o demo1 hello.go
```

- Go environment variables

    - GOARCH : Processor Architecture (32bit/ 64bit arm or amd etc)
    - GOOS : Operating system (windows, darwin, mips, linux, ios, android)

- To check supported os/arch

```
go tool dist list
```
- output 
```
aix/ppc64
android/386
android/amd64
android/arm
android/arm64
darwin/amd64
darwin/arm64
dragonfly/amd64
freebsd/386
freebsd/amd64
freebsd/arm
freebsd/arm64
freebsd/riscv64
illumos/amd64
ios/amd64
ios/arm64
js/wasm
linux/386
linux/amd64
linux/arm
linux/arm64
linux/loong64
linux/mips
linux/mips64
linux/mips64le
linux/mipsle
linux/ppc64
linux/ppc64le
linux/riscv64
linux/s390x
netbsd/386
netbsd/amd64
netbsd/arm
netbsd/arm64
openbsd/386
openbsd/amd64
openbsd/arm
openbsd/arm64
openbsd/ppc64
plan9/386
plan9/amd64
plan9/arm
solaris/amd64
wasip1/wasm
windows/386
windows/amd64
windows/arm
windows/arm64
```
- To cross compile build 

```
GOOS=windows GOARCH=amd64 go build -o hello.exe hello.go
```

- Go build with details 

```
go build -x -v
```

# for Link issue
```
sudo GODEBUG=installgoroot=all go install std
```
