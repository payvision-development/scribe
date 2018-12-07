package freshservice

import "time"

// RequestItilChange type
type RequestItilChange struct {
	ItilChange struct {
		Email            string `json:"email"`
		Subject          string `json:"subject"`
		DescriptionHTML  string `json:"description_html"`
		Status           int    `json:"status"`
		Priority         int    `json:"priority"`
		ChangeType       int    `json:"change_type"`
		Risk             int    `json:"risk"`
		Impact           int    `json:"impact"`
		PlannedStartDate string `json:"planned_start_date,omitempty"`
		PlannedEndDate   string `json:"planned_end_date,omitempty"`
	} `json:"itil_change"`
}

// ResponseItilChange type
type ResponseItilChange struct {
	Status bool `json:"status"`
	Item   struct {
		ItilChange struct {
			ID             int64       `json:"id"`
			DisplayID      int64       `json:"display_id"`
			RequesterID    int64       `json:"requester_id"`
			OwnerID        interface{} `json:"owner_id"`
			GroupID        interface{} `json:"group_id"`
			Priority       int         `json:"priority"`
			Impact         int         `json:"impact"`
			Status         int         `json:"status"`
			Risk           int         `json:"risk"`
			ChangeType     int         `json:"change_type"`
			ApprovalStatus int         `json:"approval_status"`
			Deleted        bool        `json:"deleted"`
			Subject        string      `json:"subject"`
			CreatedAt      time.Time   `json:"created_at"`
			UpdatedAt      time.Time   `json:"updated_at"`
			CcEmail        struct {
			} `json:"cc_email"`
			PlannedStartDate   time.Time     `json:"planned_start_date"`
			PlannedEndDate     time.Time     `json:"planned_end_date"`
			ImportID           interface{}   `json:"import_id"`
			DepartmentID       interface{}   `json:"department_id"`
			EmailConfigID      interface{}   `json:"email_config_id"`
			ProjectID          interface{}   `json:"project_id"`
			ApprovalType       interface{}   `json:"approval_type"`
			WfEventID          interface{}   `json:"wf_event_id"`
			StateFlowID        interface{}   `json:"state_flow_id"`
			StateTraversal     []int         `json:"state_traversal"`
			StatusName         string        `json:"status_name"`
			ImpactName         string        `json:"impact_name"`
			PriorityName       string        `json:"priority_name"`
			RequesterName      string        `json:"requester_name"`
			OwnerName          interface{}   `json:"owner_name"`
			GroupName          interface{}   `json:"group_name"`
			RiskType           string        `json:"risk_type"`
			ChangeTypeName     string        `json:"change_type_name"`
			ApprovalStatusName string        `json:"approval_status_name"`
			Description        string        `json:"description"`
			AssocReleaseID     interface{}   `json:"assoc_release_id"`
			AssociatedAssets   []interface{} `json:"associated_assets"`
			Attachments        []interface{} `json:"attachments"`
			Notes              []struct {
				Note struct {
					ID         int64       `json:"id"`
					Body       string      `json:"body"`
					BodyHTML   string      `json:"body_html"`
					ToEmails   interface{} `json:"to_emails"`
					CcEmails   interface{} `json:"cc_emails"`
					Deleted    bool        `json:"deleted"`
					UserID     int64       `json:"user_id"`
					CreatedAt  time.Time   `json:"created_at"`
					UpdatedAt  time.Time   `json:"updated_at"`
					HeaderInfo interface{} `json:"header_info"`
				} `json:"note"`
			} `json:"notes"`
			CustomFieldValues struct {
			} `json:"custom_field_values"`
		} `json:"itil_change"`
	} `json:"item"`
	Redirect interface{} `json:"redirect"`
}

// RequestNote type
type RequestNote struct {
	Note struct {
		BodyHTML string `json:"body_html"`
	} `json:"itil_note"`
}

// ResponseNote type
type ResponseNote struct {
	Status bool `json:"status"`
	Item   struct {
		Note struct {
			ID          int64       `json:"id"`
			Body        string      `json:"body"`
			BodyHTML    string      `json:"body_html"`
			ToEmails    interface{} `json:"to_emails"`
			CcEmails    interface{} `json:"cc_emails"`
			Deleted     bool        `json:"deleted"`
			NotableType string      `json:"notable_type"`
			NotableID   int64       `json:"notable_id"`
			UserID      int64       `json:"user_id"`
			AccountID   int         `json:"account_id"`
			CreatedAt   time.Time   `json:"created_at"`
			UpdatedAt   time.Time   `json:"updated_at"`
			HeaderInfo  interface{} `json:"header_info"`
		} `json:"note"`
	} `json:"item"`
	Redirect interface{} `json:"redirect"`
}
