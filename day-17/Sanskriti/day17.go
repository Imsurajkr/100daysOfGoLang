// Promoted Fields in Golang Structure

/*
  In Go structure, promoted fields are just like anonymous fields,
  the type of the field is the name of the field. We use this concept
  in the nested structure where a structure is a field in another structure,
  simply by just adding the name of the structure into another structure
  and it behaves like the Anonymous Field to the nested structure. And the
  fields of that structure (other than nested structure) are the part of the
  nested structure, such type of fields are known as Promoted fields. If the
  anonymous structure or nested structure and parent structure contains a field
  that has the same name, then that field doesn’t promote, only different
  name fields get promoted to the structure.
*/

package main

import "fmt"

// Structure
type details struct {

	// Fields of the details structure
	name   string
	age    int
	gender string
}

// Nested structure
type student struct {
	branch string
	year   int
	details
}

func main() {

	// Initializing the fields of the student structure
	values := student{
		branch: "CSE",
		year:   2010,
		details: details{

			name:   "Sumit",
			age:    28,
			gender: "Male",
		},
	}

	// Promoted fields of the student structure
	fmt.Println("Name: ", values.name)
	fmt.Println("Age: ", values.age)
	fmt.Println("Gender: ", values.gender)

	// Normal fields of the student structure
	fmt.Println("Year: ", values.year)
	fmt.Println("Branch : ", values.branch)
}

/*
   In the above program, we have two structures named as details and student.
   Where details structure is the normal structure and student structure is the
   nested structure which contains the details structure as fields in it just
   like anonymous fields. Now, the fields of the details structure,
   i.e, name, age, and gender are promoted to the student structure and
   known as promoted fields. Now, you can directly access them with the help
   of the student structure like values.name, values.age, and values.gender.
*/
