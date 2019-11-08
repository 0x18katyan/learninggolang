package main

import "fmt"

type person struct {
	first, last string
}

type graduate struct {
	person
	degreeName, university string
}

type postGraduate struct {
	graduate
	degreeName, university string
}

type doctorate struct {
	postGraduate
	specialization, university string
}

func (education graduate) brag() {
	fmt.Printf("I am a graduate from %v with a degree in %v.\n", education.university, education.degreeName)
}

func (education postGraduate) brag() {
	fmt.Printf("I am a post-graduate from %v with a degree in %v.\n", education.university, education.degreeName)
}

func (education doctorate) brag() {
	fmt.Printf("I am a Doctor from %v with specialization in %v.\n", education.university, education.specialization)
}

type collegeStudent interface {
	brag()
}

func discount(student collegeStudent) {
	fmt.Println("Congratulations! For being a College Student, you get NOTHING!!!!")
}

func main() {

	student1 := graduate{
		person: person{
			first: "Hobo",
			last:  "Scientist"},
		degreeName: "Bachelors of Science in Computer Networking and IT Security",
		university: "London Metropolitan University",
	}

	student2 := postGraduate{
		graduate: graduate{
			person: person{
				first: "Ashim",
				last:  "Mahara"},
			degreeName: "Bachelors of Science in Computer Networking and IT Security",
			university: "London Metropolitan University",
		},
		degreeName: "Masters of Science in Data Analytics",
		university: "London Metropolitan University",
	}

	student1.brag()
	student2.brag()
	discount(student1)

}
