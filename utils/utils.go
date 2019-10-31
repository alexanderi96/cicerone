package utils

import "strings"

func GetRedirectUrl(referer string) string {
	var redirectUrl string
	url := strings.Split(referer, "/")

	if len(url) > 4 {
		redirectUrl = "/" + strings.Join(url[3:], "/")
	} else {
		redirectUrl = "/"
	}
	return redirectUrl
}


//maybe one day will implement epoch data storage
func DateToUnix(date string) int64 {

}

func UnixToDate(unix int64) string {
	
} 