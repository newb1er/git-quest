package main

import (
	"fmt"
	"git-quest-be/internal/quests"
	"os"

	"github.com/urfave/cli/v2"
)

func getQuest(questName string) (quests.Quest, error) {
	quest, ok := quests.Quests[questName]
	if !ok {
		return nil, fmt.Errorf("quest not found: %s", questName)
	}

	return quest, nil
}

func main() {
	app := &cli.App{
		Name:        "git-quest",
		Description: "Git learning quests",
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List all quests",
				Action: func(c *cli.Context) error {
					for name := range quests.Quests {
						fmt.Println(name)
					}
					return nil
				},
			},
			{
				Name:  "setup",
				Usage: "Setup a quest",
				Action: func(c *cli.Context) error {
					if c.NArg() == 0 {
						return fmt.Errorf("quest name is required")
					}

					questName := c.Args().First()

					quest, err := getQuest(questName)
					if err != nil {
						return err
					}

					err = quest.Setup()
					return err
				},
			},
			{
				Name:  "validate",
				Usage: "Validate a quest",
				Action: func(c *cli.Context) error {
					if c.NArg() == 0 {
						return fmt.Errorf("quest name is required")
					}

					questName := c.Args().First()
					quest, err := getQuest(questName)
					if err != nil {
						return err
					}

					pass, err := quest.Validate()
					if err != nil {
						return err
					}

					if !pass {
						fmt.Println("Quest failed")
					} else {
						fmt.Println("Quest passed")
					}

					return nil
				},
			},
			{
				Name:  "teardown",
				Usage: "Teardown a quest",
				Action: func(c *cli.Context) error {
					if c.NArg() == 0 {
						return fmt.Errorf("quest name is required")
					}

					questName := c.Args().First()
					quest, err := getQuest(questName)
					if err != nil {
						return err
					}

					err = quest.Teardown()
					return err
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
