# Switch

```ruby
name = "John"
switch name
case "John", "Bob"
    gender = "masculine"
case "Alice", "Karol"
    gender = "femenine"
default
    gender = "Couldn't guess"
end

println("The name: " + name + " is usually used as " + gender)
```