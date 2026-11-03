package models

import "time"

// Todo — задача, принадлежащая пользователю.
type Todo struct {
	// Уникальный идентификатор задачи.
	ID int `json:"id" example:"1"`
	// Идентификатор владельца задачи.
	UserID int `json:"user_id" example:"1"`
	// Краткий заголовок, до 200 символов.
	Title string `json:"title" example:"Купить молоко"`
	// Подробное описание, до 1000 символов.
	Description string `json:"description" example:"2 литра, 3.2%"`
	// Признак выполнения.
	Done bool `json:"done" example:"false"`
	// Время создания в формате RFC3339.
	CreatedAt time.Time `json:"created_at" example:"2026-11-03T08:00:00Z"`
	// Время последнего изменения в формате RFC3339.
	UpdatedAt time.Time `json:"updated_at" example:"2026-11-03T08:00:00Z"`
}

// CreateTodoRequest — тело запроса на создание задачи.
type CreateTodoRequest struct {
	// Идентификатор пользователя, которому принадлежит задача.
	UserID int `json:"user_id" validate:"required" example:"1"`
	// Заголовок задачи. Обязательное поле, не длиннее 200 символов.
	Title string `json:"title" validate:"required,max=200" example:"Купить молоко"`
	// Описание задачи, не длиннее 1000 символов.
	Description string `json:"description" validate:"max=1000" example:"2 литра, 3.2%"`
	// Начальный статус выполнения.
	Done bool `json:"done" example:"false"`
}

// UpdateTodoRequest — частичное обновление задачи.
// Передаются только те поля, которые нужно изменить.
type UpdateTodoRequest struct {
	// Новый заголовок. Необязательное поле, не длиннее 200 символов.
	Title *string `json:"title" validate:"omitempty,max=200" example:"Купить хлеб"`
	// Новое описание. Необязательное поле, не длиннее 1000 символов.
	Description *string `json:"description" validate:"omitempty,max=1000" example:"Бородинский"`
	// Новый статус выполнения.
	Done *bool `json:"done" example:"true"`
}
