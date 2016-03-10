package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	// Ensure our goroutines run across all cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

const imageIDLength = 10

type Video struct {
	ID          string
	title       string
	Username    string
	Description string
	Location    string
	upvotes     int64
	downvotes   int64
	mature      bool
}

func NewVideo(user *User) *Video {
	return &Video{
		ID:       GenerateID("vid", videoIDLength),
		Username: user.ID,
	}
}

func (video *Video) StaticRoute() string {
	return "/vid/" + video.Location
}

func (video *Video) StaticThumbnailRoute() string {
	return "/vid/thumbnail/" + video.Location
}

func (video *Video) StaticPreviewRoute() string {
	return "/vid/preview/" + video.Location
}

func (video *Video) ShowRoute() string {
	return "/video/" + video.ID
}

// A map of accepted mime types and their file extension
var mimeExtensions = map[string]string{
	"image/png":  ".png",
	"image/jpeg": ".jpg",
	"image/gif":  ".gif",
}

func (video *Video) CreateFromFile(file multipart.File, headers *multipart.FileHeader) error {

	// Move our file to an appropriate place, with an appropriate name
	video.title = headers.Filename
	video.Location = video.ID + filepath.Ext(video.title)

	// Open a file at target location
	savedFile, err := os.Create("M:/videos/" + video.Location)
	if err != nil {
		return err
	}

	defer savedFile.Close()

	// Copy the uploaded file to the target location
	size, err := io.Copy(savedFile, file)
	if err != nil {
		return err
	}
	var in = size
	fmt.Println(in)

	// Save the video to the database
	return globalVideoStore.Save(video)
}
