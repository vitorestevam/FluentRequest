# Fluent Request

This package is an option to do http requests using fluent interface

``` go
resp, err := FluentRequest().
                Method("GET").
                Url("https://jsonplaceholder.typicode.com/todos/1").
                Run()

body, _ := ioutil.ReadAll(resp.Body)

fmt.Println(string(body))

// output
//{
//   "userId": 1,
//   "id": 1,
//   "title": "delectus aut autem",
//   "completed": false
//}
```

