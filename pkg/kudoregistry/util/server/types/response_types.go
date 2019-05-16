package types

const (
	FileNotFound       = "File not found."
	BackendNotFound    = "Storage Backend does not exist."
	YamlContentType    = "application/x-yaml"
	PackageContentType = "application/x-tar"
)

type Response struct {
	Data string `json:"message"`
}

type ResponseError struct {
	Data string `json:"error"`
}
