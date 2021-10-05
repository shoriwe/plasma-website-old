# Super

When you need to cast an object as other class, especially useful for the methods defined in that other class. You can
approach the function super to cast a proxy to that target class.

```ruby
class MyNumber(Float, Integer)
    def Initialize(number)
        super(self, Float).Initialize()
        super(self, Integer).Initialize()
    end
end
```
