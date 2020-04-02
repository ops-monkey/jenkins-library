// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type sonarExecuteScanOptions struct {
	Instance                  string `json:"instance,omitempty"`
	Host                      string `json:"host,omitempty"`
	Token                     string `json:"token,omitempty"`
	Organization              string `json:"organization,omitempty"`
	CustomTLSCertificateLinks string `json:"customTlsCertificateLinks,omitempty"`
	SonarScannerDownloadURL   string `json:"sonarScannerDownloadUrl,omitempty"`
	ProjectVersion            string `json:"projectVersion,omitempty"`
	Options                   string `json:"options,omitempty"`
	BranchName                string `json:"branchName,omitempty"`
	ChangeID                  string `json:"changeId,omitempty"`
	ChangeBranch              string `json:"changeBranch,omitempty"`
	ChangeTarget              string `json:"changeTarget,omitempty"`
	PullRequestProvider       string `json:"pullRequestProvider,omitempty"`
	Owner                     string `json:"owner,omitempty"`
	Repository                string `json:"repository,omitempty"`
	GithubToken               string `json:"githubToken,omitempty"`
	DisableInlineComments     bool   `json:"disableInlineComments,omitempty"`
	LegacyPRHandling          bool   `json:"legacyPRHandling,omitempty"`
	GithubAPIURL              string `json:"githubApiUrl,omitempty"`
}

// SonarExecuteScanCommand Executes the Sonar scanner
func SonarExecuteScanCommand() *cobra.Command {
	metadata := sonarExecuteScanMetadata()
	var stepConfig sonarExecuteScanOptions
	var startTime time.Time

	var createSonarExecuteScanCmd = &cobra.Command{
		Use:   "sonarExecuteScan",
		Short: "Executes the Sonar scanner",
		Long:  `The step executes the [sonar-scanner](https://docs.sonarqube.org/display/SCAN/Analyzing+with+SonarQube+Scanner) cli command to scan the defined sources and publish the results to a SonarQube instance.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			startTime = time.Now()
			log.SetStepName("sonarExecuteScan")
			log.SetVerbose(GeneralConfig.Verbose)
			return PrepareConfig(cmd, &metadata, "sonarExecuteScan", &stepConfig, config.OpenPiperFile)
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
			telemetry.Initialize(GeneralConfig.NoTelemetry, "sonarExecuteScan")
			sonarExecuteScan(stepConfig, &telemetryData)
			telemetryData.ErrorCode = "0"
		},
	}

	addSonarExecuteScanFlags(createSonarExecuteScanCmd, &stepConfig)
	return createSonarExecuteScanCmd
}

func addSonarExecuteScanFlags(cmd *cobra.Command, stepConfig *sonarExecuteScanOptions) {
	cmd.Flags().StringVar(&stepConfig.Instance, "instance", "SonarCloud", "Jenkins only: The name of the SonarQube instance defined in Jenkins. DEPRECATED: use host parameter instead")
	cmd.Flags().StringVar(&stepConfig.Host, "host", os.Getenv("PIPER_host"), "The URL to the Sonar backend.")
	cmd.Flags().StringVar(&stepConfig.Token, "token", os.Getenv("PIPER_token"), "Token used to authenticate with the Sonar Server.")
	cmd.Flags().StringVar(&stepConfig.Organization, "organization", os.Getenv("PIPER_organization"), "SonarCloud.io only: Organization that the project will be assigned to in SonarCloud.io.")
	cmd.Flags().StringVar(&stepConfig.CustomTLSCertificateLinks, "customTlsCertificateLinks", os.Getenv("PIPER_customTlsCertificateLinks"), "List of comma-separated download links to custom TLS certificates. This is required to ensure trusted connections to instances with custom certificates.")
	cmd.Flags().StringVar(&stepConfig.SonarScannerDownloadURL, "sonarScannerDownloadUrl", "https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-4.3.0.2102-linux.zip", "URL to the sonar-scanner-cli archive.")
	cmd.Flags().StringVar(&stepConfig.ProjectVersion, "projectVersion", os.Getenv("PIPER_projectVersion"), "The project version that is reported to SonarQube.")
	cmd.Flags().StringVar(&stepConfig.Options, "options", os.Getenv("PIPER_options"), "A list of options which are passed to the sonar-scanner.")
	cmd.Flags().StringVar(&stepConfig.BranchName, "branchName", os.Getenv("PIPER_branchName"), "Pull-Request only: The name of the pull-request branch.")
	cmd.Flags().StringVar(&stepConfig.ChangeID, "changeId", os.Getenv("PIPER_changeId"), "Pull-Request only: The id of the pull-request.")
	cmd.Flags().StringVar(&stepConfig.ChangeBranch, "changeBranch", os.Getenv("PIPER_changeBranch"), "Pull-Request only: The name of the pull-request branch.")
	cmd.Flags().StringVar(&stepConfig.ChangeTarget, "changeTarget", os.Getenv("PIPER_changeTarget"), "Pull-Request only: The name of the base branch.")
	cmd.Flags().StringVar(&stepConfig.PullRequestProvider, "pullRequestProvider", "GitHub", "Pull-Request only: The scm provider.")
	cmd.Flags().StringVar(&stepConfig.Owner, "owner", os.Getenv("PIPER_owner"), "Pull-Request only: The owner of the scm repository.")
	cmd.Flags().StringVar(&stepConfig.Repository, "repository", os.Getenv("PIPER_repository"), "Pull-Request only: The scm repository.")
	cmd.Flags().StringVar(&stepConfig.GithubToken, "githubToken", os.Getenv("PIPER_githubToken"), "Pull-Request only: Token for Github to set status on the Pull-Request.")
	cmd.Flags().BoolVar(&stepConfig.DisableInlineComments, "disableInlineComments", false, "Pull-Request only: Disables the pull-request decoration with inline comments. DEPRECATED: only supported in SonarQube < 7.2")
	cmd.Flags().BoolVar(&stepConfig.LegacyPRHandling, "legacyPRHandling", false, "Pull-Request only: Activates the pull-request handling using the [GitHub Plugin](https://docs.sonarqube.org/display/PLUG/GitHub+Plugin). DEPRECATED: only supported in SonarQube < 7.2")
	cmd.Flags().StringVar(&stepConfig.GithubAPIURL, "githubApiUrl", "https://api.github.com", "Pull-Request only: The URL to the Github API. see [GitHub plugin docs](https://docs.sonarqube.org/display/PLUG/GitHub+Plugin#GitHubPlugin-Usage) DEPRECATED: only supported in SonarQube < 7.2")

}

// retrieve step metadata
func sonarExecuteScanMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:    "sonarExecuteScan",
			Aliases: []config.Alias{},
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name:        "instance",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "host",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "sonarServerUrl"}},
					},
					{
						Name:        "token",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "sonarToken"}},
					},
					{
						Name:        "organization",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "customTlsCertificateLinks",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "sonarScannerDownloadUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "projectVersion",
						ResourceRef: []config.ResourceReference{{Name: "commonPipelineEnvironment", Param: "artifactVersion"}},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "options",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "branchName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "changeId",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "changeBranch",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "changeTarget",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "pullRequestProvider",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "owner",
						ResourceRef: []config.ResourceReference{{Name: "commonPipelineEnvironment", Param: "github/owner"}},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "githubOrg"}},
					},
					{
						Name:        "repository",
						ResourceRef: []config.ResourceReference{{Name: "commonPipelineEnvironment", Param: "github/repository"}},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "githubRepo"}},
					},
					{
						Name:        "githubToken",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "disableInlineComments",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "legacyPRHandling",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "githubApiUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
				},
			},
		},
	}
	return theMetaData
}
