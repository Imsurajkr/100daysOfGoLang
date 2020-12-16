<p>
  <h2 align="center"> DAY 2</h2>

# Two Go rules :
  * ### You either use a Go package or you do not include it

        
          - Go has strict rules about package usage.
          - Therefore, you cannot just include any package you might think that you will need and then not use it afterward.
          - So, using an underscore character in front of a package name in the import list will not create an error message in the compilation process
        

  * ### There is only one way to format curly braces
    ```GO
    package main
    import (
           "fmt"
    )
    func main()
    {
          fmt.Println("Go has strict rules for curly braces!")
    }
    ```
      * <p>Putting the opening curly brace ({) in its own line will make the Go compiler insert a semicolon at the end of the previous line (func main()), which is the cause of the error message.
      </p> 
</p>

<p> 
  
  # Downloading External packages for Go

   <img align="center" src="‪https://github.com/Imsurajkr/100daysOfGoLang/blob/main/day-2/bhimra/externalcode.png" width="230">

   <p>
      <h3>$ go get -v rsc.io/quote/v3</h3>

      - download the external package from web
  </p>



