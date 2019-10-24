package daemon

// Daemon is used to group together data related to gourdd
type Daemon struct {
	ID string
}

func init() {
	MakeSocket()
}

// GetSocketPath returns the filesystem path to the command socket
func GetSocketPath() string {
	return "/run/gourd/gourdd.sock"
}
