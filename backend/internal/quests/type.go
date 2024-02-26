package quests

type QuestMeta struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	repoPath    string
}

type Quest interface {
	Setup() error
	Validate() (bool, error)
	Teardown() error
	GetMeta() QuestMeta
	GetPath() string
	Prompt() string
}
