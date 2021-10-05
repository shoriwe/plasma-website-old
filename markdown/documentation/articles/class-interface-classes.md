# Classes

## Definition

```ruby
class MyClass
    name = "Machine"
    
    # Overwrite ToString
    def ToString()
        return self.name
    end
    
    def print(message)
        println(self.name + " says " + message)
    end
end
```

## Inheritance

Inheritance can be achieved by capturing in parentheses the subclasses.

```ruby
class CustomInteger(Value, Integer) # Inherits from Integer and Value 
end
```

By default, every class inherits from `Value`, even if it was not specified.

## Self

You can access the constructed object from its methods by using the keyboard `self`

```ruby
class MyClass
    name = "Machine"
    
    def print(message)
        println(self.name + " says " + message)
    end
end
```

## There is no static methods

By design of the language when you need to isolate a functionality in a namespace, use the keyboard `module`.