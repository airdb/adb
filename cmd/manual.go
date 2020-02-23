package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var manCommand = &cobra.Command{
	Use:   "man",
	Short: "man command like linux",
	Long:  "display manual pages",
}

var gitInitCommand = &cobra.Command{
	Use:   "git",
	Short: "git operation",
	Long:  "git operation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git config --global core.hooksPath .github/hooks")
	},
}

var dockerInitCommand = &cobra.Command{
	Use:   "docker",
	Short: "docker operation",
	Long:  "docker operation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("docker exec -it -e COLUMNS=`tput cols` -e LINES=`tput lines`  -it airdb/go bash")
		fmt.Println("docker exec -it -e COLUMNS=$(tput cols) -e LINES=$(tput lines) airdb/go bash")
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

var toolsInitCommand = &cobra.Command{
	Use:   "tools",
	Short: "install binary tools",
	Long:  "install binary tools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wget https://github.com/airdb/init/releases/latest/download/tools-linux-amd64.zip")
	},
}

var brewInitCommand = &cobra.Command{
	Use:   "brew",
	Short: "brew operation",
	Long:  "brew operation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Brew Common Command:")
		fmt.Println("\tbrew outdated")
		fmt.Println("\tbrew install github/gh/gh")
		fmt.Println("\tbrew install aliyun-cli")
		fmt.Println()
		fmt.Println("\tbrew tap aidb/taps")
		fmt.Println("\tbrew install airdb/taps/adb")
		fmt.Println("\tbrew install adb")
		fmt.Println()
	},
}

var githubInitCommand = &cobra.Command{
	Use:   "gh",
	Short: "github operation",
	Long:  "github operation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("brew install github/gh/gh")
		fmt.Println("gh --repo bbhj/lbs issue status")
		fmt.Println("gh --repo bbhj/lbs issue view 1")
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
