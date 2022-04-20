package cmd

import (
	"fmt"

	"github.com/airdb/adb/internal/adblib"
	"github.com/spf13/cobra"
)

var manCommand = &cobra.Command{
	Use:     "man",
	Short:   "man command like linux",
	Long:    "display manual pages",
	Example: adblib.MysqlDoc,
}

var gitInitCommand = &cobra.Command{
	Use:     "git",
	Short:   "git operation",
	Long:    "git operation",
	Aliases: []string{"github", "gh"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(adblib.GithubDoc)
	},
}

var dockerInitCommand = &cobra.Command{
	Use:     "docker",
	Short:   "docker operation",
	Long:    "docker operation",
	Example: adblib.DockerDoc,
	Aliases: []string{"podman"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(adblib.DockerDoc)
	},
}

var cloudInitCommand = &cobra.Command{
	Use:   "cloud",
	Short: "cloud operation",
	Long:  "cloud operation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("brew install aliyun-cli")
		fmt.Println("https://github.com/aliyun/aliyun-cli")
	},
}

var brewInitCommand = &cobra.Command{
	Use:   "brew",
	Short: "brew operation",
	Long:  "brew operation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(adblib.BrewDoc)
	},
}

var vimInitCommand = &cobra.Command{
	Use:   "vim",
	Short: "vim configuration",
	Long:  "vim configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("VIM Common Plugins:")
		fmt.Println("")
		fmt.Printf("\tgit clone %s %s\n",
			"http://github.com/fatih/vim-go.git",
			"~/.vim/pack/plugins/start/vim-go",
		)
	},
}

var osinitCommand = &cobra.Command{
	Use:   "osinit",
	Short: "init linux os",
	Long:  "init linux os",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("User-Data for Cloud Server(EC2/CVM/ECS)")
		fmt.Println("")
		fmt.Println("\tcurl https://init.airdb.host/osinit/ubuntu_init.sh | bash -")
	},
}

var kubeCommand = &cobra.Command{
	Use:   "kube",
	Short: "kubeneters command",
	Long:  "kubeneters command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kubeneters commands)")
		fmt.Println("")
		fmt.Println("\tingress")
		fmt.Println("\tkubectl describe ingress")
		fmt.Println("\tkubectl get ingress -o yaml")
		fmt.Println("")
	},
}

var helmCommand = &cobra.Command{
	Use:   "helm",
	Short: "helm command",
	Long:  "helm command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(adblib.HelmDoc)
	},
}

var terraformCommand = &cobra.Command{
	Use:   "terraform",
	Short: "terraform command",
	Long:  "terraform command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(adblib.TerraformDoc)
	},
}

var opensslCommand = &cobra.Command{
	Use:     "openssl",
	Aliases: []string{"ssl", "tls"},
	Short:   "openssl command",
	Long:    "openssl command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(adblib.OpenSSLDoc)
	},
}

var tcpdumpCommand = &cobra.Command{
	Use:     "tcpdump",
	Short:   "tcpdump command",
	Long:    "tcpdump command",
	Aliases: []string{"wireshake"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(adblib.TcpdumpDoc)
	},
}

var toolsCommand = &cobra.Command{
	Use:   "tools",
	Short: "tools command",
	Long:  "tools command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(adblib.ToolsDoc)
	},
}

var perfCommand = &cobra.Command{
	Use:   "perf",
	Short: "application performance",
	Long:  "application performance",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(adblib.PerformanceDoc)
	},
}

var wrkCommand = &cobra.Command{
	Use:   "wrk",
	Short: "wrk performance",
	Long:  "wrk performance",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(adblib.WrkDoc)
	},
}

var s3Command = &cobra.Command{
	Use:     "s3",
	Short:   "s3 tools",
	Long:    "s3 tools",
	Aliases: []string{"cos", "store"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(adblib.S3Doc)
	},
}

var webCommand = &cobra.Command{
	Use:     "webserver",
	Short:   "start a webserver",
	Long:    "start a webserver",
	Aliases: []string{"web"},
	Example: AirdbWiki,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(adblib.WebserverDoc)
	},
}

func initManCommand() {
	rootCmd.AddCommand(manCommand)

	manCommand.AddCommand(gitInitCommand)
	manCommand.AddCommand(dockerInitCommand)
	manCommand.AddCommand(cloudInitCommand)
	manCommand.AddCommand(brewInitCommand)
	manCommand.AddCommand(vimInitCommand)
	manCommand.AddCommand(osinitCommand)
	manCommand.AddCommand(kubeCommand)
	manCommand.AddCommand(helmCommand)
	manCommand.AddCommand(terraformCommand)
	manCommand.AddCommand(opensslCommand)
	manCommand.AddCommand(toolsCommand)
	manCommand.AddCommand(tcpdumpCommand)
	manCommand.AddCommand(perfCommand)
	manCommand.AddCommand(wrkCommand)
	manCommand.AddCommand(s3Command)
	manCommand.AddCommand(webCommand)
}
