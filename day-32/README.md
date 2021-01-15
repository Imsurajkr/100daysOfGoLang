# Creating good Go Packages
* successful package is that its elements must be related in some way.
* It is better to split the functionality of a packageunnecessarily into multiple packages than to add too much functionality to a single Go package.
* packages should be made simple and stylish but not too simplistic and fragile.
* you should use your own packages first for a reasonable amount of time before giving them to the public.
* your packages should not export an endless list of functions.try to title your functions
using descriptive but not very long names.
* Interfaces can improve the usefulness of your functions, so when you think it is
appropriate, use an interface instead of a single type as a function parameter or
return type.
* When updating one of your packages, try not to break things and create
incompatibilities with older versions unless it is absolutely necessary.
* When developing a new Go package, try to use multiple files in order to group
similar tasks or concepts.
*  try to follow the rules that exist in the Go packages of the standard
library.
* Do not create a package that already exists from scratch. Make changes to the
existing package and maybe create your own version of it.
* Nobody wants a Go package that prints logging information on the screen. It
would be more professional to have a flag for turning on logging when needed.
* The Go code of your packages should be in harmony with the Go code of your
programs. This means that if you look at a program that uses your packages and
your function names stand out in the code in a bad way, it would be better to
change the names of your functions. As the name of a package is used almost
everywhere, try to use a concise and expressive package name.
* It is more convenient if you put new Go type definitions near the place that they
will be used for the first time because nobody, including you, wants to search
source files for definitions of new data types.
* Try to create test files for your packages, because packages with test files are
considered more professional than ones without them; small details make all the
difference and give people confidence that you are a serious developer! Notice
that writing tests for your packages is not optional and that you should avoid
using packages that do not include tests.
* Finally, do not write a Go package because you do not have anything better to do
– in that case, find something better to do and do not waste your time!
* Always remember that apart from the fact that the actual Go code in a
package should be bug-free, the next most important element of a
successful package is its documentation, as well as some code examples
that clarify its use and showcase the idiosyncrasies of the functions of the
package.