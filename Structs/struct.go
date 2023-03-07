package main

func main() {

	// defining a new type
	type courseMeta struct {
		author string
		level  string
		rating float64
	}

	var joshCourse = courseMeta{
		author: "jake",
		level: "josh",
		rating: 120.4,
	}

	//jakeCourse := new(courseMeta)

	jakeCourse := courseMeta{
		author: "jake",
		level: "josh",
		rating: 10.2,
	}


}
