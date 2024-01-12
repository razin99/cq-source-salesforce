package client

type Spec struct {
	Endpoint   string `json:"endpoint"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Token      string `json:"token"`
	ClientID   string `json:"client_id"`
	APIVersion string `json:"api_version,omitempty"`
}
