package service

import mdl "Jameson/pkg/model"

type ImageService interface {
	CreateProject(project mdl.Project) (*mdl.Project, error)
	GetProjects() []mdl.Project
	GetContainers() []mdl.TestContainer
	GetContainerByName(name string) (*mdl.TestContainer, bool)
	GetContainerById(containerId string) (*mdl.TestContainer, bool)
	ApproveReferenceForContainer(containerId string) (bool, error)
	WritingTestResultToContainer(containerId string, test mdl.Test) (bool, error)
	CreateNewTestContainer(testContainer mdl.TestContainer) (*mdl.TestContainer, error)
	UploadImage(data []byte) (*string, error)
	DownloadImage(fileName string) ([]byte, error)
}
