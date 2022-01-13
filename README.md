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
