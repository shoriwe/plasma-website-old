# Redo

```ruby
a = 0
while a == 0
    a += 1
    if a == 2
        break
    else
        redo
    end
end

println(a)
```