package checkers

// CheckFuncs : global reflection from CWE-ID to checking function, which is very significant
// You can even call external Golang functions (your own code, or from open source) in GoSan!
var CheckFuncs = map[string]interface{}{
	"128":  check128,
	"190":  check190,
	"191":  check191,
	"466":  check466,
	"478":  check478,
	"777":  check777,
	"785":  check785,
	"824":  check824,
	"1077": check1077,
	"131":  check131,
	"193":  check193,
	"369":  check369,
	"1024": check1024,
	"787":  check787,
}
