package vss

import "time"

// Release struct
type Release struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	CreatedOn  time.Time `json:"createdOn"`
	ModifiedOn time.Time `json:"modifiedOn"`
	ModifiedBy struct {
		ID          string `json:"id"`
		DisplayName string `json:"displayName"`
		UniqueName  string `json:"uniqueName"`
		URL         string `json:"url"`
		ImageURL    string `json:"imageUrl"`
	} `json:"modifiedBy"`
	CreatedBy struct {
		ID          string `json:"id"`
		DisplayName string `json:"displayName"`
		UniqueName  string `json:"uniqueName"`
		URL         string `json:"url"`
		ImageURL    string `json:"imageUrl"`
	} `json:"createdBy"`
	Environments []struct {
		ID        int    `json:"id"`
		ReleaseID int    `json:"releaseId"`
		Name      string `json:"name"`
		Status    string `json:"status"`
		Variables struct {
		} `json:"variables"`
		PreDeployApprovals   []interface{} `json:"preDeployApprovals"`
		PostDeployApprovals  []interface{} `json:"postDeployApprovals"`
		PreApprovalsSnapshot struct {
			Approvals []struct {
				Rank             int  `json:"rank"`
				IsAutomated      bool `json:"isAutomated"`
				IsNotificationOn bool `json:"isNotificationOn"`
				Approver         struct {
					ID          string `json:"id"`
					DisplayName string `json:"displayName"`
					UniqueName  string `json:"uniqueName"`
					URL         string `json:"url"`
					ImageURL    string `json:"imageUrl"`
				} `json:"approver"`
				ID int `json:"id"`
			} `json:"approvals"`
			ApprovalOptions struct {
				RequiredApproverCount                                   int  `json:"requiredApproverCount"`
				ReleaseCreatorCanBeApprover                             bool `json:"releaseCreatorCanBeApprover"`
				AutoTriggeredAndPreviousEnvironmentApprovedCanBeSkipped bool `json:"autoTriggeredAndPreviousEnvironmentApprovedCanBeSkipped"`
				EnforceIdentityRevalidation                             bool `json:"enforceIdentityRevalidation"`
				TimeoutInMintues                                        int  `json:"timeoutInMintues"`
			} `json:"approvalOptions"`
		} `json:"preApprovalsSnapshot"`
		PostApprovalsSnapshot struct {
			Approvals []struct {
				Rank             int  `json:"rank"`
				IsAutomated      bool `json:"isAutomated"`
				IsNotificationOn bool `json:"isNotificationOn"`
				ID               int  `json:"id"`
			} `json:"approvals"`
		} `json:"postApprovalsSnapshot"`
		DeploySteps             []interface{} `json:"deploySteps"`
		Rank                    int           `json:"rank"`
		DefinitionEnvironmentID int           `json:"definitionEnvironmentId"`
		EnvironmentOptions      struct {
			EmailNotificationType   string `json:"emailNotificationType"`
			EmailRecipients         string `json:"emailRecipients"`
			SkipArtifactsDownload   bool   `json:"skipArtifactsDownload"`
			TimeoutInMinutes        int    `json:"timeoutInMinutes"`
			EnableAccessToken       bool   `json:"enableAccessToken"`
			PublishDeploymentStatus bool   `json:"publishDeploymentStatus"`
		} `json:"environmentOptions"`
		Demands              []interface{} `json:"demands"`
		Conditions           []interface{} `json:"conditions"`
		WorkflowTasks        []interface{} `json:"workflowTasks"`
		DeployPhasesSnapshot []struct {
			DeploymentInput struct {
				ParallelExecution struct {
					ParallelExecutionType string `json:"parallelExecutionType"`
				} `json:"parallelExecution"`
				SkipArtifactsDownload bool          `json:"skipArtifactsDownload"`
				TimeoutInMinutes      int           `json:"timeoutInMinutes"`
				QueueID               int           `json:"queueId"`
				Demands               []interface{} `json:"demands"`
				EnableAccessToken     bool          `json:"enableAccessToken"`
			} `json:"deploymentInput"`
			Rank          int    `json:"rank"`
			PhaseType     string `json:"phaseType"`
			Name          string `json:"name"`
			WorkflowTasks []struct {
				TaskID           string `json:"taskId"`
				Version          string `json:"version"`
				Name             string `json:"name"`
				Enabled          bool   `json:"enabled"`
				AlwaysRun        bool   `json:"alwaysRun"`
				ContinueOnError  bool   `json:"continueOnError"`
				TimeoutInMinutes int    `json:"timeoutInMinutes"`
				DefinitionType   string `json:"definitionType"`
				Inputs           struct {
					ScriptType          string `json:"scriptType"`
					ScriptName          string `json:"scriptName"`
					Arguments           string `json:"arguments"`
					InlineScript        string `json:"inlineScript"`
					WorkingFolder       string `json:"workingFolder"`
					FailOnStandardError string `json:"failOnStandardError"`
				} `json:"inputs"`
			} `json:"workflowTasks"`
		} `json:"deployPhasesSnapshot"`
		Owner struct {
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
			UniqueName  string `json:"uniqueName"`
			URL         string `json:"url"`
			ImageURL    string `json:"imageUrl"`
		} `json:"owner"`
		Schedules []interface{} `json:"schedules"`
		Release   struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			URL   string `json:"url"`
			Links struct {
				Web struct {
					Href string `json:"href"`
				} `json:"web"`
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
		} `json:"release"`
		ReleaseDefinition struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			URL   string `json:"url"`
			Links struct {
				Web struct {
					Href string `json:"href"`
				} `json:"web"`
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
		} `json:"releaseDefinition"`
		ReleaseCreatedBy struct {
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
		} `json:"releaseCreatedBy"`
		TriggerReason string `json:"triggerReason"`
	} `json:"environments"`
	Variables struct {
	} `json:"variables"`
	VariableGroups []interface{} `json:"variableGroups"`
	Artifacts      []struct {
		SourceID            string `json:"sourceId"`
		Type                string `json:"type"`
		Alias               string `json:"alias"`
		DefinitionReference struct {
			ArtifactSourceDefinitionURL struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"artifactSourceDefinitionUrl"`
			DefaultVersionBranch struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"defaultVersionBranch"`
			DefaultVersionSpecific struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"defaultVersionSpecific"`
			DefaultVersionTags struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"defaultVersionTags"`
			DefaultVersionType struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"defaultVersionType"`
			Definition struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"definition"`
			Project struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"project"`
			Version struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"version"`
			ArtifactSourceVersionURL struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"artifactSourceVersionUrl"`
			Branch struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"branch"`
		} `json:"definitionReference"`
		IsPrimary bool `json:"isPrimary"`
	} `json:"artifacts"`
	ReleaseDefinition struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		URL   string `json:"url"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Web struct {
				Href string `json:"href"`
			} `json:"web"`
		} `json:"_links"`
	} `json:"releaseDefinition"`
	Description                string `json:"description"`
	Reason                     string `json:"reason"`
	ReleaseNameFormat          string `json:"releaseNameFormat"`
	KeepForever                bool   `json:"keepForever"`
	DefinitionSnapshotRevision int    `json:"definitionSnapshotRevision"`
	LogsContainerURL           string `json:"logsContainerUrl"`
	URL                        string `json:"url"`
	Links                      struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Web struct {
			Href string `json:"href"`
		} `json:"web"`
	} `json:"_links"`
	Tags             []interface{} `json:"tags"`
	ProjectReference struct {
		ID   string      `json:"id"`
		Name interface{} `json:"name"`
	} `json:"projectReference"`
	Properties struct {
	} `json:"properties"`
}

// WorkItems struct
type WorkItems struct {
	Count int `json:"count"`
	Value []struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"value"`
}

// WorkItem struct
type WorkItem struct {
	ID     int `json:"id"`
	Rev    int `json:"rev"`
	Fields struct {
		SystemAreaPath                     string    `json:"System.AreaPath"`
		SystemTeamProject                  string    `json:"System.TeamProject"`
		SystemIterationPath                string    `json:"System.IterationPath"`
		SystemWorkItemType                 string    `json:"System.WorkItemType"`
		SystemState                        string    `json:"System.State"`
		SystemReason                       string    `json:"System.Reason"`
		SystemAssignedTo                   string    `json:"System.AssignedTo"`
		SystemCreatedDate                  time.Time `json:"System.CreatedDate"`
		SystemCreatedBy                    string    `json:"System.CreatedBy"`
		SystemChangedDate                  time.Time `json:"System.ChangedDate"`
		SystemChangedBy                    string    `json:"System.ChangedBy"`
		SystemTitle                        string    `json:"System.Title"`
		MicrosoftVSTSCommonStateChangeDate time.Time `json:"Microsoft.VSTS.Common.StateChangeDate"`
		MicrosoftVSTSCommonActivatedDate   time.Time `json:"Microsoft.VSTS.Common.ActivatedDate"`
		MicrosoftVSTSCommonActivatedBy     string    `json:"Microsoft.VSTS.Common.ActivatedBy"`
		MicrosoftVSTSCommonPriority        int       `json:"Microsoft.VSTS.Common.Priority"`
		MicrosoftVSTSCommonValueArea       string    `json:"Microsoft.VSTS.Common.ValueArea"`
		SystemTags                         string    `json:"System.Tags"`
	} `json:"fields"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		WorkItemUpdates struct {
			Href string `json:"href"`
		} `json:"workItemUpdates"`
		WorkItemRevisions struct {
			Href string `json:"href"`
		} `json:"workItemRevisions"`
		WorkItemHistory struct {
			Href string `json:"href"`
		} `json:"workItemHistory"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		WorkItemType struct {
			Href string `json:"href"`
		} `json:"workItemType"`
		Fields struct {
			Href string `json:"href"`
		} `json:"fields"`
	} `json:"_links"`
	URL string `json:"url"`
}
