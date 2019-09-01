package language

import (
	"github.com/coderbyte-org/cb-code-runner/language/javascript"
	"github.com/coderbyte-org/cb-code-runner/language/python"
	"github.com/coderbyte-org/cb-code-runner/language/ruby"
	"github.com/coderbyte-org/cb-code-runner/language/php"
	"github.com/coderbyte-org/cb-code-runner/language/java"
	"github.com/coderbyte-org/cb-code-runner/language/swift"
	"github.com/coderbyte-org/cb-code-runner/language/golang"
	"github.com/coderbyte-org/cb-code-runner/language/cpp"
	"github.com/coderbyte-org/cb-code-runner/language/csharp"
	"github.com/coderbyte-org/cb-code-runner/language/c"
	"github.com/coderbyte-org/cb-code-runner/language/kotlin"
)

type runFn func([]string, string) (string, string, error)

var languages = map[string]runFn{
	"javascript":   javascript.Run,
	"python":       python.Run,
	"ruby":         ruby.Run,
	"php":          php.Run,
	"java":         java.Run,
	"swift":        swift.Run,
	"go":           golang.Run,
	"cpp":          cpp.Run,
	"csharp":       csharp.Run,
	"c":            c.Run,
	"kotlin":       kotlin.Run,
}

func IsSupported(lang string) bool {
	_, supported := languages[lang]
	return supported
}

func Run(lang string, files []string, stdin string) (string, string, error) {
	return languages[lang](files, stdin)
}
