package transformer

type RegisterTransformer struct {
	Token string `json:"token"`
}

type LoginTransformer struct {
	Token string `json:"token"`
}

type MeTransformer struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
