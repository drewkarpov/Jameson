package model

import "github.com/drewkarpov/Jameson/pkg/image"

type Project struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type TestContainer struct {
	ID          string           `json:"id" bson:"id"`
	Name        string           `json:"name" bson:"name"`
	ProjectId   string           `json:"project_id" bson:"project_id"`
	ReferenceId string           `json:"reference_id" bson:"reference_id"`
	Approved    bool             `json:"approved" bson:"approved"`
	VoidZones   []image.VoidZone `json:"void_zones" bson:"void_zones"`
	Tests       []Test           `json:"tests" bson:"tests"`
}

type Test struct {
	ID          string     `json:"id" bson:"id"`
	CandidateId string     `json:"candidate_id" bson:"candidate_id"`
	Result      TestResult `json:"result" bson:"result"`
}

type TestResult struct {
	ID         string  `json:"id" bson:"id"`
	Percentage float64 `json:"percentage" bson:"percentage"`
}

type ResultContainer struct {
	Percentage float64         `json:"percentage"`
	Images     ImagesContainer `json:"images"`
}

type ImagesContainer struct {
	ReferenceId string `json:"reference"`
	CandidateId string `json:"candidate"`
	DiffId      string `json:"diff"`
}

type Reference struct {
	ID string `json:"reference_id"`
}
