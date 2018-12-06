package freshservice

import "time"

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
