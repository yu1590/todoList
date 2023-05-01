package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TodoListModel = (*customTodoListModel)(nil)

type (
	// TodoListModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTodoListModel.
	TodoListModel interface {
		todoListModel
	}

	customTodoListModel struct {
		*defaultTodoListModel
	}
)

// NewTodoListModel returns a model for the database table.
func NewTodoListModel(conn sqlx.SqlConn) TodoListModel {
	return &customTodoListModel{
		defaultTodoListModel: newTodoListModel(conn),
	}
}
