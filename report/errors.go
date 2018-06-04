package report

type ErrorReportNotFound struct {
	message string
}

func (err ErrorReportNotFound) Error() string {
	return err.message
}

func NewErrorReportNotFound(message string) ErrorReportNotFound {
	return ErrorReportNotFound{message: message}
}

type ErrorInvalidReportData struct {
	message string
}

func (err ErrorInvalidReportData) Error() string {
	return err.message
}

func NewErrorInvalidReportData(message string) ErrorInvalidReportData {
	return ErrorInvalidReportData{message: message}
}

type ErrorReportConflict struct {
	message string
}

func NewErrorConflictReportData(message string) ErrorReportConflict {
	return ErrorReportConflict{
		message: message,
	}
}
