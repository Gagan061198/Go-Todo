package model

/*type Task struct{
	Id int 'json:"id"'
	Desc string 'json:"desc"'
	TimeStamp time.Time 'json:"timestamp"'
	Status string 'json:"status"'
}

type user struct{
	userId int 'json:"userID"'
	Name *name 'json:"name"'
	DOB *DOB 'json:"DOB"'
	Email string 'json:"email"'
}

type name struct{
	FName string 'json:"FName"'
	LName string 'json:"LName"'

}

type DOB struct{
	Date int32 'json:"Date"'
	Month int32 'json:"Month"'
	Year int32 'json:"Year"'

}*/

type TodoItemModel struct {
	Id          int `gorm:"primary_key"`
	Description string
	Completed   bool
}
