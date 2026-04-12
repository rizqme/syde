package model

// RelationType constants for the relationship model.
const (
	RelBelongsTo  = "belongs_to"
	RelDependsOn  = "depends_on"
	RelExposes    = "exposes"
	RelConsumes   = "consumes"
	RelUses       = "uses"
	RelInvolves   = "involves"
	RelReferences = "references"
	RelAppliesTo  = "applies_to"
	RelModifies   = "modifies"
	RelVisualizes = "visualizes"
)

// ValidRelationshipTypes returns all valid relationship type strings.
func ValidRelationshipTypes() []string {
	return []string{
		RelBelongsTo, RelDependsOn, RelExposes, RelConsumes,
		RelUses, RelInvolves, RelReferences, RelAppliesTo,
		RelModifies, RelVisualizes,
	}
}
