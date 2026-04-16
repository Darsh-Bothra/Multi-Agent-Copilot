package response

type Success struct {
	Data interface{} `json:"data"`
}

type Error struct {
	Error string `json:"error"`
}