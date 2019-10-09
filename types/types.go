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