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


