package definition

// Interface providing the basic API for a definition
type Interface interface {
	AddArgument(arg Reference) (def *Definition, err error)
}
