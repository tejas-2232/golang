package types

type Student struct {
	Id    int
	Name  string `validate:"required"`
	Email string `validate:"required"`
	Age   int    `json:"age" validate:"required"`
}
