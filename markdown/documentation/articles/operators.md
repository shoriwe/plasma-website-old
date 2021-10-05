# Operators

This means operations with the structure:

```
LEFT OPERATOR RIGHT
```

## Binary operators

### Addition `+`

#### Supports

- Integer/Floats with Integers/Floats
- Strings with Strings
- Bytes with Bytes
- Arrays with Arrays

### Subtraction `-`

#### Supports

- Integer/Floats with Integers/Floats

### Multiplication `*`

#### Supports

- Integer/Float with Integer/Float
- String/Integer with Integer/String

### Division `/`

#### Supports

- Integer/Float with Integer/Float

### Floor division `//`

#### Supports

- Integer/Float with Integer/Float

### Modulus `%`

#### Supports

- Integer with Integer

### Power Of `**`

#### Supports

- Integer/Float with Integer/Float

### Equals `==`

#### Supports

- Any with Any

### Not equals `!=`

#### Supports

- Any with Any

### Less than `<`

#### Supports

- Integer/Float with Integer/Float

### Less or equal than `<=`

#### Supports

- Integer/Float with Integer/Float

### Greater than `>`

#### Supports

- Integer/Float with Integer/Float

### Greater or equal than `>=`

#### Supports

- Integer/Float with Integer/Float

### Bitwise left `<<`

#### Supports

- Integer

### Bitwise right `>>`

#### Supports

- Integer

### Bitwise or `|`

#### Supports

- Integer

### Bitwise and `&`

#### Supports

- Integer

### Bitwise xor `^`

#### Supports

- Integer

### Bool And `and`

#### Supports

- Any with Any

### Bool Or `or`

#### Supports

- Any with Any

### Bool Xor `xor`

#### Supports

- Any with Any

### Contains `in`

#### Supports

- Any with Array/Tuple/HashTable
- String with String
- Bytes with Bytes

## Unary

This means operations with the structure:

```
OPERATOR VALUE
```

### Negate Bits `~`

#### Supports

- Integer

### Bool negation `not` or `!`

#### Supports

- Any