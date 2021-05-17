package main

import "fmt"

type Course struct {
	Name   string
	Videos []Video
}

type Video struct {
	Name   string
	Course Course
}

func main() {
	video1 := Video{Name: "Introducction"}
	video2 := Video{Name: "What is Go?"}
	video3 := Video{Name: "Install Go"}

	course := Course{Name: "Go Professional Course"}

	video1.Course = course
	video2.Course = course
	video3.Course = course

	videos := []Video{video1, video2, video3}
	course.Videos = videos

	fmt.Println(course.Name)
	for _, video := range course.Videos {
		fmt.Println(video.Name)
	}
}
