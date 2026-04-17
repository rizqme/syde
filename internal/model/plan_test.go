package model

import (
	"reflect"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestPlanChangeTasksRoundTripYAML(t *testing.T) {
	in := PlanEntity{
		Changes: PlanChanges{
			Components: ChangeLane{
				Deleted: []DeletedChange{{
					ID:    "del1",
					Slug:  "old-component",
					Why:   "no longer used",
					Tasks: []string{"task-a", "task-b"},
				}},
				Extended: []ExtendedChange{{
					ID:    "ext1",
					Slug:  "existing-component",
					What:  "add behavior",
					Why:   "needed by plan",
					Tasks: []string{"task-c"},
				}},
				New: []NewChange{{
					ID:    "new1",
					Name:  "New Component",
					What:  "create it",
					Why:   "needed by plan",
					Draft: map[string]interface{}{"responsibility": "do work"},
					Tasks: []string{"task-d"},
				}},
			},
		},
	}

	data, err := yaml.Marshal(in)
	if err != nil {
		t.Fatalf("marshal plan: %v", err)
	}

	var out PlanEntity
	if err := yaml.Unmarshal(data, &out); err != nil {
		t.Fatalf("unmarshal plan: %v", err)
	}

	gotLane := out.Changes.Components
	if !reflect.DeepEqual(gotLane.Deleted[0].Tasks, []string{"task-a", "task-b"}) {
		t.Fatalf("deleted tasks did not round trip: %#v", gotLane.Deleted[0].Tasks)
	}
	if !reflect.DeepEqual(gotLane.Extended[0].Tasks, []string{"task-c"}) {
		t.Fatalf("extended tasks did not round trip: %#v", gotLane.Extended[0].Tasks)
	}
	if !reflect.DeepEqual(gotLane.New[0].Tasks, []string{"task-d"}) {
		t.Fatalf("new tasks did not round trip: %#v", gotLane.New[0].Tasks)
	}
}
