package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const videoIDLength = 10

func HandleNavigateToUpload(c *gin.Context) {
	session := RequestSession(c.Request)

	if session != nil {
		c.HTML(http.StatusOK, "upload.tmpl.html", gin.H{"currentuser": session.UserID})
	} else {
		c.HTML(http.StatusOK, "login.tmpl.html", nil)
	}
}
func createNewVideo(user *User) *Video {
	return &Video{
		ID:        GenerateID("vid", videoIDLength),
		Username:  user.Username,
		upvotes:   0,
		downvotes: 0,
	}
}
func HandleNewVideo(c *gin.Context) {

	user := RequestUser(c.Request)
	if user == nil {
		user = &User{
			ID:             "5",
			Email:          "hello@gmail.com",
			HashedPassword: "browns",
			Username:       "fred",
		}
	}
	video := createNewVideo(user)
	video.Description = c.Request.FormValue("description")
	video.title = c.Request.FormValue("title")
	video.mature = c.Request.FormValue("matureContent") == "on"

	file, headers, err := c.Request.FormFile("file")

	// No file was uploaded.
	if file == nil {
		c.HTML(http.StatusOK, "upload.tmpl.html", gin.H{
			"Error": errNoImage,
			"Video": video,
		})
		return
	}

	// A file was uploaded, but an error occurredd
	if err != nil {
		panic(err)
	}

	defer file.Close()

	err = video.HandleCreateVideo(file, headers)
	if err != nil {
		fmt.Println("Failed.")
		c.String(http.StatusExpectationFailed, "failed to upload to the database", gin.H{
			"Error": err,
			"Video": video,
		})
		return
	}
	c.Redirect(303, "/")
	//c.HTML(http.StatusFound, "index.tmpl.html", gin.H{"currentuser": 5})
}
func (vid *Video) HandleCreateVideo(file multipart.File, headers *multipart.FileHeader) error {
	// Move our file to an appropriate place, with an appropriate name
	vid.Location = vid.ID + filepath.Ext(headers.Filename)

	// Open a file at target location
	savedFile, err := os.Create("videos/" + vid.Location)
	if err != nil {
		return err
	}

	defer savedFile.Close()

	// Copy the uploaded file to the target location
	_, err = io.Copy(savedFile, file)
	if err != nil {
		return err
	}
	// Save the image to the database
	return globalVideoStore.Save(vid)
}
