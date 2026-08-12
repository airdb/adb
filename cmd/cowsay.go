package cmd

import (
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

const cowsayWrapWidth = 40
const cowsayCustomFigureDir = ".config/bar"

type cowsayFigure struct {
	body []string
}

var (
	cowsayFigureName string
	cowsayListOnly   bool
)

var cowsayFigures = map[string]cowsayFigure{
	"cow": {
		body: []string{
			"        \\   ^__^",
			"         \\  (oo)\\_______",
			"            (__)\\       )\\/\\",
			"                ||----w |",
			"                ||     ||",
		},
	},
	"tux": {
		body: []string{
			"   \\",
			"    \\",
			"        .--.",
			"       |o_o |",
			"       |:_/ |",
			"      //   \\ \\",
			"     (|     | )",
			"    /'\\_   _/`\\\\",
			"    \\___)=(___/",
		},
	},
	"dragon": {
		body: []string{
			"           \\                    / \\  //\\\\",
			"            \\    |\\___/|      /   \\//  \\\\",
			"                 /0  0  \\__  /    //  | \\ \\",
			"                /     /  \\/_/    //   |  \\  \\",
			"                \\_^_\\'/   \\/_   //    |   \\   \\",
			"                //_^_/     \\/_ //     |    \\    \\",
			"             ( //) |        \\///      |     \\     \\",
			"           ( / /) _|_ /   )  //       |      \\     _\\",
			"         ( // /) '/,_ _ _/  ( ; -.    |    _ _\\.-~        .-~~~^-.",
			"       (( / / )) ,-{        _      `-.|.-~-.           .~         `.",
			"      (( // / ))  '/\\      /                 ~-. _ .-~      .-~^-.  \\",
			"      (( /// ))      `.   {            }                   /      \\  \\",
			"       (( / ))     .----~-.\\        \\-'                 .~         \\  `. \\^-.",
			"                  ///.----..>        \\             _ -~             `.  ^-`  ^-_",
			"                    ///-._ _ _ _ _ _ _}^ - - - - ~                     ~--,   .-~",
			"                                                                       /.-~",
		},
	},
	"bunny": {
		body: []string{
			"        \\",
			"         \\",
			"          (\\_/)",
			"          (o.o)",
			"          /|_|\\\\",
		},
	},
	"sheep": {
		body: []string{
			"        \\",
			"         \\",
			"           __",
			"         .-  '-.",
			"        /        \\",
			"       |  .--.  .-|",
			"       | (    \\/  |",
			"        \\      _.-'",
			"         '-.__.-'",
		},
	},
}

var cowsayUserHomeDir = os.UserHomeDir

var cowsayFortunes = []string{
	"Ship small, iterate fast.",
	"Latency is a feature when it is low enough.",
	"Read the error before you rewrite the code.",
	"Cache invalidation becomes easier after naming things well.",
	"The simplest rollback plan is usually the best plan.",
	"Measure first, then optimize the hot path.",
	"Logs are only useful when they explain the failure.",
	"Healthy tooling saves more time than clever heroics.",
	"Delete dead code before it turns into folklore.",
	"Good defaults remove half the documentation burden.",
	"行而不辍，未来可期。",
	"路虽远，行则将至。",
	"守正而后出新。",
	"见微知著，行稳致远。",
	"于细微处见功夫。",
	"于无声处见担当。",
	"复杂之中，守住分寸。",
	"持续、克制与笃定，自有力量。",
	"真正可贵的是判断与分寸。",
	"不事张扬，自有方向。",
}

var cowsayCommand = &cobra.Command{
	Use:     "cowsay [message]",
	Short:   "Print a random fortune with an ASCII character",
	Long:    "Print a cowsay-style speech bubble. Without a message, adb cowsay picks a random fortune and character.",
	Aliases: []string{"say"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runCowsay(args)
	},
}

func init() {
	cowsayCommand.Flags().StringVarP(&cowsayFigureName, "figure", "f", "", "ASCII figure to use")
	cowsayCommand.Flags().BoolVarP(&cowsayListOnly, "list", "l", false, "List available figures")
}

func runCowsay(args []string) error {
	figures, err := loadCowsayFigures()
	if err != nil {
		return err
	}

	if cowsayListOnly {
		fmt.Println(strings.Join(cowsayFigureNames(figures), "\n"))
		return nil
	}

	message := strings.TrimSpace(strings.Join(args, " "))
	if message == "" {
		message = cowsayFortunes[rand.IntN(len(cowsayFortunes))]
	}

	figureName := cowsayFigureName
	if figureName == "" {
		names := cowsayFigureNames(figures)
		figureName = names[rand.IntN(len(names))]
	}

	figure, ok := figures[figureName]
	if !ok {
		return fmt.Errorf("unknown figure %q, available: %s", figureName, strings.Join(cowsayFigureNames(figures), ", "))
	}

	for _, line := range renderBubble(message) {
		fmt.Println(line)
	}
	for _, line := range figure.body {
		fmt.Println(line)
	}

	return nil
}

func loadCowsayFigures() (map[string]cowsayFigure, error) {
	figures := make(map[string]cowsayFigure, len(cowsayFigures))
	for name, figure := range cowsayFigures {
		figures[name] = figure
	}

	customFigures, err := loadCustomCowsayFigures()
	if err != nil {
		return nil, err
	}
	for name, figure := range customFigures {
		figures[name] = figure
	}

	return figures, nil
}

func loadCustomCowsayFigures() (map[string]cowsayFigure, error) {
	homeDir, err := cowsayUserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("resolve home directory: %w", err)
	}

	baseDir := filepath.Join(homeDir, cowsayCustomFigureDir)
	if _, err := os.Stat(baseDir); err != nil {
		if os.IsNotExist(err) {
			return map[string]cowsayFigure{}, nil
		}
		return nil, fmt.Errorf("stat custom cowsay figure directory %q: %w", baseDir, err)
	}

	figures := map[string]cowsayFigure{}
	walkErr := filepath.WalkDir(baseDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		body, err := readCowsayFigureFile(path)
		if err != nil {
			return fmt.Errorf("read %q: %w", path, err)
		}

		relativePath, err := filepath.Rel(baseDir, path)
		if err != nil {
			return fmt.Errorf("make relative path for %q: %w", path, err)
		}

		figureName := filepath.ToSlash(relativePath)
		figures[figureName] = cowsayFigure{body: body}
		return nil
	})
	if walkErr != nil {
		return nil, fmt.Errorf("load custom cowsay figures from %q: %w", baseDir, walkErr)
	}

	return figures, nil
}

func readCowsayFigureFile(path string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	normalized := strings.ReplaceAll(string(content), "\r\n", "\n")
	normalized = strings.TrimSuffix(normalized, "\n")

	return strings.Split(normalized, "\n"), nil
}

func cowsayFigureNames(figures map[string]cowsayFigure) []string {
	names := make([]string, 0, len(figures))
	for name := range figures {
		names = append(names, name)
	}
	slices.Sort(names)
	return names
}

func renderBubble(message string) []string {
	lines := wrapText(message, cowsayWrapWidth)
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	rendered := []string{
		" " + strings.Repeat("_", maxWidth+2),
	}

	if len(lines) == 1 {
		rendered = append(rendered, fmt.Sprintf("< %-*s >", maxWidth, lines[0]))
	} else {
		for i, line := range lines {
			left, right := "|", "|"
			if i == 0 {
				left, right = "/", "\\"
			} else if i == len(lines)-1 {
				left, right = "\\", "/"
			}
			rendered = append(rendered, fmt.Sprintf("%s %-*s %s", left, maxWidth, line, right))
		}
	}

	rendered = append(rendered, " "+strings.Repeat("-", maxWidth+2))
	return rendered
}

func wrapText(message string, width int) []string {
	rawLines := strings.Split(message, "\n")
	wrapped := make([]string, 0, len(rawLines))

	for _, rawLine := range rawLines {
		words := strings.Fields(rawLine)
		if len(words) == 0 {
			wrapped = append(wrapped, "")
			continue
		}

		current := words[0]
		for _, word := range words[1:] {
			if len(current)+1+len(word) <= width {
				current += " " + word
				continue
			}
			wrapped = append(wrapped, current)
			current = word
		}
		wrapped = append(wrapped, current)
	}

	if len(wrapped) == 0 {
		return []string{""}
	}

	return wrapped
}
