package service

import mdl "github.com/drewkarpov/Jameson/pkg/model"

//go:generate mockgen -source=image_service.go -destination=mocks/mock.go

type ImageService interface {
	CreateProject(project *mdl.Project) (*mdl.Project, error)
	GetProjects() ([]mdl.Project, error)
	GetContainers() []mdl.TestContainer
	GetContainerByName(name string) (*mdl.TestContainer, bool)
	GetContainerById(containerId string) (*mdl.TestContainer, bool)
	ApproveReferenceForContainer(containerId string) (bool, error)
	WritingTestResultToContainer(candidate, result []byte, percentage float64, containerId string) (*mdl.TestResult, error)
	CreateNewTestContainer(testContainer mdl.TestContainer) (*mdl.TestContainer, error)
	SetNewReferenceForContainer(containerId string, reference mdl.Reference) (bool, error)
	DeleteContainerById(containerId string) (bool, error)
	UploadImage(data []byte) (*string, error)
	DownloadImage(fileName string) ([]byte, error)
	GetContainerByTestId(testId string) (*mdl.TestContainer, bool)
}
