package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zuiwuchang/ng-xi18n/cmd/cmdupdate"
	"log"
	"strings"
)

func init() {
	context := &cmdupdate.Context{}
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update message file",
		Long: `new message file
	ng-xi18n update -s src/messages.xlf -d src/locale/zh-Hant.xlf 
	ng-xi18n update -s src/messages.xlf -d src/locale/zh-Hant.xlf
	ng-xi18n update -s src/messages.xlf -l zh-Hant
`,
		Run: func(cmd *cobra.Command, args []string) {
			context.Dist = strings.TrimSpace(context.Dist)
			context.Locale = strings.TrimSpace(context.Locale)
			if context.Dist == "" && context.Locale == "" {
				log.Fatalln("unknow locale")
			}
			if context.Dist == "" {
				context.Dist = fmt.Sprintf(`src/locale/%s.xlf`, context.Locale)
			}

			e := context.Update()
			if e != nil {
				log.Fatalln(e)
			}
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&context.Src,
		"src", "s",
		"src/messages.xlf",
		"source file,use ng xi18n get it.",
	)
	flags.StringVarP(&context.Dist,
		"dist", "d",
		"",
		"distribution file,if empty use src/locale/$locale.xlf",
	)
	flags.StringVarP(&context.Locale,
		"locale", "l",
		"",
		"locale zh-Hant zh-Hant-TW zh-Hant-HK ...",
	)
	rootCmd.AddCommand(cmd)
}
