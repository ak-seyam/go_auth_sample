package signup

type SignUpRequest struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Groups   []string `json:"groups"`
}

type SignUpResponse struct {
	Username string   `json:"username"`
	ID       string   `json:"id"`
	Groups   []string `json:"groups"`
}
