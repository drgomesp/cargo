package cargo

// Error represents an error with context information
type Error struct {
	Code int
	Err  error
}

func (e Error) Error() string {
	return e.Err.Error()
}
