package vss

// Event types values
const (
	DeploymentStartedEvent           string = "ms.vss-release.deployment-started-event"
	DeploymentApprovalPendingEvent   string = "ms.vss-release.deployment-approval-pending-event"
	DeploymentApprovalCompletedEvent string = "ms.vss-release.deployment-approval-completed-event"
	DeploymentCompletedEvent         string = "ms.vss-release.deployment-completed-event"
)

// Event types values
const (
	ArtifactBuildType                string = "Build"
	ArtifactJenkinsType              string = "Jenkins"
	ArtifactGitHubType               string = "GitHub"
	ArtifactNugetType                string = "Nuget"
	ArtifactTeamBuildType            string = "Team Build (external)"
	ArtifactExternalTFSBuildType     string = "ExternalTFSBuild"
	ArtifactGitType                  string = "Git"
	ArtifactTFVCType                 string = "TFVC"
	ArtifactExternalTfsXamlBuildType string = "ExternalTfsXamlBuild"
)
