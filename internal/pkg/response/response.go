package response

type Position struct {
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Message string  `json:"message"`
	Data    data
}

type data struct {
	Result string
}
