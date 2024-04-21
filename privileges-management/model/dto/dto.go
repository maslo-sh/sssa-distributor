package dto

type RequestAccessPayload struct {
	ResourceID string
	Username   string
}

type ApprovalPayload struct {
	Share string
}

type DenialPayload struct {
	Justification string
}

type ApproverAssignmentPayload struct {
	ResourceID string
	Approvers  []string
}
