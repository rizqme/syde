package model

// LearningCategory represents the type of learning.
type LearningCategory string

const (
	CatGotcha      LearningCategory = "gotcha"
	CatConstraint  LearningCategory = "constraint"
	CatConvention  LearningCategory = "convention"
	CatContext     LearningCategory = "context"
	CatDependency  LearningCategory = "dependency"
	CatPerformance LearningCategory = "performance"
	CatWorkaround  LearningCategory = "workaround"
)

// LearningSource represents how the learning was discovered.
type LearningSource string

const (
	SourceSessionObservation LearningSource = "session-observation"
	SourceDebugging          LearningSource = "debugging"
	SourceCodeReview         LearningSource = "code-review"
	SourceIncident           LearningSource = "incident"
	SourceUserReport         LearningSource = "user-report"
)

// Confidence represents how certain we are about a learning.
type Confidence string

const (
	ConfidenceHigh   Confidence = "high"
	ConfidenceMedium Confidence = "medium"
	ConfidenceLow    Confidence = "low"
)

// LearningEntity represents a captured design learning.
type LearningEntity struct {
	BaseEntity   `yaml:",inline"`
	Category     LearningCategory `yaml:"category,omitempty"`
	EntityRefs   []string         `yaml:"entity_refs,omitempty"`
	Source       LearningSource   `yaml:"source,omitempty"`
	ConfLevel    Confidence       `yaml:"confidence,omitempty"`
	PromotedTo   string           `yaml:"promoted_to,omitempty"`
	DiscoveredAt string           `yaml:"discovered_at,omitempty"`
}
