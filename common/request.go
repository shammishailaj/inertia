package common

const (
	// MsgDaemonOK is the OK response upon successfully reaching daemon
	MsgDaemonOK = "I'm a little Webhook, short and stout!"

	// Container is a constant used in HTTP GET query strings
	Container = "container"

	// Stream is a constant used in HTTP GET query strings
	Stream = "stream"

	// DaemonProjectPath is the path the daemon uses for the deployed project
	DaemonProjectPath = "/app/host/project"
)

// DaemonRequest is the configurable body of a request to the daemon.
type DaemonRequest struct {
	Stream     bool        `json:"stream"`
	Container  string      `json:"container,omitempty"`
	Project    string      `json:"project"`
	BuildType  string      `json:"build_type"`
	GitOptions *GitOptions `json:"git_options"`
	Secret     string      `json:"secret"`
}

// GitOptions represents GitHub-related deployment options
type GitOptions struct {
	RemoteURL string `json:"remote"`
	Branch    string `json:"branch"`
}

// UserRequest is used for logging in or modifying users
type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Admin    bool   `json:"admin"`
}

// EnvRequest represents a request to manage environment variables
type EnvRequest struct {
	Name    string `json:"name,omitempty"`
	Value   string `json:"value,omitempty"`
	Encrypt bool   `json:"encrypt,omitempty"`

	Remove bool `json:"remove,omitempty"`
}
