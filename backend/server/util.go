package server

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func OpenBrowser(url string) {
	var err error

	fmt.Println()

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		fmt.Println(err)
	}

	s := "If your web browser doesn't open automatically, navigate to " + url + " manually"
	t := strings.Repeat("=", len(s))
	fmt.Println()
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t)

}
