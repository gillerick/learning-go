## Learning Go

#### 1. Creating Reusable Code

In Go, it is possible to attach functions to custom types. They are then referred to as _**methods**_

As opposed to more completely OOP languages such as Java where _methods are members of a class_, in Go a _method is a
member of a type_

```
type Dog struct {
	Breed string
	Weight int
	Sound string
}

func (d Dog) Speak(){
	fmt.Println(d.Sound)
}
```

In the above function,`(d Dog)` is known as the **receiver** where `d` is the **identifier** and `Dog` is the **type**.
Here, we are _passing in a reference_ to a dog object.

#### 2. Working with Files and the Web

Reading a file always produces an array of bytes.

`defer` keyword defers/delays the execution of a function until the surrounding function returns. The deferred call's
arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

Golang's `net/http` standard library provides HTTP client and server implementations

##### 1. Writing to a file

```
length, err := io.WriteString(file, content)
```

##### 2. Reading from a file

```
func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file:", string(data))
}
```

##### 3. Reading from a web page

```
response, err := http.Get(url)
```

##### 4. Posting data to a web server

```
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
```

#### 3. Managing program workflow

There is no use of the `break` keyword in Go. As soon as one of these cases is evaluated as true, it will execute the
code in that case and then jump to the end of the switch statement.

By default, the `switch` statement matches goes through all the case statement from top to bottom and tries to find case
expression that matches the switch expression. Once the matching case is found, it exits and does not consider the other
cases.

`fallthrough` allows way around this limitation by looping through all the subsequent cases to look for any other match.