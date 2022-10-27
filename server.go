package main

import (
	SharedIntrastructure "github.com/marcosperez/remember-go/contexts/shared/infrastructure"
	Users "github.com/marcosperez/remember-go/contexts/users"
)


func main(){

	userApp := Users.NewUserApp()

	_, err := userApp.Start()
	if(err != nil) {
		SharedIntrastructure.Logger.Error(err)
	}

	SharedIntrastructure.Logger.Log("App start succefuly...")

}