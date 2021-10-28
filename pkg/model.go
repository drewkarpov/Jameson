package pkg

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	ID             primitive.ObjectID `json:"id" bson:"id"`
	Name           string             `json:"name" bson:"name"`
	TestContainers []TestContainer    `json:"containers" bson:"containers"`
}

type TestContainer struct {
	ID          primitive.ObjectID `json:"id" bson:"id"`
	Name        string             `json:"name" bson:"name"`
	ReferenceId string             `json:"reference_id" bson:"reference_id"`
	Approved    bool               `json:"approved" bson:"approved"`
	Tests       []Test             `json:"tests" bson:"tests"`
}
type TestContainerDTO struct {
	ProjectId   string `json:"project_id" bson:"project_id"`
	Name        string `json:"name" bson:"name"`
	ReferenceId string `json:"reference_id" bson:"reference_id"`
	Approved    bool   `json:"approved" bson:"approved"`
}

type Test struct {
	CandidateId string     `json:"candidate_id" bson:"candidate_id"`
	Result      TestResult `json:"result" bson:"result"`
}

type TestResult struct {
	ResultId   string  `json:"result_id" bson:"result_id"`
	Percentage float64 `json:"percentage" bson:"percentage"`
}
