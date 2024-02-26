package quests

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type QuestBranch struct {
	QuestMeta
}

var _ Quest = &QuestBranch{} // Interface Static Check

func NewQuestBranch() *QuestBranch {
	return &QuestBranch{
		QuestMeta: QuestMeta{
			ID:          2,
			Title:       "Branch",
			Description: "Create a new branch",
			repoPath:    "/tmp/branch",
		},
	}
}

func (q *QuestBranch) GetMeta() QuestMeta {
	return q.QuestMeta
}

func (q *QuestBranch) GetPath() string {
	return q.repoPath
}

func (q *QuestBranch) Setup() error {
	// make directory
	err := os.MkdirAll(q.repoPath, 0755)
	if err != nil {
		return fmt.Errorf("make directory failed: %+v", err)
	}

	initOpts := &git.PlainInitOptions{
		InitOptions: git.InitOptions{
			DefaultBranch: plumbing.Main,
		},

		Bare: false,
	}

	_, err = git.PlainInitWithOptions(q.repoPath, initOpts)
	if err != nil {
		return fmt.Errorf("init repo failed: %+v", err)
	}

	return nil
}

func (q *QuestBranch) Validate() (bool, error) {
	// Check if the branch exists
	repo, err := git.PlainOpen(q.repoPath)
	if err != nil {
		return false, fmt.Errorf("open repo failed: %+v", err)
	}

	branches, err := repo.Branches()
	if err != nil {
		return false, fmt.Errorf("get branches failed: %+v", err)
	}

	branchExists := false

	branches.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name().Short() == "feature" {
			branchExists = true
		}
		return nil
	})

	return branchExists, nil
}

func (q *QuestBranch) Teardown() error {
	// Remove the repo
	err := os.RemoveAll(q.repoPath)
	if err != nil {
		return fmt.Errorf("remove repo failed: %+v", err)
	}

	return nil
}

func (q *QuestBranch) Prompt() string {
	return "Create a new branch called 'feature'. Add a empty file named 'feature.txt' and commit it."
}
