// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-api/stainless-api-go"
	"github.com/stainless-api/stainless-api-go/option"
	"github.com/urfave/cli/v3"
)

var projectsConfigsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Retrieve configuration files for a project",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "project",
		},
		&cli.StringFlag{
			Name:   "branch",
			Action: getAPIFlagAction[string]("query", "branch"),
		},
	},
	Before:          initAPICommand,
	Action:          handleProjectsConfigsRetrieve,
	HideHelpCommand: true,
}

var projectsConfigsGuess = cli.Command{
	Name:  "guess",
	Usage: "Generate configuration suggestions based on an OpenAPI spec",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "project",
		},
		&cli.StringFlag{
			Name:   "spec",
			Action: getAPIFlagAction[string]("body", "spec"),
		},
		&cli.StringFlag{
			Name:   "branch",
			Action: getAPIFlagAction[string]("body", "branch"),
		},
	},
	Before:          initAPICommand,
	Action:          handleProjectsConfigsGuess,
	HideHelpCommand: true,
}

func handleProjectsConfigsRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(ctx, cmd)
	params := stainlessv0.ProjectConfigGetParams{}
	if cmd.IsSet("project") {
		params.Project = stainlessv0.String(cmd.Value("project").(string))
	}
	res := []byte{}
	_, err := cc.client.Projects.Configs.Get(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", colorizeJSON(string(res), os.Stdout))
	return nil
}

func handleProjectsConfigsGuess(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(ctx, cmd)
	params := stainlessv0.ProjectConfigGuessParams{}
	if cmd.IsSet("project") {
		params.Project = stainlessv0.String(cmd.Value("project").(string))
	}
	res := []byte{}
	_, err := cc.client.Projects.Configs.Guess(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithRequestBody("application/json", cc.body),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", colorizeJSON(string(res), os.Stdout))
	return nil
}
