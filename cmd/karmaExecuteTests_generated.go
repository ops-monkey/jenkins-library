package cmd

import (
	"fmt"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type karmaExecuteTestsOptions struct {
	InstallCommand string `json:"installCommand,omitempty"`
	ModulePath     string `json:"modulePath,omitempty"`
	RunCommand     string `json:"runCommand,omitempty"`
}

// KarmaExecuteTestsCommand Executes the Karma test runner
func KarmaExecuteTestsCommand() *cobra.Command {
	metadata := karmaExecuteTestsMetadata()
	var stepConfig karmaExecuteTestsOptions
	var startTime time.Time

	var createKarmaExecuteTestsCmd = &cobra.Command{
		Use:   "karmaExecuteTests",
		Short: "Executes the Karma test runner",
		Long: `In this step the ([Karma test runner](http://karma-runner.github.io)) is executed.

The step is using the ` + "`" + `seleniumExecuteTest` + "`" + ` step to spin up two containers in a Docker network:

* a Selenium/Chrome container (` + "`" + `selenium/standalone-chrome` + "`" + `)
* a NodeJS container (` + "`" + `node:lts-stretch` + "`" + `)

In the Docker network, the containers can be referenced by the values provided in ` + "`" + `dockerName` + "`" + ` and ` + "`" + `sidecarName` + "`" + `, the default values are ` + "`" + `karma` + "`" + ` and ` + "`" + `selenium` + "`" + `. These values must be used in the ` + "`" + `hostname` + "`" + ` properties of the test configuration ([Karma](https://karma-runner.github.io/1.0/config/configuration-file.html) and [WebDriver](https://github.com/karma-runner/karma-webdriver-launcher#usage)).

!!! note
    In a Kubernetes environment, the containers both need to be referenced with ` + "`" + `localhost` + "`" + `.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			startTime = time.Now()
			log.SetStepName("karmaExecuteTests")
			log.SetVerbose(GeneralConfig.Verbose)
			return PrepareConfig(cmd, &metadata, "karmaExecuteTests", &stepConfig, config.OpenPiperFile)
		},
		Run: func(cmd *cobra.Command, args []string) {
			telemetryData := telemetry.CustomData{}
			telemetryData.ErrorCode = "1"
			handler := func() {
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetry.Send(&telemetryData)
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, "karmaExecuteTests")
			karmaExecuteTests(stepConfig, &telemetryData)
			telemetryData.ErrorCode = "0"
		},
	}

	addKarmaExecuteTestsFlags(createKarmaExecuteTestsCmd, &stepConfig)
	return createKarmaExecuteTestsCmd
}

func addKarmaExecuteTestsFlags(cmd *cobra.Command, stepConfig *karmaExecuteTestsOptions) {
	cmd.Flags().StringVar(&stepConfig.InstallCommand, "installCommand", "npm install --quiet", "The command that is executed to install the test tool.")
	cmd.Flags().StringVar(&stepConfig.ModulePath, "modulePath", ".", "Define the path of the module to execute tests on.")
	cmd.Flags().StringVar(&stepConfig.RunCommand, "runCommand", "npm run karma", "The command that is executed to start the tests.")

	cmd.MarkFlagRequired("installCommand")
	cmd.MarkFlagRequired("modulePath")
	cmd.MarkFlagRequired("runCommand")
}

// retrieve step metadata
func karmaExecuteTestsMetadata() config.StepData {
	var theMetaData = config.StepData{
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name:        "installCommand",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "modulePath",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "runCommand",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
				},
			},
		},
	}
	return theMetaData
}