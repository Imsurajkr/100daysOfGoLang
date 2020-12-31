// Promoted Methods in Golang Structure
/*
   In Go structure, the working of promoted methods is just like Promoted Fields.
   We use this concept in the nested structure where a structure is a field in
   another structure, simply by just adding the name of the structure into
   another structure and it behaves like the Anonymous Field to the nested structure.
   And the methods of that structure (other than nested structure) are the part of the
   nested structure, such type of methods are known as Promoted Methods. Or in other words,
   prompted methods are those methods which are implemented by the child structure,
   and is accessible by the parent structure.
*/

package main

import "fmt"

// Structure
type details struct {

	// Fields of the details structure
	name    string
	age     int
	gender  string
	psalary int
}

// Nested structure
type employee struct {
	post string
	eid  int
	details
}

// Method
func (d details) promotmethod(tsalary int) int {
	return d.psalary * tsalary
}

func main() {

	// Initializing the fields of the employee structure
	values := employee{
		post: "HR",
		eid:  4567,
		details: details{

			name:    "Sumit",
			age:     28,
			gender:  "Male",
			psalary: 890,
		},
	}

	// Promoted fields of the
	// employee structure
	fmt.Println("Name: ", values.name)
	fmt.Println("Age: ", values.age)
	fmt.Println("Gender: ", values.gender)
	fmt.Println("Per day salary: ", values.psalary)

	// Promoted method of the
	// employee structure
	fmt.Println("Total Salary: ", values.promotmethod(30))

	// Normal fields of
	// the employee structure
	fmt.Println("Post: ", values.post)
	fmt.Println("Employee id: ", values.eid)
}

/*
    Important Points:

   1.If the child and the parent structure contain a method that has the same name
	 but different type of receiver, then both the methods are available in the
	 parent structure as shown in Example 2. Here both child and parent structure
	 contain methods of the same names.

   2.If child structure contains two methods with the same name and same receiver,
	 then these methods are not promoted in the parent structure and if you try to do,
	 then the compiler will give an error.
*/
