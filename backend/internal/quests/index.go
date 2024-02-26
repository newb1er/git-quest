package quests

var Quests map[string]Quest = map[string]Quest{
	NewQuestCommit().Title: NewQuestCommit(),
	NewQuestBranch().Title: NewQuestBranch(),
}
