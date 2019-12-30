package cmd

import (
	"fmt"
	"runtime"

	"github.com/imroc/req"
	"github.com/spf13/cobra"
)

type GithubReleaseInfo struct {
	Assets []struct {
		BrowserDownloadURL string      `json:"browser_download_url"`
		ContentType        string      `json:"content_type"`
		CreatedAt          string      `json:"created_at"`
		DownloadCount      int         `json:"download_count"`
		ID                 int         `json:"id"`
		Label              interface{} `json:"label"`
		Name               string      `json:"name"`
		NodeID             string      `json:"node_id"`
		Size               int         `json:"size"`
		State              string      `json:"state"`
		UpdatedAt          string      `json:"updated_at"`
		Uploader           struct {
			AvatarURL         string `json:"avatar_url"`
			EventsURL         string `json:"events_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			GravatarID        string `json:"gravatar_id"`
			HTMLURL           string `json:"html_url"`
			ID                int    `json:"id"`
			Login             string `json:"login"`
			NodeID            string `json:"node_id"`
			OrganizationsURL  string `json:"organizations_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			ReposURL          string `json:"repos_url"`
			SiteAdmin         bool   `json:"site_admin"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			Type              string `json:"type"`
			URL               string `json:"url"`
		} `json:"uploader"`
		URL string `json:"url"`
	} `json:"assets"`
	AssetsURL string `json:"assets_url"`
	Author    struct {
		AvatarURL         string `json:"avatar_url"`
		EventsURL         string `json:"events_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		GravatarID        string `json:"gravatar_id"`
		HTMLURL           string `json:"html_url"`
		ID                int    `json:"id"`
		Login             string `json:"login"`
		NodeID            string `json:"node_id"`
		OrganizationsURL  string `json:"organizations_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		ReposURL          string `json:"repos_url"`
		SiteAdmin         bool   `json:"site_admin"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		Type              string `json:"type"`
		URL               string `json:"url"`
	} `json:"author"`
	Body            string `json:"body"`
	CreatedAt       string `json:"created_at"`
	Draft           bool   `json:"draft"`
	HTMLURL         string `json:"html_url"`
	ID              int    `json:"id"`
	Name            string `json:"name"`
	NodeID          string `json:"node_id"`
	Prerelease      bool   `json:"prerelease"`
	PublishedAt     string `json:"published_at"`
	TagName         string `json:"tag_name"`
	TarballURL      string `json:"tarball_url"`
	TargetCommitish string `json:"target_commitish"`
	UploadURL       string `json:"upload_url"`
	URL             string `json:"url"`
	ZipballURL      string `json:"zipball_url"`
}

var installCommand = &cobra.Command{
	Use:   "install",
	Short: "Install Binnary From Github.",
	Long:  "Install Binnary From Github.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			installUsage()
		} else {
			installToUsrLocalBin(args[0])
		}
	},
}

func installUsage() {
	fmt.Printf("Usage: adb install [package_name]\n")
	fmt.Println()
	fmt.Printf("Package Support List:\n")
	fmt.Printf("\tkubectl-iexec\n")
	fmt.Printf("\thelm\n")
}

func installToUsrLocalBin(packageName string) {
	apiurl := "https://api.github.com/repos/"
	switch packageName {
	case "kubectl-iexec":
		apiurl += "gabeduke/kubectl-iexec"
	case "helm":
		apiurl += "helm/helm"
	default:
		installUsage()
		return
	}

	apiurl += "/releases/latest"
	resp, err := req.Get(apiurl)
	if err != nil {
		fmt.Println("install package failed")
		return
	}

	var release GithubReleaseInfo
	err = resp.ToJSON(&release)
	if err != nil {
		fmt.Println("install package failed")
		return
	}

	name := ""
	downloadURL := ""
	switch packageName {
	case "kubectl-iexec":
		name = fmt.Sprintf("kubectl-iexec_%s_%s_x86_64.tar.gz",
			release.TagName,
			runtime.GOOS,
		)
		downloadURL = fmt.Sprintf("https://github.com/gabeduke/kubectl-iexec/releases/download/%s/%s",
			release.TagName,
			name,
		)
	case "helm":
		name = fmt.Sprintf("helm-%s-%s-%s.tar.gz",
			release.TagName,
			runtime.GOOS,
			runtime.GOARCH,
		)
		downloadURL = fmt.Sprintf("https://get.helm.sh/%s",
			name,
		)
	default:
		installUsage()
		return
	}

	fmt.Println(name)
	fmt.Println(apiurl)
	fmt.Println(downloadURL)
	downloadAndUnzip(name, downloadURL)

}

func downloadAndUnzip(name, downloadUrl string) {
	resp, err := req.Get(downloadUrl)
	if err != nil {
		fmt.Println("download package failed")
		return
	}
	filePath := "/tmp/" + name
	err = resp.ToFile(filePath)
	if err != nil {
		fmt.Println("download package failed")
	}
}
