package vocabulary

type Input struct {
	Language   string
	Vocabulary string
	Workspace  string
}

type Output struct {
	Directory string
	Summary   string
}

type Runner interface {
	Run(input Input) (Output, error)
}

type Handler struct{}

func (Handler) Run(_ Input) (Output, error) {
	return Output{}, nil
}
