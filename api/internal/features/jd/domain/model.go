package domain

import (
	"time"

	"github.com/google/uuid"
)

// Seniority is optional: nil means the JD supplies insufficient evidence.
type Seniority string

const (
	Intern Seniority = "intern"
	Junior Seniority = "junior"
	Mid    Seniority = "mid"
	Senior Seniority = "senior"
	Lead   Seniority = "lead"
)

type SkillCategory string

const (
	ProgrammingLanguage SkillCategory = "programming_language"
	Framework           SkillCategory = "framework"
	Database            SkillCategory = "database"
	Tool                SkillCategory = "tool"
	Technology          SkillCategory = "technology"
	Other               SkillCategory = "other"
)

type Requirement string

const (
	Required  Requirement = "required"
	Preferred Requirement = "preferred"
)

type ExtractedSkill struct {
	Name     string        `json:"name"`
	Category SkillCategory `json:"category"`
	// Empty means the JD does not state whether the skill is required or preferred.
	Requirement Requirement `json:"requirement,omitempty"`
}

// ExtractionCandidate is an untrusted proposal, including when decoded successfully.
// Nil collections indicate missing/null structural data; empty arrays are valid.
type ExtractionCandidate struct {
	Title           string           `json:"title"`
	SeniorityLevel  *Seniority       `json:"seniorityLevel,omitempty"`
	Skills          []ExtractedSkill `json:"skills"`
	Technologies    []string         `json:"technologies"`
	DomainKnowledge []string         `json:"domainKnowledge"`
}

// StructuredJD is returned only after deterministic domain validation.
type StructuredJD struct {
	Title           string           `json:"title"`
	SeniorityLevel  *Seniority       `json:"seniorityLevel,omitempty"`
	Skills          []ExtractedSkill `json:"skills"`
	Technologies    []string         `json:"technologies"`
	DomainKnowledge []string         `json:"domainKnowledge"`
}

// Status is the persisted lifecycle state of a job description.
type Status string

const (
	StatusUploaded   Status = "uploaded"
	StatusParsing    Status = "parsing"
	StatusParsed     Status = "parsed"
	StatusCustomized Status = "customized"
	StatusLocked     Status = "locked"
	StatusFailed     Status = "failed"
)

// JD is the repository-facing representation of a persisted job description.
type JD struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	RawText      string
	StructuredJD StructuredJD
	Status       Status
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// ListItem contains the summary fields of a persisted job description.
type ListItem struct {
	ID             uuid.UUID `json:"id"`
	Title          string    `json:"title"`
	SeniorityLevel Seniority `json:"seniorityLevel"`
	Status         Status    `json:"status"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// CreateInput contains the original JD text and its reviewed structure.
type CreateInput struct {
	RawText      string
	StructuredJD StructuredJD
}

// UpdateInput contains only fields which may change after JD creation.
type UpdateInput struct {
	StructuredJD StructuredJD
}
