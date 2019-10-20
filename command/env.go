package command

import (
	"fmt"
)

func env() error {
	fmt.Println("adb_config=~/.adb/config.json")
	fmt.Println("git_hook_path: ./.github/hooks/")
	fmt.Println("aliyun_config: ~/.aliyun/config.json")
	// fmt.Println("git config --list")
	return nil
}
