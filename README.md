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

##### 2.1. Writing to a file

```
length, err := io.WriteString(file, content)
```

##### 2.2. Reading from a file

```
func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file:", string(data))
}
```

##### 2.3. Reading from a web page

```
response, err := http.Get(url)
```

##### 2.4. Posting data to a web server

```
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
```

#### 3. Managing program workflow

##### 3.1. Switch statements

There is no use of the `break` keyword in Go. As soon as one of these cases is evaluated as true, it will execute the
code in that case and then jump to the end of the switch statement.

By default, the `switch` statement matches goes through all the case statement from top to bottom and tries to find case
expression that matches the switch expression. Once the matching case is found, it exits and does not consider the other
cases.

`fallthrough` allows way around this limitation by looping through all the subsequent cases to look for any other match.

Go also allows for placing the executed statement after the `switch` statement then evaluating it from a variable. This
is shown below:

_Normal flow_

```
rand.Seed(time.Now().Unix())
	dayOfWeek := rand.Intn(8) + 1
	fmt.Println("Day of Week", dayOfWeek)

	var result string
	switch dayOfWeek {
	case 1:
		result = "It's Sunday. Going riding!"
```

_Alternative flow_

```
rand.Seed(time.Now().Unix())
	var result string
	switch dayOfWeek := rand.Intn(8) + 1; dayOfWeek {
	case 1:
		result = "It's Sunday. Going riding!"
```

##### 3.2. For loops

Go offers several ways of looping through a collection using the for loop.

Given an array of strings `authors := []string{"Charles Dickens", "Leo Tolstoy", "Vladimir Nabokov"}`, these are some
possible ways to loop through each of the items.

_Traditional for loop as in Java and C++_

```
for i := 0; i < len(authors); i++ {
		fmt.Println(authors[i])
	}
```

_Go-styled for loops_

```
for i := range authors {
		fmt.Println(authors[i])
	}
```

```
for _, author := range authors {
		fmt.Println(author)
	}
```

```
value := 1
	for value < 10 {
		fmt.Println("Value:", value)
		value++
	}
```

_Looping till a condition is reached_

```
sum := 1
	for sum < 1000 {
		sum += sum
		if sum > 300 {
			goto theEnd
		}
	}

theEnd:
	fmt.Println("End of program")
```

#### 4. Managing complex types and collections

##### 4.1. Memory management

The Go runtime is statically linked into the application with memory being allocated and deallocated automatically.

`make()` or `new()` can be used to initialize maps, slices and channels

The `new()` function allocates but does not initialize memory. That is to say, it allocates a zeroed storage but returns
a memory address. The `make()` function however, both allocates and initializes memory - it allocates a non-zeroed
storage and returns a memory address.

Therefore, using new instead of make in initializing a map can cause a crash. `make()` should then be used whenever the
intent is to add data immediately to the initialized collection.

`m := make(map[string]int)`

##### 4.2. Pointers

Pointers are variables that store the memory address of other variables. A fun way to think about them is, a person
through whom you get to another person.

_Unassigned pointer_

```
var p *int
fmt.Println("Value of p:", p)       //<nil>
fmt.Println("Invalid pointer:", *p) //Crashes the application
```

_Valid assigned pointer_

```
someInteger := 67
var p = &someInteger
fmt.Println("Value of variable someInteger:", *p) //67
fmt.Println("Value of memory address of someInteger:", p) // The memory address of someInteger (0xc0000b2008)

```

_Changing a value from a pointer_

It should be noted that changing a value from a pointer changes the actual variable being pointed. This can be seen in the code snippet below:

```
four := 4
pointer1 := &four
fmt.Println("four from pointer:", *pointer1)

*pointer1 = *pointer1/2 // Same as four/2
fmt.Println("Pointer 1:", *pointer1) //2
fmt.Println("four:", four) //2

```


