package command

type SCommandArg[T any] struct {
	Index       int    `json:"index"`
	Name        string `json:"name"`
	Value       T      `json:"value"`
	Description string `json:"description"`
}

type ICommandArg[T any] interface {
	ChangeIndex(index int) *SCommandArg[T]
	ChangeName(name string) *SCommandArg[T]
	ChangeValue(value T) *SCommandArg[T]
	ChangeDescription(description string) *SCommandArg[T]
}

func (arg *SCommandArg[T]) ChangeIndex(index int) *SCommandArg[T] {
	arg.Index = index
	return arg
}

func (arg *SCommandArg[T]) ChangeName(name string) *SCommandArg[T] {
	arg.Name = name
	return arg
}

func (arg *SCommandArg[T]) ChangeValue(value T) *SCommandArg[T] {
	arg.Value = value
	return arg
}

func (arg *SCommandArg[T]) ChangeDescription(description string) *SCommandArg[T] {
	arg.Description = description
	return arg
}

func NewCommandArg[T any](name string) *SCommandArg[T] {
	var val T

	return &SCommandArg[T]{
		Index:       0,
		Name:        name,
		Value:       val,
		Description: "",
	}
}
