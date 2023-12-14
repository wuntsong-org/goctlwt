package migrationnotes

import (
	"github.com/wuntsong-org/goctlwt/config"
	"github.com/wuntsong-org/goctlwt/util/format"
)

// BeforeCommands run before comamnd run to show some migration notes
func BeforeCommands(dir, style string) error {
	if err := migrateBefore1_3_4(dir, style); err != nil {
		return err
	}
	return nil
}

func getModelSuffix(style string) (string, error) {
	cfg, err := config.NewConfig(style)
	if err != nil {
		return "", err
	}
	baseSuffix, err := format.FileNamingFormat(cfg.NamingFormat, "_model")
	if err != nil {
		return "", err
	}
	return baseSuffix + ".go", nil
}
