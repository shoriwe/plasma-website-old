# Types

## Integers

Numeric values internally represented with `int64`.

Notice that all the notations supports **`_`** to pretty write every number.

### Notations

#### Decimal

All life notation for numbers.

```ruby
a = 10
b = 1_000_000
```

#### Hex

A **`python`** like hex number notation.

```ruby
a = 0xA # With letters in upper case
b = 0xa # With letters in lower case
c = 0xFF_FF
d = 0XfF_aFD
```

#### Binary

A **`python`** like binary number notation.

```ruby
a = 0b01
b = 0B01
c = 0b1111_1111_1101
d = 0B0000_0000_000_0001
```

#### Octal

A **`python`** like octal number notation.

```ruby
a = 0o075
b = 0O075
c = 0o556_012
d = 0O0
```

## Floats

Internally floating numbers are been represented with **`float64`**

### Notations

#### Basic

The all life decimal number notation.

```ruby
a = 1000.0001
a = 1_000.000_1
```

#### Scientific

A **`python`** like scientific number notation.

```ruby
a = 10e-2 # 0.1
a = 10e+2 # 1000
```

## Strings

Notice that strings definitions support multiple lines.

### Escape characters

#### Hex

You will need to approach the next pattern `\xNN` or `\XNN`.

```ruby
a = '\x43' # C
```

#### Unicode

You will need to approach the next pattern `\uNNNN`.

```ruby
a = '\u0043' # C
```

### Single quote

```ruby
a = 'Hello world!'
b = 'Hello
World from other line'
```

### Double quote

```ruby
a = 'Hello world!'
b = 'Hello
World from other line'
```

## Bytes

For bytes, the same notation for strings apply, but you will also include a `b` as the string prefix.

```ruby
a = b"My bytes"
b = b"Multi
line
bytes
strings"
```

## Bool

```ruby
a = True
b = False
```

## None

```ruby
a = None
```

## Tuples

Tuples are a immutable data type, anyway, you still can modify its members.

```ruby
a = (1, "hello world", 1 + 4 % 10)
```

## Arrays

Array values can be modified without problems

```ruby
a = [1, "hello world", 1 + 4 % 10]

a[0] = "John" # ["John", "hello world", 1 + 4 % 10]
```

## Hash Tables

```ruby
people = {
    "John": (30, "New York"),
    "Alice": (32, "Chicago")
}
```