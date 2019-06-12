package models

type Option struct {
	Index string `json:"index"`
	Desc  string `json:"desc"`
	// Analyses   []string `json:"analyses"`
	Knowledges []string `json:"knowledges"`
}

type PaperJsonSimple struct {
	IsBasicPackage bool        `json:"isBasicPackage"`
	Pager          Pager       `json:"pager"`
	Paper          Paper       `json:"paper"`
	AnalyseData    AnalyseData `json:"analyseData"`
}

type Pager struct {
	TotalCount int        `json:"totalCount"`
	List       []Question `json:"list"`
}

type Question struct {
	Score       float64        `json:"score"`
	Number      string         `json:"number"`
	Section     Section        `json:"section"`
	Difficulty  Difficulty     `json:"difficulty"`
	Knowledges  []Knowledgeaaa `json:"knowledges"`
	SubQuestion []SubQuestion  `json:"subQuestions"`
	Materials   []Material     `json:"materials"`
}

type Material struct {
	No   string `json:"no"`
	Html string `json:"html"`
}

type Section struct {
	Code         string  `json:"code"`
	Name         string  `json:"name"`
	CategoryCode string  `json:"categoryCode"`
	CategoryName string  `json:"categoryName"`
	Score        float64 `json:"score"`
}

type Difficulty struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Knowledgeaaa struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type SubQuestion struct {
	Stem    string   `json:"stem"`
	Code    string   `json:"code"`
	Options []Option `json:"options"`
	Answers []Answer `json:"answers"`
	Score   float64  `json:"score"`
}

type Answer struct {
	Desc string `json:"desc"`
}

type Paper struct {
	Title         string      `json:"title"`
	Year          string      `json:"year"`
	PaperType     PaperType   `json:"paperType"`
	QuestionCount int         `json:"questionCount"`
	Grade         []Grade     `json:"grade"`
	Phase         Phase       `json:"phase"`
	DefaultArea   DefaultArea `json:"defaultArea"`
	Subject       Subject     `json:"subject"`
	DateTime      int64       `json:"dateTime"`
}

type PaperType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Grade struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Phase struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type DefaultArea struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Subject struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type AnalyseData struct {
	TotalScore              float64                   `json:"totalScore"`
	PaperAmountDistribution []PaperAmountDistribution `json:"paperAmountDistribution"`
	DifficultyAnalysis      []DifficultyAnalysis      `json:"difficultyAnalysis"`
}

type PaperAmountDistribution struct {
	Score                   float64 `json:"score"`
	ScorePercentage         float64 `json:"scorePercentage"`
	QuestionCount           int     `json:"questionCount"`
	QuestionCountPercentage float64 `json:"questionCountPercentage"`
	Name                    string  `json:"Name"`
}

type DifficultyAnalysis struct {
	Score                   float64 `json:"score"`
	ScorePercentage         float64 `json:"scorePercentage"`
	QuestionCount           int     `json:"questionCount"`
	QuestionCountPercentage float64 `json:"questionCountPercentage"`
	DifficultyValue         int     `json:"difficultyValue"`
	DifficultyName          string  `json:"difficultyName"`
}
