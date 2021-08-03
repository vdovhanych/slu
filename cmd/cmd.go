package cmd

import (
	_ "github.com/sikalabs/slut/cmd/expand"
	_ "github.com/sikalabs/slut/cmd/expand/file"
	_ "github.com/sikalabs/slut/cmd/expand/string"
	_ "github.com/sikalabs/slut/cmd/file_templates"
	_ "github.com/sikalabs/slut/cmd/file_templates/editorconfig"
	_ "github.com/sikalabs/slut/cmd/file_templates/gitignore"
	_ "github.com/sikalabs/slut/cmd/generate_docs"
	_ "github.com/sikalabs/slut/cmd/mysql"
	_ "github.com/sikalabs/slut/cmd/mysql/create"
	_ "github.com/sikalabs/slut/cmd/mysql/drop"
	_ "github.com/sikalabs/slut/cmd/mysql/list"
	_ "github.com/sikalabs/slut/cmd/postgres"
	_ "github.com/sikalabs/slut/cmd/postgres/create"
	_ "github.com/sikalabs/slut/cmd/postgres/drop"
	_ "github.com/sikalabs/slut/cmd/postgres/list"
	"github.com/sikalabs/slut/cmd/root"
	_ "github.com/sikalabs/slut/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.RootCmd.Execute())
}
