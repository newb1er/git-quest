package quests

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type QuestCommit struct {
	QuestMeta
}

var _ Quest = &QuestCommit{} // Interface Static Check

func NewQuestCommit() *QuestCommit {
	return &QuestCommit{
		QuestMeta: QuestMeta{
			ID:          1,
			Title:       "Commit",
			Description: "Commit a new file",
			repoPath:    "/tmp/commit",
		},
	}
}

func (q *QuestCommit) GetMeta() QuestMeta {
	return q.QuestMeta
}

func (q *QuestCommit) GetPath() string {
	return q.repoPath
}

func (q *QuestCommit) Setup() error {
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

func (q *QuestCommit) Teardown() error {
	// Remove the repo
	err := os.RemoveAll(q.repoPath)
	if err != nil {
		return fmt.Errorf("remove repo failed: %+v", err)
	}

	return nil
}

func (q *QuestCommit) Validate() (bool, error) {
	repo, err := git.PlainOpen(q.repoPath)
	if err != nil {
		return false, fmt.Errorf("open git repo failed: %+v", err)
	}

	head, err := repo.Head()
	if err != nil {
		return false, fmt.Errorf("get head failed: %+v", err)
	}

	logOptions := &git.LogOptions{
		From: head.Hash(),
	}
	commits, err := repo.Log(logOptions)
	if err != nil {
		return false, fmt.Errorf("get commits failed: %+v", err)
	}

	lastCommit, err := commits.Next()
	if err != nil {
		return false, fmt.Errorf("get last commit failed: %+v", err)
	}

	helloFile, err := lastCommit.File("hello.txt")
	if err != nil {
		return false, fmt.Errorf("get hello.txt failed: %+v", err)
	}

	content, err := helloFile.Contents()
	if err != nil {
		return false, fmt.Errorf("get hello.txt content failed: %+v", err)
	}

	content = strings.TrimRight(content, "\n")
	if content != "Hello World" {
		return false, fmt.Errorf("hello.txt content is not 'Hello World'")
	}

	_, err = commits.Next()
	if err == nil {
		return false, fmt.Errorf("more than one commit")
	} else if err != io.EOF {
		return false, fmt.Errorf("get next commit failed: %+v", err)
	}

	return true, nil
}

func (q *QuestCommit) Prompt() string {
	return "Please commit a new file 'hello.txt' with the content 'Hello World'"
}
