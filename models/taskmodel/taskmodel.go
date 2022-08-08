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
			&data.Deadline,
			&data.Status)
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

func (m *TaskModel) Find(id int64, task *entites.Task) error {
	return m.db.QueryRow("select * from task where id_task = ?", id).Scan(
		&task.Id_task,
		&task.Assignee,
		&task.Deadline,
		&task.Status)
}

func (m *TaskModel) Update(task entites.Task) error {

	_, err := m.db.Exec("update task set assignee = ?, deadline = ? where id_task = ?",
		task.Assignee, task.Deadline, task.Id_task)

	if err != nil {
		return err
	}

	return nil
}

func (m *TaskModel) Complete(id int64) error {
	_, err := m.db.Exec("update task set status = 1 where id_task = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (m *TaskModel) Delete(id int64) error {
	_, err := m.db.Exec("delete from task where id_task = ?", id)
	if err != nil {
		return err
	}
	return nil
}
