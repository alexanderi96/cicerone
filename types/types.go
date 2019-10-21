package types

import (
	"net/http"
	//"github.com/alexanderi96/cicerone/db"
	//"github.com/alexanderi96/cicerone/sessions"
)
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
	Categories 	[]CategoryCount
	Referer    	string
}

//method isCicerone
func (c Context) IsCicerone() bool {
	return c.User.Cicerone
}


//load Context
func LoadContext(r *http.Request) (Context, error) {
	return Context{[]Event{}, User{1, "aless", "ianne", "asd", "asd", "mam", 123, 123, false}, "", "", "", "", []CategoryCount{}, ""}, nil
	//var c.Events := db.LoadEvents()
	//c.user := db.GetUserInfo(sessions.GetCurrentUserName(r))
	/*
	LOADING MOK USERINFO
	*/
	//c.user.Uid, c.user.Name, c.user.Surname, c.user.Username, c.user.Cicerone = 

	
}



//CategoryCount is the struct used to populate the sidebar
//which contains the category name and the count of the tasks
//in each category
type CategoryCount struct {
	Name  string
	Count int
}