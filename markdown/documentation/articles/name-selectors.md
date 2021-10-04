# Name selectors

## Direct

This follows the pattern:

```ruby
EXPRESSION.NAME
```

### Example

```ruby
a = 10
a.ToString()

(10 - 50 ** 2).ToFloat()
```

## Nested

This follows the pattern:

```ruby
EXPRESSION.NAME[.NAME[.NAME[...]]]
```

### Example

```ruby
a = MyObject
a.symbol_1.symbol_2.symbol_3
```