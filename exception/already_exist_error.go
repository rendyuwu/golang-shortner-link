package exception

type AlreadyExistError struct {
	Error string
}

func NewAlreadyExistError(error string) AlreadyExistError {
	return AlreadyExistError{
		Error: error,
	}
}
