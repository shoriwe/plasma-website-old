# Lambda functions

You will use the keyboard `lambda` followed by the arguments that requires and `:` then the expression to execute.

## Example

```ruby
hello = lambda: println("Hello world!")

hello = lambda name, last_name: println("Hello " + name + " " + last_name)

pow = map((lambda n: n**2), range(0, 1000, 1))
```