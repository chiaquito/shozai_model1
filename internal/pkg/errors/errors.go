package errors

import (
	"fmt"
)

const (
	BusinessErrorCode          int = 400
	RecordNotFoundErrorCode    int = 404
	OptimisticLockingErrorCode int = 409
	SystemErrorCode            int = 500
)

type AppError struct {
	Code    int    // エラーコード
	Message string // エラーメッセージ
	Err     error  // 元のエラー
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return fmt.Sprintf("code: %d  message: %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func IsSystemError(err error) bool {
	if e, ok := err.(*AppError); ok {
		return e.Code == SystemErrorCode
	}
	return false
}

func IsBusinessError(err error) bool {
	if e, ok := err.(*AppError); ok {
		return e.Code == SystemErrorCode
	}
	return false
}

func IsOptimisticLockingError(err error) bool {
	if e, ok := err.(*AppError); ok {
		return e.Code == OptimisticLockingErrorCode
	}
	return false
}

func MakeSystemError(err error) *AppError {
	return &AppError{
		Code:    SystemErrorCode,
		Message: "システムエラー 情報システム部に問い合わせてください。",
		Err:     err,
	}
}

func MakeRecordNotFoundError(err error) *AppError {
	return &AppError{
		Code:    RecordNotFoundErrorCode,
		Message: "データがありません。",
		Err:     err,
	}
}

func MakeOptimisticLockingError(id int) *AppError {
	return &AppError{
		Code:    OptimisticLockingErrorCode,
		Message: "他のユーザーが操作している可能性があります。ページを再更新してください。",
		Err:     fmt.Errorf("id: %d", id),
	}
}

func MakeBusinessError(message string, err error) *AppError {
	return &AppError{
		Code:    BusinessErrorCode,
		Message: message,
		Err:     err,
	}
}
