package pkg

type Project struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type TestContainer struct {
	ID          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	ProjectId   string `json:"project_id" bson:"project_id"`
	ReferenceId string `json:"reference_id" bson:"reference_id"`
	Approved    bool   `json:"approved" bson:"approved"`
	Tests       []Test `json:"tests" bson:"tests"`
}
type TestContainerDTO struct {
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
