<p>
  <h2 align="center"> DAY 2</h2>

### Resolving day 1 issue: 
</p>
You need to specify in your app.yaml file what your app access point is. You need to specify

   `<|`

	`application: hello`
    
    `|`

### refer to this [link](https://cloud.google.com/appengine/docs/standard/go/quickstart)
<p>
- The `entry point` of each go program is main.main, i.e. a function called main in a package called main. You have to provide such a main package.
- `GAE` is an exception though. They add a main package, containing the main function automatically to your project. Therefore, you are not allowed to write your own.
</p>

