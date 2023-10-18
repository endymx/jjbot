package sumgpt

type User struct {
	Uid      int64
	Username string
}

type API struct {
	Code    int    `json:"code"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
	RawText string `json:"rawText"`
}
