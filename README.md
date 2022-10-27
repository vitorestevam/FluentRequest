<div align="center">

<img  src="https://user-images.githubusercontent.com/55667307/198409675-2b624488-3bae-4f66-b5c8-5b07727e0bc7.png" alt="drawing" style="width:300px;"/>

# Fluent Request

<div align="left" />

Fluent Request is a small library implemented over http package. Its objective is to provide a friendly and useful interface when running http requests.

## Usage

Fluent request's usage is based on Fluent Interface or Chain Callbacks where the class setting methods returns the class itself. Check the example:

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

