package models

import "time"

// User — учётная запись пользователя.
type User struct {
	// Уникальный идентификатор пользователя.
	ID int `json:"id" example:"1"`
	// Адрес электронной почты.
	Email string `json:"email" example:"anna@example.com"`
	// Отображаемое имя, до 100 символов.
	Name string `json:"name" example:"Анна"`
	// Время создания учётной записи в формате RFC3339.
	CreatedAt time.Time `json:"created_at" example:"2026-11-03T08:00:00Z"`
}

// CreateUserRequest — тело запроса на создание пользователя.
type CreateUserRequest struct {
	// Email. Обязательное поле, должно быть корректным адресом.
	Email string `json:"email" validate:"required,email" example:"anna@example.com"`
	// Имя пользователя. Обязательное поле, не длиннее 100 символов.
	Name string `json:"name" validate:"required,max=100" example:"Анна"`
}

// UpdateUserRequest — частичное обновление пользователя.
// Передаются только те поля, которые нужно изменить.
type UpdateUserRequest struct {
	// Новый email. Необязательное поле.
	Email *string `json:"email" validate:"omitempty,email" example:"anna.new@example.com"`
	// Новое имя. Необязательное поле, не длиннее 100 символов.
	Name *string `json:"name" validate:"omitempty,max=100" example:"Анна Иванова"`
}

// ErrorResponse — тело ответа с описанием ошибки.
type ErrorResponse struct {
	// Машиночитаемый код ошибки.
	Error string `json:"error" example:"not_found"`
	// Сообщение, пригодное для показа пользователю.
	Message string `json:"message" example:"Задача не найдена"`
}

// SuccessResponse — общая обёртка успешного ответа.
// Конкретный тип поля data зависит от операции: объект User или Todo либо массив таких объектов.
// Если список пуст, data равен null.
type SuccessResponse struct {
	// Признак успешного выполнения запроса.
	Success bool `json:"success" example:"true"`
	// Полезная нагрузка. Отсутствует, если данных нет.
	Data interface{} `json:"data,omitempty"`
}
