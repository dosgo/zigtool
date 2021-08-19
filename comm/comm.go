package comm

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Build(buildName string) {
	envs := getGoEnvs()
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
	if envs["GOARCH"] == envs["GOHOSTARCH"] && envs["GOOS"] == envs["GOHOSTOS"] {
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

func getGoEnvs() map[string]string {
	var envs = make(map[string]string, 0)
	cmd := exec.Command("go", "env")
	out, _ := cmd.CombinedOutput()
	lines := strings.Split(string(out), "\n")
	for _, v := range lines {
		kv := strings.Split(v, "=")
		if len(kv) > 1 {
			envs[strings.Trim(strings.Replace(kv[0], "set", "", -1), " ")] = kv[1]
		}
	}
	return envs
}
