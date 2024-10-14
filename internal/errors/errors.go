package errors

type systemError struct {
}

func (e *systemError) Error() string {
	return "システムエラー"
}

// func MakeSystemError() systemError {
// 	return errors.New()
// }
