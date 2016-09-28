package definition

// Interface providing the basic API for a definition
type Interface interface {
	AddArguments(arg ...Reference) *Definition
	AddMethodCall(method Method) *Definition
}
