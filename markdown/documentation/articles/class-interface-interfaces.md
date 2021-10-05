# Interfaces

## Definition

Interfaces are almost equal to Classes, the only difference is that they only support method definitions inside its
body.

```ruby
interface IMachine
    # Overwrite ToString
    def ToString()
        return selfname
    end
    
    def print(message)
        println(self.name + " says " + message)
    end
end
```

## Inheritance

Inheritance can be achieved by capturing in parentheses the sub-interfaces.

```ruby
interface ICustomInteger(Value, Integer) # Inherits from Integer and Value 
end
```

By default, every interface inherits from `Value`, even if it was not specified.

## Self

You can access the constructed object from its methods by using the keyboard `self`

```ruby
interface IObject
    def print(message)
        println(self.name + " says " + message) # Don't worry about not having `name` defined, it will be resolved at runtime when you inherit this interface.
    end
end
```

## There is no static methods

By design of the language when you need to isolate a functionality in a namespace, use the keyboard `module`.