package entity

type Response[T any] struct {
	Msg  string `json:"msg"`
	Code string `json:"code"`
	Data T      `json:"data,omitempty"`
}
