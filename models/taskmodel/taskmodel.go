package taskmodel

import (
	"database/sql"

	"github.com/primaaku/project-go-prima/config"
	"github.com/primaaku/project-go-prima/entites"
)

type TaskModel struct {
	db *sql.DB
}

func New() *TaskModel {
	db, err := config.DBConnection()
	if err != nil {
		panic(err)
	}
	return &TaskModel{db: db}
}

func (m *TaskModel) FindAll(task *[]entites.Task) error {
	rows, err := m.db.Query("select * from task")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var data entites.Task
		rows.Scan(
			&data.Id_task,
			&data.Assignee,
			&data.Deadline)
		*task = append(*task, data)
	}

	return nil
}

func (m *TaskModel) Create(task *entites.Task) error {
	result, err := m.db.Exec("insert into task (assignee, deadline) values(?,?)",
		task.Assignee, task.Deadline)

	if err != nil {
		return err
	}

	lastInsertId, _ := result.LastInsertId()
	task.Id_task = lastInsertId
	return nil
}
