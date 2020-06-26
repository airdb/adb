package cmd

type FlagStruct struct {
	Type string
}

var GlobalFlags = FlagStruct{}

type generateStruct struct {
	GitURL string
	Owner  string
	Repo   string
}

var generateFlags = generateStruct{}
