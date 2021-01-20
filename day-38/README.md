# Reflection
* Reflection gives you the ability to examine types at runtime. It also allows you to examine, modify, and create variables, functions, and structs at runtime.
##### Reflection in Go is built around three concepts:
* Types
* Kinds
* Values
#### Why Reflection?
* Most of the time, variables, types, and functions in Go are pretty straightforward. When you need a type, you define a type:
        type Foo struct {
        A int
        B string
        }
* When you need a variable, you define a variable:
        var x Foo
* And when you need a function, you define a function:
        func DoSomething(f Foo) {
        fmt.Println(f.A, f.B)
        }
* But sometimes you want to work with variables at runtime using information that didn’t exist when the program was written. Maybe you’re trying to map data from a file or network request into a variable. Maybe you want to build a tool that works with different types. In those situations, you need to use reflection. 
