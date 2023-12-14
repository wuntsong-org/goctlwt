package api

import (
	"github.com/spf13/cobra"
	"github.com/wuntsong-org/goctlwt/api/apigen"
	"github.com/wuntsong-org/goctlwt/api/dartgen"
	"github.com/wuntsong-org/goctlwt/api/docgen"
	"github.com/wuntsong-org/goctlwt/api/format"
	"github.com/wuntsong-org/goctlwt/api/gogen"
	"github.com/wuntsong-org/goctlwt/api/javagen"
	"github.com/wuntsong-org/goctlwt/api/ktgen"
	"github.com/wuntsong-org/goctlwt/api/new"
	"github.com/wuntsong-org/goctlwt/api/tsgen"
	"github.com/wuntsong-org/goctlwt/api/validate"
	"github.com/wuntsong-org/goctlwt/config"
	"github.com/wuntsong-org/goctlwt/internal/cobrax"
	"github.com/wuntsong-org/goctlwt/plugin"
)

var (
	// Cmd describes an api command.
	Cmd       = cobrax.NewCommand("api", cobrax.WithRunE(apigen.CreateApiTemplate))
	dartCmd   = cobrax.NewCommand("dart", cobrax.WithRunE(dartgen.DartCommand))
	docCmd    = cobrax.NewCommand("doc", cobrax.WithRunE(docgen.DocCommand))
	formatCmd = cobrax.NewCommand("format", cobrax.WithRunE(format.GoFormatApi))
	goCmd     = cobrax.NewCommand("go", cobrax.WithRunE(gogen.GoCommand))
	newCmd    = cobrax.NewCommand("new", cobrax.WithRunE(new.CreateServiceCommand),
		cobrax.WithArgs(cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs)))
	validateCmd = cobrax.NewCommand("validate", cobrax.WithRunE(validate.GoValidateApi))
	javaCmd     = cobrax.NewCommand("java", cobrax.WithRunE(javagen.JavaCommand), cobrax.WithHidden())
	ktCmd       = cobrax.NewCommand("kt", cobrax.WithRunE(ktgen.KtCommand))
	pluginCmd   = cobrax.NewCommand("plugin", cobrax.WithRunE(plugin.PluginCommand))
	tsCmd       = cobrax.NewCommand("ts", cobrax.WithRunE(tsgen.TsCommand))
)

func init() {
	var (
		apiCmdFlags      = Cmd.Flags()
		dartCmdFlags     = dartCmd.Flags()
		docCmdFlags      = docCmd.Flags()
		formatCmdFlags   = formatCmd.Flags()
		goCmdFlags       = goCmd.Flags()
		javaCmdFlags     = javaCmd.Flags()
		ktCmdFlags       = ktCmd.Flags()
		newCmdFlags      = newCmd.Flags()
		pluginCmdFlags   = pluginCmd.Flags()
		tsCmdFlags       = tsCmd.Flags()
		validateCmdFlags = validateCmd.Flags()
	)

	apiCmdFlags.StringVar(&apigen.VarStringOutput, "o")
	apiCmdFlags.StringVar(&apigen.VarStringHome, "home")
	apiCmdFlags.StringVar(&apigen.VarStringRemote, "remote")
	apiCmdFlags.StringVar(&apigen.VarStringBranch, "branch")

	dartCmdFlags.StringVar(&dartgen.VarStringDir, "dir")
	dartCmdFlags.StringVar(&dartgen.VarStringAPI, "api")
	dartCmdFlags.BoolVar(&dartgen.VarStringLegacy, "legacy")
	dartCmdFlags.StringVar(&dartgen.VarStringHostname, "hostname")
	dartCmdFlags.StringVar(&dartgen.VarStringScheme, "scheme")

	docCmdFlags.StringVar(&docgen.VarStringDir, "dir")
	docCmdFlags.StringVar(&docgen.VarStringOutput, "o")

	formatCmdFlags.StringVar(&format.VarStringDir, "dir")
	formatCmdFlags.BoolVar(&format.VarBoolIgnore, "iu")
	formatCmdFlags.BoolVar(&format.VarBoolUseStdin, "stdin")
	formatCmdFlags.BoolVar(&format.VarBoolSkipCheckDeclare, "declare")

	goCmdFlags.StringVar(&gogen.VarStringDir, "dir")
	goCmdFlags.StringVar(&gogen.VarStringAPI, "api")
	goCmdFlags.StringVar(&gogen.VarStringHome, "home")
	goCmdFlags.StringVar(&gogen.VarStringRemote, "remote")
	goCmdFlags.StringVar(&gogen.VarStringBranch, "branch")
	goCmdFlags.StringVarWithDefaultValue(&gogen.VarStringStyle, "style", config.DefaultFormat)

	javaCmdFlags.StringVar(&javagen.VarStringDir, "dir")
	javaCmdFlags.StringVar(&javagen.VarStringAPI, "api")

	ktCmdFlags.StringVar(&ktgen.VarStringDir, "dir")
	ktCmdFlags.StringVar(&ktgen.VarStringAPI, "api")
	ktCmdFlags.StringVar(&ktgen.VarStringPKG, "pkg")

	newCmdFlags.StringVar(&new.VarStringHome, "home")
	newCmdFlags.StringVar(&new.VarStringRemote, "remote")
	newCmdFlags.StringVar(&new.VarStringBranch, "branch")
	newCmdFlags.StringVarWithDefaultValue(&new.VarStringStyle, "style", config.DefaultFormat)

	pluginCmdFlags.StringVarP(&plugin.VarStringPlugin, "plugin", "p")
	pluginCmdFlags.StringVar(&plugin.VarStringDir, "dir")
	pluginCmdFlags.StringVar(&plugin.VarStringAPI, "api")
	pluginCmdFlags.StringVar(&plugin.VarStringStyle, "style")

	tsCmdFlags.StringVar(&tsgen.VarStringDir, "dir")
	tsCmdFlags.StringVar(&tsgen.VarStringAPI, "api")
	tsCmdFlags.StringVar(&tsgen.VarStringCaller, "caller")
	tsCmdFlags.BoolVar(&tsgen.VarBoolUnWrap, "unwrap")

	validateCmdFlags.StringVar(&validate.VarStringAPI, "api")

	// Add sub-commands
	Cmd.AddCommand(dartCmd, docCmd, formatCmd, goCmd, javaCmd, ktCmd, newCmd, pluginCmd, tsCmd, validateCmd)
}
