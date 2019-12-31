package config

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/Unknwon/goconfig"
	"github.com/wangcong15/go-sanitizer/checkers"
)

// Config : global information
type Config struct {
	// global variables
	version string
	welcome string
	goodbye string

	// argument variables
	ArgPkg      string
	ArgCWE      string
	ArgLanguage string

	// error messages
	errNeedInputPkg string
	errPkgNotFound  string
	ErrTransAST     string
	ErrWFile        string
}

// Load : configure file
func Load(c *Config) {
	// load configure parameters
	var absPath string
	if _, err := os.Stat("~/go/src/github.com/wangcong15/go-sanitizer/config"); err == nil {
		absPath = "~/go/src/github.com/wangcong15/go-sanitizer/config"
	} else {
		absPath = toAbsPath("github.com/wangcong15/go-sanitizer/config")
	}
	configFile := fmt.Sprintf("config-%s.ini", c.ArgLanguage)
	configFile = path.Join(absPath, configFile)
	cfg, err := goconfig.LoadConfigFile(configFile)
	if err != nil {
		log.Println("非常抱歉，我们目前不支持该语言")
		log.Fatal(err)
	}
	c.version, err = cfg.GetValue("", "version")
	c.welcome, err = cfg.GetValue("", "welcome")
	c.goodbye, err = cfg.GetValue("", "goodbye")
	c.errNeedInputPkg, err = cfg.GetValue("", "errNeedInputPkg")
	c.errPkgNotFound, err = cfg.GetValue("", "errPkgNotFound")
	c.ErrTransAST, err = cfg.GetValue("", "errTransAST")
	c.ErrWFile, err = cfg.GetValue("", "errWFile")

	// load checkers

}

// GetWelcome : Get welcome string
func (c Config) GetWelcome() string {
	return fmt.Sprintf("%s %s", c.welcome, c.version)
}

// GetGoodbye : Get goodbye string
func (c Config) GetGoodbye() string {
	return c.goodbye
}

// Validate : the correctness of parameters
func (c Config) Validate() {
	if c.ArgPkg == "" {
		log.Println(c.errNeedInputPkg)
		os.Exit(1)
	}
	if s, err := os.Stat(c.ArgPkg); err != nil || !s.IsDir() {
		log.Println(c.errPkgNotFound)
		os.Exit(1)
	}
}

// GetCWEs : get CWE-IDs
func (c Config) GetCWEs() []string {
	if c.ArgCWE == "" {
		r := make([]string, 0, len(checkers.CheckFuncs))
		for k := range checkers.CheckFuncs {
			r = append(r, k)
		}
		return r
	}
	return strings.Split(c.ArgCWE, ",")
}

func toAbsPath(path string) (absPath string) {
	var goListScript string
	goListScript = fmt.Sprintf("go list -e -f {{.Dir}} %s", path)
	absPath, _ = ExecShell(goListScript)
	return
}

func ExecShell(s string) (string, string) {
	cmd := exec.Command("/bin/bash", "-c", s)
	var stdout, stderr bytes.Buffer
	var outStr, errStr string

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Run()
	outStr, errStr = string(stdout.Bytes()), string(stderr.Bytes())
	outStr = strings.TrimRight(outStr, "]\n")
	errStr = strings.TrimRight(errStr, "]\n")
	return outStr, errStr
}
