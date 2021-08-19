package comm

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Build(buildName string) {
	var envs = make(map[string]string)
	envs["GOOS"] = os.Getenv("GOOS")
	envs["GOARCH"] = os.Getenv("GOARCH")
	if envs["GOOS"] == "" {
		envs["GOOS"] = runtime.GOOS
	}
	if envs["GOARCH"] == "" {
		envs["GOARCH"] = runtime.GOARCH
	}
	var target = ""
	var goos = envs["GOOS"]
	var goarch = envs["GOARCH"]
	if goos == "windows" {
		switch goarch {
		case "386":
			target = "i386-windows-gnu"
			break
		case "amd64":
			target = "x86_64-windows-gnu"
			break
		case "arm":
			target = "arm-windows-gnu"
			break
		case "arm64":
			target = "aarch64-windows-gnu"
			break
		}
	}
	if goos == "linux" {

		switch goarch {
		case "386":
			target = "i386-linux-gnu"
			break
		case "amd64":
			target = "x86_64-linux-gnu"
			break
		case "arm":
			target = "arm-linux-gnueabi"
			break
		case "arm64":
			target = "aarch64-linux-gnu"
			break
		case "mips":
			target = "mips-linux-gnu"
			break
		case "mips64":
			target = "mips64-linux-musl"
			break
		case "mips64le":
			target = "mips64el-linux-musl"
			break
		case "mipsle":
			target = "mipsel-linux-gnu"
			break
		}
	}
	//no cross build
	if envs["GOARCH"] == runtime.GOARCH && envs["GOOS"] == runtime.GOOS {
		target = ""
	}

	var zigTarget = os.Getenv("ZIGTARGET")
	if len(zigTarget) > 0 {
		target = zigTarget
	}
	var args = []string{buildName}
	if len(target) > 0 {
		args = append(args, "-target")
		args = append(args, target)
	}
	for index, value := range os.Args {
		if index > 0 {
			args = append(args, value)
		}
	}
	cmd := exec.Command("zig", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("zig cmd:%s\r\n", strings.Join(cmd.Args, " "))
	cmd.Start()
	cmd.Wait()
}
