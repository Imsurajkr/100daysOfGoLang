<p>
  <h2 align="center"> DAY 1</h2>
`package hello`  // package name `hello`   
`import (
    "fmt"
)
func Hello() {
    fmt.Println(" run if you can")
}`

Above code showed a `ERROR : "go run: cannot run non-main package"` due to package name is not main.</p>

- To fix it, I just need to name the package to main. But I don't understand why I need to do that. I should be able to name the package whatever I want.
- Another question, I know main function is the entry point of the program, you need it. otherwise it will not work. But I see some codes that didn't have main function still works.
- Click on this link, the example at the bottom of the page didn't use package main and main function, and it still works. just curious <h4>WHY?</h4>

go to this [link](https://developers.google.com/appengine/docs/go/gettingstarted/usingdatastore)</p>








