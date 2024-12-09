package structures

type Todo struct {
	Id int `json:"id" bson: "_id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}