package figya

type Module interface {
	Run() *Result
	Name() string
}
