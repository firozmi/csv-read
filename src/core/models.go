package core

// A ServerStatus is an Response that Dispalys Server Status
type ServerStatus struct {
	Code    int    `json:"statusCode"`
	Message string `json:"message"`
}
