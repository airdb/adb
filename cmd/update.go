package cmd

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/minio/selfupdate"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:                "update",
	Short:              "Self update adb",
	Long:               "Self update adb",
	DisableFlagParsing: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return doUpdate()
	},
}

const updateTimeout = 5 * time.Minute

func downloadURL() string {
	dl := "https://github.com/airdb/adb/releases/latest/download/adb"

	// Keep the legacy artifact names for amd64: "adb" (linux) and
	// "adb-darwin". Other combinations use "adb-<goos>-<goarch>".
	switch {
	case runtime.GOOS == "linux" && runtime.GOARCH == "amd64":
	case runtime.GOOS == "darwin" && runtime.GOARCH == "amd64":
		dl += "-darwin"
	default:
		dl += "-" + runtime.GOOS + "-" + runtime.GOARCH
	}

	return dl
}

func doUpdate() error {
	dl := downloadURL()

	fmt.Printf("It will take about 1 minute for downloading.\nDownload url: %s\n", dl)

	client := &http.Client{Timeout: updateTimeout}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, dl, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download %s: unexpected status %s", dl, resp.Status)
	}

	if err := selfupdate.Apply(resp.Body, selfupdate.Options{}); err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	fmt.Println("update successfully!")

	return nil
}

func updateCmdInit() {
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(completionBashCmd)

	completionBashCmd.PersistentFlags().BoolVarP(&writeCompletionFile, "write_file", "w", false, "write completion file")
}

var writeCompletionFile bool

const bashCompletionCompatShim = `# adb bash-completion compatibility shim.
if ! declare -F _get_comp_words_by_ref >/dev/null 2>&1; then
_get_comp_words_by_ref()
{
    local exclude cur_ prev_ cword_
    local words_=( "${COMP_WORDS[@]}" )

    while [[ "$1" == -* ]]; do
        case "$1" in
            -n)
                exclude="$2"
                shift 2
                ;;
            *)
                shift
                ;;
        esac
    done

    cword_=$COMP_CWORD
    cur_="${COMP_WORDS[COMP_CWORD]}"
    if (( COMP_CWORD > 0 )); then
        prev_="${COMP_WORDS[COMP_CWORD-1]}"
    fi

    if [[ -n ${exclude} ]]; then
        local filtered=()
        local i word
        cword_=0
        for i in "${!words_[@]}"; do
            word="${words_[i]}"
            if [[ -n ${word} && ${exclude} == *${word:0:1}* ]]; then
                if (( i < COMP_CWORD )); then
                    ((cword_--))
                fi
                continue
            fi
            filtered+=("${word}")
        done
        words_=( "${filtered[@]}" )
        cur_="${words_[cword_]}"
        if (( cword_ > 0 )); then
            prev_="${words_[cword_-1]}"
        else
            prev_=""
        fi
    fi

    while [[ $# -gt 0 ]]; do
        case "$1" in
            cur)
                printf -v "$1" '%s' "$cur_"
                ;;
            prev)
                printf -v "$1" '%s' "$prev_"
                ;;
            words)
                eval "$1=(\"\${words_[@]}\")"
                ;;
            cword)
                printf -v "$1" '%s' "$cword_"
                ;;
        esac
        shift
    done
}
fi

`

var completionBashCmdLongDesc = `To load completion run

. <(adb completion)

To configure your bash shell to load completions for each session add to your bashrc

# MacOS:
# adb completion >/usr/local/etc/bash_completion.d/adb
# ~/.bashrc or ~/.profile
. <(adb completion)
`

func writeBashCompletionScript(w io.Writer) error {
	if _, err := io.WriteString(w, bashCompletionCompatShim); err != nil {
		return err
	}

	return rootCmd.GenBashCompletionV2(w, false)
}

// CompletionCmd represents the completion command.
var completionBashCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates bash completion scripts",
	Long:  completionBashCmdLongDesc,
	RunE: func(cmd *cobra.Command, args []string) error {
		if writeCompletionFile {
			completionFile := "/usr/local/etc/bash_completion.d/adb"

			file, err := os.Create(completionFile)
			if err != nil {
				return err
			}
			defer file.Close()

			if err := writeBashCompletionScript(file); err != nil {
				return err
			}

			fmt.Println("Generates bash completion scripts successfully, file:", completionFile)

			return nil
		}

		return writeBashCompletionScript(os.Stdout)
	},
}
