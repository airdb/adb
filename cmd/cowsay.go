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
const cowsayFortuneFileName = "fortune.txt"

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
	"凡是教你赚钱的，一定是想赚你的钱，把这句话焊死在脑子里",
	"假如有一天你没钱了，谁会出来为你兜底",
	"没钱的时候，你才知道自己到底有几个真朋友",
	"穷一次，你就看清了所有人的真面目",
	"别人对你的脸色，取决于你兜里的厚度",
	"钱是照妖镜，一照一个准",
	"有钱时身边全是好人，没钱时身边全是路人",
	"别高估感情，别低估人性",
	"落难时开口借钱，是检验人情最狠的办法",
	"钱不是万能的，但没钱你连说这句话的资格都没有",
	"人走茶凉是常态，别把凉茶当热茶捧着",
	"你以为的知己，多半只是你的一厢情愿",
	"落魄时还在你身边的，才配得上你富贵时的惦记",
	"别人帮你是情分，不帮你是本分",
	"别把客气当真心，别把敷衍当热情",
	"成年人的告别，从来都是悄无声息的",
	"你成功时鼓掌的人多，你失败时递手的人少",
	"关系再铁，也经不起利益反复称量",
	"别指望雪中送炭，多数人只会锦上添花",
	"你把他当兄弟，他把你当人脉",
	"靠山山会倒，靠人人会跑，靠自己最牢",
	"你的价值，决定了别人对你客气的程度",
	"与其求人，不如让自己变得值钱",
	"人脉不是你认识谁，是你能帮到谁",
	"别把平台当本事，离了平台你还剩什么",
	"真正的成熟，是学会一个人扛下所有",
	"别抱怨没人懂你，先问问自己值不值得被懂",
	"你能走多远，取决于你熬过了多少黑夜",
	"别在低谷时认命，也别在高峰时忘形",
	"你不强大，没人为你坚强",
	"别活在别人的嘴里，你不是他们的谈资",
	"讨厌你的人，你连呼吸都是错的",
	"别人夸你，别当真；别人骂你，别上心",
	"真正关心你的没几个，看热闹的一大把",
	"别把别人的客气，当成喜欢",
	"你过得好有人眼红，你过得差有人偷笑",
	"别解释，懂你的人不需要，不懂的人不配",
	"世界不关心你的自尊，它只认你的结果",
	"别人的成功学是别人的剧本，不是你的答案",
	"别拿别人的尺子，量自己的人生",
	"成年人的世界，没有容易二字",
	"生活不会因为你惨，就对你手下留情",
	"理想很丰满，现实很骨感",
	"别天真了，社会不是学校，没人给你补考",
	"这个世界的真相，往往没那么体面",
	"你以为的岁月静好，都有人在替你扛",
	"别把运气当实力，别把时代红利当能力",
	"免费的，往往是最贵的",
	"天上不会掉馅饼，掉下来的多半是陷阱",
	"越是漂亮的话，越要留个心眼",
	"教你暴富的，十有八九想暴富你",
	"轻松赚钱的门路，都写在刑法里",
	"没有躺赚，只有躺平了被人赚",
	"生意场上，没有永远的朋友，只有永远的利益",
	"合同比人情靠谱，白纸黑字比承诺硬",
	"别赚最后一个铜板，也别信稳赚不赔",
	"便宜没好货，好货不便宜，人性也一样",
	"先小人后君子，丑话说在前头不丢人",
	"别人跟你谈感情，多半是想谈价格",
	"免费教你发财的，盯的是你的本金",
	"健康是1，其他都是0，1倒了全是0",
	"你熬的每一个夜，都在透支明天",
	"别等病了才想健康，别等老了才后悔没拼过",
	"命是自己的，别人不会替你疼",
	"身体是革命的本钱，垮了就什么都没了",
	"别用健康换钱，老了拿钱换不回健康",
	"今天偷的懒，明天加倍还",
	"你现在的样子，是过去所有选择的总和",
	"时间最公平，也最贵",
	"早点睡，明天还有硬仗要打",
	"孤独是常态，学会和自己相处",
	"你的底气，来自卡里的数字和脑子里的东西",
	"手里有粮，心里不慌；手里有钱，说话才硬",
	"别把希望寄托在别人身上，那是最不靠谱的投资",
	"安全感是自己给的，不是别人施舍的",
	"站得稳，是因为你兜得住",
	"求人不如求己，靠人不如靠己",
	"你唯一能依靠的，是那个没被生活打倒的自己",
	"没人会替你的人生负责，清醒点",
	"跌倒了自己爬起来，别等着谁来扶",
	"你对别人有用，别人才对你客气",
	"成年人的社交，本质是价值交换",
	"别指望别人理解你，每个人都在忙自己",
	"你以为的将心比心，可能换来得寸进尺",
	"善良要有锋芒，否则就是懦弱",
	"别把时间浪费在无效社交上",
	"圈子不同，不必强融",
	"低质量的社交，不如高质量的独处",
	"你的善良，得带点刺",
	"别把真心喂了狗，更别喂了不熟的狼",
	"情绪是最没用的东西，先解决问题",
	"别让情绪替你做了决定",
	"冷静，是成年人最高级的能力",
	"遇到烂人，及时止损，别纠缠",
	"别拿别人的错，惩罚自己",
	"想不通的事先放一放，时间会给你答案",
	"别在没意义的事上较真，那是跟自己过不去",
	"成年人的体面，是摔倒了也笑着站起来",
	"日子再难，也要把自己收拾干净",
	"把命运攥在自己手里，别交给任何人",
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
		fortunes, err := loadCowsayFortunes()
		if err != nil {
			return err
		}
		message = fortunes[rand.IntN(len(fortunes))]
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
		if filepath.ToSlash(relativePath) == cowsayFortuneFileName {
			return nil
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

func loadCowsayFortunes() ([]string, error) {
	homeDir, err := cowsayUserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("resolve home directory: %w", err)
	}

	customFortunes, err := readCowsayFortuneFile(filepath.Join(homeDir, cowsayCustomFigureDir, cowsayFortuneFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return cowsayFortunes, nil
		}
		return nil, fmt.Errorf("read custom cowsay fortunes: %w", err)
	}
	if len(customFortunes) == 0 {
		return cowsayFortunes, nil
	}

	return customFortunes, nil
}

func readCowsayFortuneFile(path string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	normalized := strings.ReplaceAll(string(content), "\r\n", "\n")
	lines := strings.Split(normalized, "\n")
	fortunes := make([]string, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fortunes = append(fortunes, line)
	}

	return fortunes, nil
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
