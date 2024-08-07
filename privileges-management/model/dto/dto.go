package dto

type RequestAccessPayload struct {
	ResourceID      int
	Username        string
	ValidityInHours int
}

type ApprovalPayload struct {
	Share string
}

type DenialPayload struct {
	Justification string
}

type ApproverAssignmentPayload struct {
	ResourceID int
	Approvers  []string
}

type AuthPayload struct {
	Username string
	Password string
}

type ResourceRegistrationPayload struct {
	ResourceDN        string
	SharesCreated     int
	MinSharesRequired int
}

type ApproverRegistrationPayload struct {
	Username   string
	ResourceID int
}
