package dto

type ArithmeticDto struct {
	First  float64 `json:"first" reqHeader:"first"`
	Second float64 `json:"second" reqHeader:"second"`
}
