package repository

type GitLfsHostedRepository struct {
	Name    string        `json:"name"`
	Online  bool          `json:"online"`
	Storage HostedStorage `json:"storage"`

	*Cleanup   `json:"cleanuppolicies,omitempty"`
	*Component `json:"component,omitempty"`
}
