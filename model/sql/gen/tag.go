package gen

import (
	"github.com/wuntsong-org/goctlwt/model/sql/template"
	"github.com/wuntsong-org/goctlwt/util"
	"github.com/wuntsong-org/goctlwt/util/pathx"
)

func genTag(table Table, in string) (string, error) {
	if in == "" {
		return in, nil
	}

	text, err := pathx.LoadTemplate(category, tagTemplateFile, template.Tag)
	if err != nil {
		return "", err
	}

	output, err := util.With("tag").Parse(text).Execute(map[string]any{
		"field": in,
		"data":  table,
	})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
