package generator

import (
	_ "embed"
	"path/filepath"
	"strings"

	"github.com/wuntsong-org/goctlwt/util"
	"github.com/wuntsong-org/goctlwt/util/pathx"
	"github.com/wuntsong-org/goctlwt/util/stringx"
)

//go:embed rpc.tpl
var rpcTemplateText string

// ProtoTmpl returns a sample of a proto file
func ProtoTmpl(out string) error {
	protoFilename := filepath.Base(out)
	serviceName := stringx.From(strings.TrimSuffix(protoFilename, filepath.Ext(protoFilename)))
	text, err := pathx.LoadTemplate(category, rpcTemplateFile, rpcTemplateText)
	if err != nil {
		return err
	}

	dir := filepath.Dir(out)
	err = pathx.MkdirIfNotExist(dir)
	if err != nil {
		return err
	}

	err = util.With("t").Parse(text).SaveTo(map[string]string{
		"package":     serviceName.Untitle(),
		"serviceName": serviceName.Title(),
	}, out, false)
	return err
}
