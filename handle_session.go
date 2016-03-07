package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleSessionDestroy(c *gin.Context) {
	fmt.Println("Attempting to sign out...")
	session := RequestSession(c.Request)
	if session != nil {
		err := globalSessionStore.Delete(session)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("signing out...")
	c.HTML(http.StatusOK, "index.tmpl.html", nil)
}

func HandleSessionNew(c *gin.Context) {
	var form Register

	if c.Bind(&form) == nil {
		if form.RetypePassword == form.Password {

			c.HTML(http.StatusOK, "view.tmpl.html", nil)
			fmt.Println("it worked!")
		} else {
			fmt.Println("email worked: ", form.Email == "user")
			fmt.Println("password worked: ", form.RetypePassword == form.Password, form.RetypePassword, form.Password)
			fmt.Println("You didn't get logged in")
			c.HTML(http.StatusUnauthorized, "register.tmpl.html", nil)
		}
	}
}

func HandleSessionCreate(c *gin.Context) {
	fmt.Println("Attempting to sign in...")
	var form Login
	// This will infer what binder to use depending on the content-type header.
	if c.Bind(&form) == nil {
		user, err := FindUser(form.Email, form.Password)
		if err != nil {
			fmt.Println(err)
		} else {
			session := FindOrCreateSession(c.Writer, c.Request)
			session.UserID = user.ID
			err = globalSessionStore.Save(session)
			if err != nil {
				panic(err)
			}

			fmt.Println("Signed in.")
			c.HTML(http.StatusOK, "viewVideo.tmpl.html", gin.H{"currentuser": "words"})
		}
	}

}
