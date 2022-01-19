# Structs in Go

Go's *structs* are typed collections of fields. They are useful for grouping data together.

## Syntax and Struct Operations

### Declaring struct type

```go
type person struct {
    name string
    surname string
    age int
}
```

This declares a new (custom) struct type `person`. It has `name`, `surname` and `age` fields.

### Declaring struct variables

A syntax for creating a new struct:

```go
p := person{name: "Pavel", surname: "Zaichenkov", age: 32}
```

Short version:

```go
// The order of values corresponds to the order of fields declared in the struct.
p := person{"Pavel", "Zaichenkov", 32} // {name: "Pavel", surname: "Zaichenkov", age: 32}
```

Omitted fields are be zero-valued:

```go
p := person{"Pavel"} // {name: "Pavel", surname: "", age: 0}
```

Set struct to a zero value:

```go
var p person // {name: "", surname: "", age: 0}
```

### Accessing struct fields

Struct fields are accessed with a dot `.`:

```go
fmt.Printf("%s is %d years old.", p.name, p.age) // Pavel is 32 years old.
```

Structs are mutable. Fields can be redefined.

```go
p.age = 15
fmt.Println(p) // {Pavel Zaichenkov 15}
```

Structs of the same type can be assigned to each other:

```go
var p1 person
p2 := person{
    name: "Jaroslavs",
    surname: "Samcuks",
    age: 41,
}
p1 = p2
```

## Anonymous structs

It is possible to declare structs without creating a new data type. These types of structs are called *anonymous structs*.

```go
book := struct {
    title string
    author string
    pages int
}{
    title: "Animal Farm",
    author: "George Orwell",
    pages: 112, 
}
fmt.Println(book) // {Animal Farm George Orwell 112}
fmt.Printf("%+v", book) // {title:Animal Farm author:George Orwell pages:112}
```

### Use of anonymous structs in tests

You have already been using anonymous structs in tests. See an example below.

```go
func perimeter(width float64, height float64) float64 {
    return 2 * (width + height)
}

func TestPerimeter(t *testing.T) {
    for _, tc := range []struct {
        desc string
        width float64
        height float64
        want float64
    } {
        {"zeros", 0, 0, 0},
        {"int values", 5, 3, 16},
        {"float values", math.Pi, 5, 2 * (math.Pi + 5)},
    } {
        if got := perimeter(tc.width, tc.height); got != tc.want {
            t.Errorf("%s: got %.2f want %.2f", tc.desc, got, tc.want)
        }
    }
}
```

## Storing Structs in a File

Structs and struct slices can be stored in a
[JSON](https://en.wikipedia.org/wiki/JSON) file while preserving a structure
(using `encoding/json` package).

### Encoding 

Only capitalized fields (e.g. `Population`, not `population`) can be stored in a
file.

```go
f, err := os.Create("countries.json")
if err != nil {
    log.Fatal(err)
}
defer f.Close()
w := bufio.NewWriter(f)
defer w.Flush()
// countries is a slice of anonymous structs in this example.
countries := []struct {
    Name, Capital string
    population    int // This is an unexported field.
}{
    {"Latvia", "Riga", 1_902_000},
    {"Switzerland", "Bern", 8_637_000},
}
if err := json.NewEncoder(w).Encode(countries); err != nil {
    log.Fatal(err)
}
```

The content of `countries.json`:

```
[{"Name":"Latvia","Capital":"Riga"},{"Name":"Switzerland","Capital":"Bern"}]
```

### Decoding

```go
type country struct{
    Name, Capital string
    population int
    Space int
}

var countries []country
f, err := os.Open("countries.json")
if err != nil {
    log.Fatal(err)
}
defer f.Close()
if err := json.NewDecoder(bufio.NewReader(f)).Decode(&countries); err != nil {
  log.Fatal(err)
}
fmt.Printf("%+v", countries)
// [{Name:Latvia Capital:Riga population:0 Space:0} {Name:Switzerland Capital:Bern population:0 Space:0}]
```

# Field Tags

[Go struct specification](https://go.dev/ref/spec#Struct_types) allows adding tags to the fields e.g.

```go
struct {
    field1 string "this is my tag"
    field2 int    `json:"foo"`
}
```

The tags are arbitrary strings, however various libraries use them for [metadata](https://en.wikipedia.org/wiki/Metadata) purpose.
Examples of such libraries include `encoding/json`, `encoding/xml`, `encoding/gob`, etc. As you can see, encoding libraries use this
feature a lot! Defacto syntax for the tags is the following: \`library1:"data1[,data2[,...]]"[ library2:"..."]\`.

Consider the following [example](https://go.dev/play/p/DlOaU9w0Ux4):

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Record struct {
	Field1 string `json:"-"`
	Field2 string `json:"foo,omitempty"`
	Field3 string `json:"bar"`
	Field4 string
}

func main() {
	out, err := json.MarshalIndent([]Record{
		{Field1: "a", Field2: "b", Field3: "c", Field4: "d"},
		{Field1: "a"},
	}, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
```

and its output

```text
[
  {
    "foo": "b",
    "bar": "c",
    "Field4": "d"
  },
  {
    "bar": "",
    "Field4": ""
  }
]
```

Note the following aspects:

* `Record.Field1` is omitted.
* `Record.Field2` is renamed to `foo` and is omitted for empty strings.
* `Record.Field3` is renamed to `bar` and is present even for empty values.
* `Record.Field4` goes untouched.

## Methods

https://go.dev/tour/methods/1

Go does not have [classes](https://en.wikipedia.org/wiki/Class_(computer_programming)), however it allows to define methods on types.
Methods are functions with a special _receiver_ argument, which appears between `func` keyword and the method name.

In this example, the `Abs` method has a receiver of type `Vertex` named `v`:

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

Note: I recommend `Type.Method` syntax, in which case you always know both - method name and its receiver type.

Methods are attached to their receivers and allow to create functions that implement algorithms without knowing about
about the data those algorithms work with. We will be dealing with this in the future. For now it's important that you
understand what methods are and can use them as almost all of the libraries provide structures with methods e.g.
[os.File](https://pkg.go.dev/os#File), which has more than 20 methods like `os.File.Read()`, `os.File.Close()`, etc.
