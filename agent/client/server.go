package client

// The gRPC server to retrieve tasks from, and send results to.
type Server struct {
	ip   string
	port uint16
}
