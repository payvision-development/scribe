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
		DisplayName string `json:"displayName"`
		URL         string `json:"url"`
		Links       struct {
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"_links"`
		ID         string `json:"id"`
		UniqueName string `json:"uniqueName"`
		ImageURL   string `json:"imageUrl"`
	} `json:"modifiedBy"`
	CreatedBy struct {
		DisplayName string `json:"displayName"`
		URL         string `json:"url"`
		Links       struct {
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"_links"`
		ID         string `json:"id"`
		UniqueName string `json:"uniqueName"`
		ImageURL   string `json:"imageUrl"`
	} `json:"createdBy"`
	Environments []struct {
		ID        int    `json:"id"`
		ReleaseID int    `json:"releaseId"`
		Name      string `json:"name"`
		Status    string `json:"status"`
		Variables struct {
		} `json:"variables"`
		VariableGroups     []interface{} `json:"variableGroups"`
		PreDeployApprovals []struct {
			ID       int `json:"id"`
			Revision int `json:"revision"`
			Approver struct {
				DisplayName string `json:"displayName"`
				URL         string `json:"url"`
				Links       struct {
					Avatar struct {
						Href string `json:"href"`
					} `json:"avatar"`
				} `json:"_links"`
				ID         string `json:"id"`
				UniqueName string `json:"uniqueName"`
				ImageURL   string `json:"imageUrl"`
			} `json:"approver"`
			ApprovalType     string    `json:"approvalType"`
			CreatedOn        time.Time `json:"createdOn"`
			ModifiedOn       time.Time `json:"modifiedOn"`
			Status           string    `json:"status"`
			Comments         string    `json:"comments"`
			IsAutomated      bool      `json:"isAutomated"`
			IsNotificationOn bool      `json:"isNotificationOn"`
			TrialNumber      int       `json:"trialNumber"`
			Attempt          int       `json:"attempt"`
			Rank             int       `json:"rank"`
			Release          struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				URL   string `json:"url"`
				Links struct {
				} `json:"_links"`
			} `json:"release"`
			ReleaseDefinition struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				URL   string `json:"url"`
				Links struct {
				} `json:"_links"`
			} `json:"releaseDefinition"`
			ReleaseEnvironment struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Links struct {
				} `json:"_links"`
			} `json:"releaseEnvironment"`
			URL string `json:"url"`
		} `json:"preDeployApprovals"`
		PostDeployApprovals  []interface{} `json:"postDeployApprovals"`
		PreApprovalsSnapshot struct {
			Approvals []struct {
				Rank             int  `json:"rank"`
				IsAutomated      bool `json:"isAutomated"`
				IsNotificationOn bool `json:"isNotificationOn"`
				Approver         struct {
					DisplayName string `json:"displayName"`
					URL         string `json:"url"`
					Links       struct {
						Avatar struct {
							Href string `json:"href"`
						} `json:"avatar"`
					} `json:"_links"`
					ID         string `json:"id"`
					UniqueName string `json:"uniqueName"`
					ImageURL   string `json:"imageUrl"`
				} `json:"approver"`
				ID int `json:"id"`
			} `json:"approvals"`
			ApprovalOptions struct {
				RequiredApproverCount                                   int    `json:"requiredApproverCount"`
				ReleaseCreatorCanBeApprover                             bool   `json:"releaseCreatorCanBeApprover"`
				AutoTriggeredAndPreviousEnvironmentApprovedCanBeSkipped bool   `json:"autoTriggeredAndPreviousEnvironmentApprovedCanBeSkipped"`
				EnforceIdentityRevalidation                             bool   `json:"enforceIdentityRevalidation"`
				TimeoutInMinutes                                        int    `json:"timeoutInMinutes"`
				ExecutionOrder                                          string `json:"executionOrder"`
			} `json:"approvalOptions"`
		} `json:"preApprovalsSnapshot"`
		PostApprovalsSnapshot struct {
			Approvals []struct {
				Rank             int  `json:"rank"`
				IsAutomated      bool `json:"isAutomated"`
				IsNotificationOn bool `json:"isNotificationOn"`
				Approver         struct {
					DisplayName string `json:"displayName"`
					URL         string `json:"url"`
					Links       struct {
						Avatar struct {
							Href string `json:"href"`
						} `json:"avatar"`
					} `json:"_links"`
					ID         string `json:"id"`
					UniqueName string `json:"uniqueName"`
					ImageURL   string `json:"imageUrl"`
				} `json:"approver"`
				ID int `json:"id"`
			} `json:"approvals"`
			ApprovalOptions struct {
				RequiredApproverCount                                   interface{} `json:"requiredApproverCount"`
				ReleaseCreatorCanBeApprover                             bool        `json:"releaseCreatorCanBeApprover"`
				AutoTriggeredAndPreviousEnvironmentApprovedCanBeSkipped bool        `json:"autoTriggeredAndPreviousEnvironmentApprovedCanBeSkipped"`
				EnforceIdentityRevalidation                             bool        `json:"enforceIdentityRevalidation"`
				TimeoutInMinutes                                        int         `json:"timeoutInMinutes"`
				ExecutionOrder                                          string      `json:"executionOrder"`
			} `json:"approvalOptions"`
		} `json:"postApprovalsSnapshot"`
		DeploySteps []struct {
			ID                  int           `json:"id"`
			DeploymentID        int           `json:"deploymentId"`
			Attempt             int           `json:"attempt"`
			Reason              string        `json:"reason"`
			Status              string        `json:"status"`
			OperationStatus     string        `json:"operationStatus"`
			ReleaseDeployPhases []interface{} `json:"releaseDeployPhases"`
			RequestedBy         struct {
				DisplayName string `json:"displayName"`
				URL         string `json:"url"`
				Links       struct {
					Avatar struct {
						Href string `json:"href"`
					} `json:"avatar"`
				} `json:"_links"`
				ID         string `json:"id"`
				UniqueName string `json:"uniqueName"`
				ImageURL   string `json:"imageUrl"`
			} `json:"requestedBy"`
			RequestedFor struct {
				DisplayName string `json:"displayName"`
				URL         string `json:"url"`
				Links       struct {
					Avatar struct {
						Href string `json:"href"`
					} `json:"avatar"`
				} `json:"_links"`
				ID         string `json:"id"`
				UniqueName string `json:"uniqueName"`
				ImageURL   string `json:"imageUrl"`
			} `json:"requestedFor"`
			QueuedOn       time.Time `json:"queuedOn"`
			LastModifiedBy struct {
				DisplayName string `json:"displayName"`
				URL         string `json:"url"`
				Links       struct {
					Avatar struct {
						Href string `json:"href"`
					} `json:"avatar"`
				} `json:"_links"`
				ID         string `json:"id"`
				UniqueName string `json:"uniqueName"`
				ImageURL   string `json:"imageUrl"`
			} `json:"lastModifiedBy"`
			LastModifiedOn time.Time     `json:"lastModifiedOn"`
			HasStarted     bool          `json:"hasStarted"`
			Tasks          []interface{} `json:"tasks"`
			RunPlanID      string        `json:"runPlanId"`
			Issues         []interface{} `json:"issues"`
		} `json:"deploySteps"`
		Rank                    int `json:"rank"`
		DefinitionEnvironmentID int `json:"definitionEnvironmentId"`
		EnvironmentOptions      struct {
			EmailNotificationType   string `json:"emailNotificationType"`
			EmailRecipients         string `json:"emailRecipients"`
			SkipArtifactsDownload   bool   `json:"skipArtifactsDownload"`
			TimeoutInMinutes        int    `json:"timeoutInMinutes"`
			EnableAccessToken       bool   `json:"enableAccessToken"`
			PublishDeploymentStatus bool   `json:"publishDeploymentStatus"`
			BadgeEnabled            bool   `json:"badgeEnabled"`
			AutoLinkWorkItems       bool   `json:"autoLinkWorkItems"`
		} `json:"environmentOptions"`
		Demands    []interface{} `json:"demands"`
		Conditions []struct {
			Result        bool   `json:"result"`
			Name          string `json:"name"`
			ConditionType string `json:"conditionType"`
			Value         string `json:"value"`
		} `json:"conditions"`
		CreatedOn            time.Time     `json:"createdOn"`
		ModifiedOn           time.Time     `json:"modifiedOn"`
		WorkflowTasks        []interface{} `json:"workflowTasks"`
		DeployPhasesSnapshot []struct {
			DeploymentInput struct {
				ParallelExecution struct {
					ParallelExecutionType string `json:"parallelExecutionType"`
				} `json:"parallelExecution"`
				SkipArtifactsDownload  bool `json:"skipArtifactsDownload"`
				ArtifactsDownloadInput struct {
					DownloadInputs []interface{} `json:"downloadInputs"`
				} `json:"artifactsDownloadInput"`
				QueueID                   int           `json:"queueId"`
				Demands                   []interface{} `json:"demands"`
				EnableAccessToken         bool          `json:"enableAccessToken"`
				TimeoutInMinutes          int           `json:"timeoutInMinutes"`
				JobCancelTimeoutInMinutes int           `json:"jobCancelTimeoutInMinutes"`
				Condition                 string        `json:"condition"`
				OverrideInputs            struct {
				} `json:"overrideInputs"`
			} `json:"deploymentInput"`
			Rank          int    `json:"rank"`
			PhaseType     string `json:"phaseType"`
			Name          string `json:"name"`
			WorkflowTasks []struct {
				TaskID           string `json:"taskId"`
				Version          string `json:"version"`
				Name             string `json:"name"`
				RefName          string `json:"refName"`
				Enabled          bool   `json:"enabled"`
				AlwaysRun        bool   `json:"alwaysRun"`
				ContinueOnError  bool   `json:"continueOnError"`
				TimeoutInMinutes int    `json:"timeoutInMinutes"`
				DefinitionType   string `json:"definitionType"`
				OverrideInputs   struct {
				} `json:"overrideInputs"`
				Condition string `json:"condition"`
				Inputs    struct {
					ScriptType          string `json:"scriptType"`
					ScriptName          string `json:"scriptName"`
					Arguments           string `json:"arguments"`
					WorkingFolder       string `json:"workingFolder"`
					InlineScript        string `json:"inlineScript"`
					FailOnStandardError string `json:"failOnStandardError"`
				} `json:"inputs"`
			} `json:"workflowTasks"`
		} `json:"deployPhasesSnapshot"`
		Owner struct {
			DisplayName string `json:"displayName"`
			URL         string `json:"url"`
			Links       struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"_links"`
			ID         string `json:"id"`
			UniqueName string `json:"uniqueName"`
			ImageURL   string `json:"imageUrl"`
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
			DisplayName string `json:"displayName"`
			URL         string `json:"url"`
			Links       struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"_links"`
			ID         string `json:"id"`
			UniqueName string `json:"uniqueName"`
			ImageURL   string `json:"imageUrl"`
		} `json:"releaseCreatedBy"`
		TriggerReason     string `json:"triggerReason"`
		ProcessParameters struct {
		} `json:"processParameters"`
		PreDeploymentGatesSnapshot struct {
			ID           int           `json:"id"`
			GatesOptions interface{}   `json:"gatesOptions"`
			Gates        []interface{} `json:"gates"`
		} `json:"preDeploymentGatesSnapshot"`
		PostDeploymentGatesSnapshot struct {
			ID           int           `json:"id"`
			GatesOptions interface{}   `json:"gatesOptions"`
			Gates        []interface{} `json:"gates"`
		} `json:"postDeploymentGatesSnapshot"`
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
			DefaultVersionType struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"defaultVersionType"`
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
	Tags                    []interface{} `json:"tags"`
	TriggeringArtifactAlias interface{}   `json:"triggeringArtifactAlias"`
	ProjectReference        struct {
		ID   string      `json:"id"`
		Name interface{} `json:"name"`
	} `json:"projectReference"`
	Properties struct {
		DownloadBuildArtifactsUsingTask struct {
			Type  string `json:"$type"`
			Value string `json:"$value"`
		} `json:"DownloadBuildArtifactsUsingTask"`
		ReleaseCreationSource struct {
			Type  string `json:"$type"`
			Value string `json:"$value"`
		} `json:"ReleaseCreationSource"`
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
		SystemAreaPath                       string    `json:"System.AreaPath"`
		SystemTeamProject                    string    `json:"System.TeamProject"`
		SystemIterationPath                  string    `json:"System.IterationPath"`
		SystemWorkItemType                   string    `json:"System.WorkItemType"`
		SystemState                          string    `json:"System.State"`
		SystemReason                         string    `json:"System.Reason"`
		SystemAssignedTo                     string    `json:"System.AssignedTo"`
		SystemCreatedDate                    time.Time `json:"System.CreatedDate"`
		SystemCreatedBy                      string    `json:"System.CreatedBy"`
		SystemChangedDate                    time.Time `json:"System.ChangedDate"`
		SystemChangedBy                      string    `json:"System.ChangedBy"`
		SystemTitle                          string    `json:"System.Title"`
		MicrosoftVSTSCommonStateChangeDate   time.Time `json:"Microsoft.VSTS.Common.StateChangeDate"`
		MicrosoftVSTSCommonActivatedDate     time.Time `json:"Microsoft.VSTS.Common.ActivatedDate"`
		MicrosoftVSTSCommonActivatedBy       string    `json:"Microsoft.VSTS.Common.ActivatedBy"`
		MicrosoftVSTSCommonClosedDate        time.Time `json:"Microsoft.VSTS.Common.ClosedDate"`
		MicrosoftVSTSCommonClosedBy          string    `json:"Microsoft.VSTS.Common.ClosedBy"`
		MicrosoftVSTSCommonPriority          int       `json:"Microsoft.VSTS.Common.Priority"`
		MicrosoftVSTSSchedulingCompletedWork float32   `json:"Microsoft.VSTS.Scheduling.CompletedWork"`
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
