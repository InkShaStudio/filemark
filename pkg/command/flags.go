package command

type SCommandFlag[T any] struct {
	Name        string `json:"name"`
	Value       *T     `json:"value"`
	Short       string `json:"short"`
	Description string `json:"description"`
}

type ICommandFlag[T any] interface {
	ChangeName(name string) *SCommandFlag[T]
	ChangeValue(value T) *SCommandFlag[T]
	ChangeShort(short string) *SCommandFlag[T]
	ChangeDescription(description string) *SCommandFlag[T]
}

func (flag *SCommandFlag[T]) ChangeName(name string) *SCommandFlag[T] {
	short := string([]rune(name)[0])
	flag.Name = name
	flag.Short = short
	return flag
}

func (flag *SCommandFlag[T]) ChangeValue(value *T) *SCommandFlag[T] {
	flag.Value = value
	return flag
}

func (flag *SCommandFlag[T]) ChangeShort(short string) *SCommandFlag[T] {
	flag.Short = short
	return flag
}

func (flag *SCommandFlag[T]) ChangeDescription(description string) *SCommandFlag[T] {
	flag.Description = description
	return flag
}

func NewCommandFlag[T any](name string) *SCommandFlag[T] {
	short := string([]rune(name)[0])
	var val *T

	return &SCommandFlag[T]{
		Name:        name,
		Value:       val,
		Short:       short,
		Description: "",
	}
}
