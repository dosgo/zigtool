# zigtool

The golang tool of the zig compiler automatically compiles different targets according to the `GOOS` and `GOARCH` environment variables. You need to install zig.

```bash
go install github.com/dosgo/zigtool/zigcc@latest
go install github.com/dosgo/zigtool/zigcpp@latest            
```

```bash
set CC=zigcc
set CXX=zigcpp
```

Manually set the compilation target

```bash
set ZIGTARGET=x86_64-windows-gnu
```
