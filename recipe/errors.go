package recipe

type ErrorRecipeNotFound struct {
	message string
}

func (err ErrorRecipeNotFound) Error() string {
	return err.message
}

func NewErrorRecipeNotFound(message string) ErrorRecipeNotFound {
	return ErrorRecipeNotFound{message: message}
}

type ErrorInvalidRecipeData struct {
	message string
}

func (err ErrorInvalidRecipeData) Error() string {
	return err.message
}

func NewErrorInvalidRecipeData(message string) ErrorInvalidRecipeData {
	return ErrorInvalidRecipeData{message: message}
}

type ErrorRecipeConflict struct {
	message string
}

func NewErrorConflictUser(message string) ErrorRecipeConflict {
	return ErrorRecipeConflict{
		message: message,
	}
}
