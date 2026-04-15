package model

// RelationType constants for the relationship model.
const (
	RelBelongsTo   = "belongs_to"
	RelDependsOn   = "depends_on"
	RelExposes     = "exposes"
	RelConsumes    = "consumes"
	RelUses        = "uses"
	RelInvolves    = "involves"
	RelReferences  = "references"
	RelRelatesTo   = "relates_to"
	RelAppliesTo   = "applies_to"
	RelModifies    = "modifies"
	RelVisualizes  = "visualizes"
	RelImplements  = "implements"
	RelRefines     = "refines"      // requirement → requirement (more specific) or requirement → component/contract/concept/system (the entity the requirement refines)
	RelDerivesFrom = "derives_from" // requirement → requirement (parent the child was derived from)
)

// ValidRelationshipTypes returns all valid relationship type strings.
func ValidRelationshipTypes() []string {
	return []string{
		RelBelongsTo, RelDependsOn, RelExposes, RelConsumes,
		RelUses, RelInvolves, RelReferences, RelRelatesTo,
		RelAppliesTo, RelModifies, RelVisualizes, RelImplements,
		RelRefines, RelDerivesFrom,
	}
}
