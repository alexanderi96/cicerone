package types

/*
Package types is used to store the context struct which
is passed while templates are executed.

MUST EDIT
*/

//must add json things
type User struct {
	Uid			int 	`json:"id"`
    Name 		string
    Surname 	string
    Username 	string
    Email 		string
    Passwd_h 	string
    Reg_date 	int
    Last_login 	int 
    Cicerone 	bool
}

type Event struct {

}

//Context is the struct passed to templates
type Context struct {
	Events      []Event
	Navigation string
	Search     string
	Message    string
	CSRFToken  string
	Categories []CategoryCount
	Referer    string
}

//CategoryCount is the struct used to populate the sidebar
//which contains the category name and the count of the tasks
//in each category
type CategoryCount struct {
	Name  string
	Count int
}