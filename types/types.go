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
	Events      []Event 	//using
	User 		 		//using
	Navigation 	string
	Search     	string
	Message    	string
	CSRFToken  	string
	Referer    	string
}

//method isCicerone
func (c Context) IsCicerone() bool {
	return c.User.Cicerone
}
