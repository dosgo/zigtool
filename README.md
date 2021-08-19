# zigtool
The golang tool of the zig compiler automatically compiles different targets according to the GOOS GOARCH environment variable. You need to install zig.


go install github.com/dosgo/zigtool/zigcc@latest

go install github.com/dosgo/zigtool/zigc++@latest


set CC=zigcc
set CXX=zigc++

Manually set the compilation target

set ZIGTARGET=x86_64-windows-gnu
