package comm

import (
	"os"
	"os/exec"
	"runtime"
)

func Build(buildName string){
	var target="";
	var goos=os.Getenv("GOOS")
	var goarch=os.Getenv("GOARCH")
	if goos=="windows"{
		switch goarch {
		case "386":
			target="i386-windows-gnu";
			break
		case "amd64":
			target="x86_64-windows-gnu";
			break
		case "arm":
			target="arm-windows-gnu";
			break
		case "arm64":
			target="aarch64-windows-gnu";
			break
		}
	}
	if runtime.GOOS=="linux"{

		switch goarch {
		case "386":
			target="i386-linux-gnu";
			break
		case "amd64":
			target="x86_64-linux-gnu";
			break
		case "arm":
			target="arm-linux-gnueabi";
			break
		case "arm64":
			target="aarch64-linux-gnu";
			break
		case "mips":
			target="mips-linux-gnu";
			break
		case "mips64":
			target="mips64-linux-musl";
			break
		case "mips64le":
			target="mips64el-linux-musl";
			break
		case "mipsle":
			target="mipsel-linux-gnu";
			break
		}
	}
	var args =[] string {buildName}
	if len(target)>0{
		args=append(args,"-target");
		args=append(args,target);
	}
	for index,value:=range os.Args{
		if index>0 {
			args=append(args,value);
		}
	}
	cmd:=exec.Command("zig",args...)
	cmd.Stdout=os.Stdout;
	cmd.Stderr=os.Stderr;
	cmd.Start();
	cmd.Wait();
}