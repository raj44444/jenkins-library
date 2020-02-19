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

type abapEnvironmentPullGitRepoOptions struct {
	Username          string `json:"username,omitempty"`
	Password          string `json:"password,omitempty"`
	RepositoryName    string `json:"repositoryName,omitempty"`
	Host              string `json:"host,omitempty"`
	CfAPIEndpoint     string `json:"cfApiEndpoint,omitempty"`
	CfOrg             string `json:"cfOrg,omitempty"`
	CfSpace           string `json:"cfSpace,omitempty"`
	CfServiceInstance string `json:"cfServiceInstance,omitempty"`
	CfServiceKey      string `json:"cfServiceKey,omitempty"`
}

// AbapEnvironmentPullGitRepoCommand Pulls a git repository to a SAP Cloud Platform ABAP Environment system
func AbapEnvironmentPullGitRepoCommand() *cobra.Command {
	metadata := abapEnvironmentPullGitRepoMetadata()
	var stepConfig abapEnvironmentPullGitRepoOptions
	var startTime time.Time

	var createAbapEnvironmentPullGitRepoCmd = &cobra.Command{
		Use:   "abapEnvironmentPullGitRepo",
		Short: "Pulls a git repository to a SAP Cloud Platform ABAP Environment system",
		Long: `Pulls a git repository (Software Component) to a SAP Cloud Platform ABAP Environment system.
Please provide either of the following options:

* The host and credentials the Cloud Platform ABAP Environment system itself. The credentials must be configured for the Communication Scenario SAP_COM_0510.
* The Cloud Foundry parameters (API endpoint, organization, space), credentials, the service instance for the ABAP service and the service key for the Communication Scenario SAP_COM_0510.
* Only provide one of those options with the respective credentials. If all values are provided, the direct communication (via host) has priority.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			startTime = time.Now()
			log.SetStepName("abapEnvironmentPullGitRepo")
			log.SetVerbose(GeneralConfig.Verbose)
			return PrepareConfig(cmd, &metadata, "abapEnvironmentPullGitRepo", &stepConfig, config.OpenPiperFile)
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
			telemetry.Initialize(GeneralConfig.NoTelemetry, "abapEnvironmentPullGitRepo")
			abapEnvironmentPullGitRepo(stepConfig, &telemetryData)
			telemetryData.ErrorCode = "0"
		},
	}

	addAbapEnvironmentPullGitRepoFlags(createAbapEnvironmentPullGitRepoCmd, &stepConfig)
	return createAbapEnvironmentPullGitRepoCmd
}

func addAbapEnvironmentPullGitRepoFlags(cmd *cobra.Command, stepConfig *abapEnvironmentPullGitRepoOptions) {
	cmd.Flags().StringVar(&stepConfig.Username, "username", os.Getenv("PIPER_username"), "User for either the Cloud Foundry API or the Communication Arrangement for SAP_COM_0510")
	cmd.Flags().StringVar(&stepConfig.Password, "password", os.Getenv("PIPER_password"), "Password for either the Cloud Foundry API or the Communication Arrangement for SAP_COM_0510")
	cmd.Flags().StringVar(&stepConfig.RepositoryName, "repositoryName", os.Getenv("PIPER_repositoryName"), "Specifies the name of the Repository (Software Component) on the SAP Cloud Platform ABAP Environment system")
	cmd.Flags().StringVar(&stepConfig.Host, "host", os.Getenv("PIPER_host"), "Specifies the host address of the SAP Cloud Platform ABAP Environment system")
	cmd.Flags().StringVar(&stepConfig.CfAPIEndpoint, "cfApiEndpoint", os.Getenv("PIPER_cfApiEndpoint"), "Cloud Foundry API Enpoint")
	cmd.Flags().StringVar(&stepConfig.CfOrg, "cfOrg", os.Getenv("PIPER_cfOrg"), "Cloud Foundry target organization")
	cmd.Flags().StringVar(&stepConfig.CfSpace, "cfSpace", os.Getenv("PIPER_cfSpace"), "Cloud Foundry target space")
	cmd.Flags().StringVar(&stepConfig.CfServiceInstance, "cfServiceInstance", os.Getenv("PIPER_cfServiceInstance"), "Cloud Foundry Service Instance")
	cmd.Flags().StringVar(&stepConfig.CfServiceKey, "cfServiceKey", os.Getenv("PIPER_cfServiceKey"), "Cloud Foundry Service Key")

	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("password")
	cmd.MarkFlagRequired("repositoryName")
}

// retrieve step metadata
func abapEnvironmentPullGitRepoMetadata() config.StepData {
	var theMetaData = config.StepData{
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name:        "username",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "password",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "repositoryName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "host",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "cfApiEndpoint",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "cloudFoundry/apiEndpoint"}},
					},
					{
						Name:        "cfOrg",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "cloudFoundry/org"}},
					},
					{
						Name:        "cfSpace",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "cloudFoundry/space"}},
					},
					{
						Name:        "cfServiceInstance",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "cloudFoundry/serviceInstance"}},
					},
					{
						Name:        "cfServiceKey",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "cloudFoundry/serviceKey"}},
					},
				},
			},
		},
	}
	return theMetaData
}
