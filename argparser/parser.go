package argparser

import (
	"flag"

	"github.com/wangcong15/go-sanitizer/config"
)

// Parse : the terminal parameters
func Parse(c *config.Config) {
	// parse user input parameters
	flag.StringVar(&c.ArgPkg, "p", "", "指定输入包以推荐断言 / Specify the checking package")
	flag.StringVar(&c.ArgCWE, "c", "", "指定要检查的CWE-ID，以英文逗号分隔 / Specify CWE-IDs seperated by comma")
	flag.StringVar(&c.ArgLanguage, "lang", "cn", "目前支持中文和英文 / language currently support cn and en")
	flag.Parse()
}
