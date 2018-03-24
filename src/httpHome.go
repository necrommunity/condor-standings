package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sillypears/condor-standings/src/log"
)

//  --------------------------------
//  The landing page for the website
//  --------------------------------

func httpHome(c *gin.Context) {
	// Local variables
	w := c.Writer

	users, usersTotal, err := db.Users.GetUsers()
	if err != nil {
		log.Error("Could not get users: ", err)
	}

	foundTables, err := db.Tables.GetTables()
	if err != nil {
		log.Error("Could not get tables: ", err)
	}

	data := TemplateData{
		Title:        "Home",
		UserAccounts: users,
		UsersTotal:   usersTotal,
		FoundTables:  foundTables,
	}
	httpServeTemplate(w, "home", data)
}
