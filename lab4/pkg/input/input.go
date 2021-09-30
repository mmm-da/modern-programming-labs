package input

type Input interface {
	GetText() string
	GetTitle() string
	GetAuthors() []string
	GetHash() [16]byte
	ReadFile(string)
}
