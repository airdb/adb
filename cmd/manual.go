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
		fmt.Println("git config --global core.excludefile .gitignore_global")
		fmt.Println("git config --global url.'ssh://git@github.com'.insteadOf https://github.com")
		fmt.Println()
		fmt.Println("For Close Github Issue, commit message should as follow:")
		fmt.Println("\t", "close #x")
		fmt.Println("\t", "closes #x")
		fmt.Println("\t", "closed #x")
		fmt.Println("\t", "fix #x")
		fmt.Println("\t", "fixes #x")
		fmt.Println("\t", "fixed #x")
		fmt.Println("\t", "resolve #x")
		fmt.Println("\t", "resolves #x")
		fmt.Println("\t", "resolved #x")
		fmt.Println("\t", "add new quick sort algorithm, fixes #4, resolve #6, closed #12")

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
		fmt.Println("\tbrew ")
		fmt.Println("\tbrew cask outdated")
		fmt.Println("\tbrew outdated adb --verbose --debug")
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
		fmt.Println("helm commands)")
		fmt.Println("")
		fmt.Println("\tsudo apt update")
		fmt.Println("\tsudo apt install snapd")
		fmt.Println("\tsudo snap install helm --classic")
		fmt.Println("")
		fmt.Println("\thelm plugin install https://github.com/airdb/helm-kube")
		fmt.Println("")
		fmt.Println("\thelm repo add airdb https://www.airdb.com/helm/")
		fmt.Println("\thelm repo update")
		fmt.Println("\thelm search repo helm/mychart")
		fmt.Println("")
		fmt.Println("\thelm show readme airdb/mychat")
		fmt.Println("")
		fmt.Println("\thelm install chart airdb/mychat")
		fmt.Println("\thelm install chart airdb/mychat --dry-run --debug")
		fmt.Println("")
		fmt.Println("\thelm get notes mychat")
		fmt.Println("")
	},
}

var terraformCommand = &cobra.Command{
	Use:   "terraform",
	Short: "terraform command",
	Long:  "terraform command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("terraform commands")
		fmt.Println("")
		fmt.Println("\tterraform -install-autocomplete")
		fmt.Println("\tterraform init -upgrade")
		fmt.Println("")
		fmt.Println("\tterraform validate")
		fmt.Println("\tterraform plan")
		fmt.Println("\tterraform apply")
		fmt.Println("\t#terraform destroy")
		fmt.Println("")
		fmt.Println("\tRefer: https://github.com/airdb/init/tree/master/terraform")
	},
}

var opensslCommand = &cobra.Command{
	Use:   "openssl",
	Short: "openssl command",
	Long:  "openssl command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("openssl commands")
		fmt.Println("")
		fmt.Println("\topenssl s_client -servername www.airdb.com  -connect www.airdb.com:443 </dev/null 2>/dev/null")
		fmt.Println("\tcert -f md  www.airdb.com")
		fmt.Println("")
		fmt.Println("\tRefer: https://github.com/genkiroid/cert")
	},
}
