package main

import (
	"os"

	"github.com/urfave/cli"
)

const help = `

`

func main() {
	app := cli.NewApp()
	app.Name = name
	app.Usage = usage
	app.UsageText = usageText
	app.Version = FmtVersion()
	app.Before = func(c *cli.Context) error {
		repo, err := OpenRepository(".")
		if err != nil {
			return exit(err)
		}
		setBacklogRepositoryToContext(c, NewBacklogRepository(repo))
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:  "pr",
			Usage: "Open the pull request list page in current repository",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "s, state",
					Value: "open",
				},
			},
			Action: func(c *cli.Context) error {
				s := c.String("state")
				return exit(getBacklogRepositoryFromContext(c).OpenPullRequestList(s))
			},
			Subcommands: []cli.Command{
				{
					Name:  "show",
					Usage: "Open the pull request page related to current branch",
					Action: func(c *cli.Context) error {
						return exit(getBacklogRepositoryFromContext(c).OpenPullRequest())
					},
				},
				{
					Name:  "add",
					Usage: "Open the page to add pull request with current branch",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "b, base",
						},
					},
					Action: func(c *cli.Context) error {
						base := c.String("base")
						return exit(getBacklogRepositoryFromContext(c).OpenAddPullRequest(base, ""))
					},
				},
			},
		},
		{
			Name:  "issue",
			Usage: "Open the issue list page in current project",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "s, state",
					Value: "not_closed",
				},
			},
			Action: func(c *cli.Context) error {
				s := c.String("state")
				return exit(getBacklogRepositoryFromContext(c).OpenIssueList(s))
			},
			Subcommands: []cli.Command{
				{
					Name:  "show",
					Usage: "Open the issue page related to current branch",
					Action: func(c *cli.Context) error {
						return exit(getBacklogRepositoryFromContext(c).OpenIssue())
					},
				},
				{
					Name:  "add",
					Usage: "Open the page to add issue in current repository's project",
					Action: func(c *cli.Context) error {
						return getBacklogRepositoryFromContext(c).OpenAddIssue()
					},
				},
			},
		},
		{
			Name:  "branch",
			Usage: "Open the branch list page in current repository",
			Action: func(c *cli.Context) error {
				return exit(getBacklogRepositoryFromContext(c).OpenBranchList())
			},
		},
		{
			Name:  "tag",
			Usage: "Open the tag list page in current repository",
			Action: func(c *cli.Context) error {
				return exit(getBacklogRepositoryFromContext(c).OpenTagList())
			},
		},
		{
			Name:  "tree",
			Usage: "Open the tree page in current branch",
			Action: func(c *cli.Context) error {
				return exit(getBacklogRepositoryFromContext(c).OpenTree(""))
			},
		},
		{
			Name:  "history",
			Usage: "Open the history page in current branch",
			Action: func(c *cli.Context) error {
				return exit(getBacklogRepositoryFromContext(c).OpenHistory(""))
			},
		},
		{
			Name:  "network",
			Usage: "Open the network page in current branch",
			Action: func(c *cli.Context) error {
				return exit(getBacklogRepositoryFromContext(c).OpenNetwork(""))
			},
		},
		{
			Name:  "repo",
			Usage: "Open the repository list page in current project",
			Action: func(c *cli.Context) error {
				return exit(getBacklogRepositoryFromContext(c).OpenRepositoryList())
			},
		},
	}
	app.Run(os.Args)
}

const contextKeyBacklogRepository = "ctx-key-backlog-repository"

func getBacklogRepositoryFromContext(c *cli.Context) *BacklogRepository {
	v, ok := c.App.Metadata[contextKeyBacklogRepository]
	if !ok {
		return nil
	}
	if repo, ok := v.(*BacklogRepository); !ok {
		return nil
	} else {
		return repo
	}
}

func setBacklogRepositoryToContext(c *cli.Context, b *BacklogRepository) {
	c.App.Metadata[contextKeyBacklogRepository] = b
}

func exit(err error) error {
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	return nil
}
