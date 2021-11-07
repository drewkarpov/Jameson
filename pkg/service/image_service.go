package service

import mdl "Jameson/pkg/model"

//go:generate mockgen -source=image_service.go -destination=mocks/mock.go

type ImageService interface {
	CreateProject(project mdl.Project) (*mdl.Project, error)
	GetProjects() []mdl.Project
	GetContainers() []mdl.TestContainer
	GetContainerByName(name string) (*mdl.TestContainer, bool)
	GetContainerById(containerId string) (*mdl.TestContainer, bool)
	ApproveReferenceForContainer(containerId string) (bool, error)
	WritingTestResultToContainer(containerId string, test mdl.Test) (bool, error)
	CreateNewTestContainer(testContainer mdl.TestContainer) (*mdl.TestContainer, error)
	SetNewReferenceForContainer(containerId, referenceId string) (bool, error)
	DeleteContainerById(containerId string) (bool, error)
	UploadImage(data []byte) (*string, error)
	DownloadImage(fileName string) ([]byte, error)
}
