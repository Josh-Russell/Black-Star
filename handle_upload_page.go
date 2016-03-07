package main

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

const videoIDLength = 10

type video struct {
	ID          string
	UserID      string
	Name        string
	Location    string
	Size        int64
	CreatedAt   time.Time
	Description string
}

func HandleNavigateToUpload(c *gin.Context) {
	session := RequestSession(c.Request)

	if session != nil {
		c.HTML(http.StatusOK, "upload.tmpl.html", gin.H{"currentuser": session.UserID})
	} else {
		c.HTML(http.StatusOK, "login.tmpl.html", nil)
	}
}
func createNewVideo(user *User) *video {
	return &video{
		ID:        GenerateID("vid", videoIDLength),
		UserID:    "5",
		CreatedAt: time.Now(),
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
		c.HTML(http.StatusOK, "upload.tmpl.html", gin.H{
			"Error": err,
			"Video": video,
		})
		return
	}

	c.HTML(http.StatusFound, "index.tmpl.html", gin.H{"currentuser": 5})
}
func (vid *video) HandleCreateVideo(file multipart.File, headers *multipart.FileHeader) error {
	// Move our file to an appropriate place, with an appropriate name
	vid.Name = headers.Filename
	vid.Location = vid.ID + filepath.Ext(vid.Name)

	// Open a file at target location
	savedFile, err := os.Create("./data/videos/" + vid.Location)
	if err != nil {
		return err
	}

	defer savedFile.Close()

	// Copy the uploaded file to the target location
	size, err := io.Copy(savedFile, file)
	if err != nil {
		return err
	}
	vid.Size = size

	return err
	// Save the image to the database
	//return globalImageStore.Save(vid)
}
