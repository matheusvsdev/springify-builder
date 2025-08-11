package model

type ComposeData struct {
	ServiceName   string
	ImageName     string
	ContainerName string
	DbName        string
	Port          string
	InternalPort  string
	VolumePath    string
	NetworkName   string
	EnvVars       map[string]string
}
