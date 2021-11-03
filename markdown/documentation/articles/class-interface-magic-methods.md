# Magic methods

Magic methods are used to operate with the language internals, to overwrite operators, to simplify an object API.

## Default

These methods are inherited from `Value` and are inherited even if it is not explicit.

#### Initialize(ANY_NUMBER_OF_ARGUMENTS, ...)

First method called when an object is created, useful no object configuration.

#### Negate()

```ruby
not self
!self
```

Negate the boolean interpretation of the object. Auto selected when used `not` and `!`

#### And(right)

```ruby
self and right
```

And boolean comparison. Auto called when using the keyboard `and`.

#### RightAnd(self)

```ruby
left and self
```

And boolean comparison. Auto called when using the keyboard `and`.

#### Or(right)

```ruby
self or right
```

Or boolean comparison. Auto called when using the keyboard `or`.

#### RightOr(left)

```ruby
left or self
```

Or boolean comparison. Auto called when using the keyboard `or`.

#### Xor(right)

```ruby
self xor right
```

Xor boolean comparison. Auto called when using the keyboard `xor`.

#### RightXor(left)

```ruby
left or self
```

Xor boolean comparison. Auto called when using the keyboard `xor`.

#### Hash()

Returns the hash for the object.

This method is used by the HashTable class to store new entries and select existing ones

#### Copy()

Copies the object.

#### Call(ANY_NUMBER_OF_ARGUMENTS, ...)

This method is called every time you use an object as a function call.

#### Class()

Returns the class that constructed the object.

#### SubClasses()

Returns an array of all the classes, even nested inheritance, of the object.

#### ToInteger()

Returns the integer representation of the object.

#### ToFloat()

Returns the float representation of the object.

#### ToString()

Returns the string representation of the object.

#### ToBool()

Returns the boolean representation of the object.

#### ToArray()

Returns the array representation of the object.

#### ToTuple()

Returns the tuple representation of the object.

#### ToBytes()

Returns the bytes representation of the object.

#### GetInteger()

Returns the internal integer.

#### GetBool()

Returns the internal boolean.

#### GetBytes()

Returns the internal bytes.

#### GetString()

Returns the internal string.

#### GetFloat()

Returns the internal float.

#### GetContent()

Returns the internal array of object for arrays and tuples.

#### GetKeyValues()

Returns the internal key and value of the object.

#### SetBool(new_value)

Sets the internal boolean to a specific value.

#### SetBytes(new_value)

Sets the internal bytes to a specific value.

#### SetString(new_value)

Sets the internal string to a specific value.

#### SetInteger(new_value)

Sets the internal integer to a specific value.

#### SetFloat(new_value)

Sets the internal float to a specific value.

#### SetContent(new_value)

Sets the internal content to a specific array.

#### SetKeyValues(new_value)

Sets the internal key and values to a specific map.

## Iterable specific

#### Index(index)

Used when indexing an object.

```ruby
self[index]
```

#### Iter

This method is called by many internal functions to reinterpret some objects as generators.

```ruby
for element in self # Here it is called to operate self as a generator
    println(element)
end
```

#### Contains(element)

This method is used by the operation `in`.

```ruby
element in self
```

### Generator specific

#### HasNext()

Returns `True` when it still has elements to iterate. And `False` when it not.

#### Next()

Returns the next value of the iteration.

### Array and Hash Table specific

#### Delete(index)

Deletes the element by its index.

#### Assign(index, value)

Used when assign is used with and index.

```ruby
self[index] = value
```

### Array specific

#### Append(value)

Puts the `value` at the end of the array.

#### Pop

Returns the last element and remove it from the array.

## String, Bytes, Array, Tuple and Hash specific

#### Length()

Returns the length of the object.

## String and Bytes specific

#### Lower()

Returns the lower case of the string.

#### Upper()

Returns the upper case of the string.

#### Split(sep)

Returns an array of strings which is the result by splitting `self` by `sep`

#### Replace(target, replace)

Replaces the `target` with `replace`.

## Numeric specific

#### Negative()

Auto used when approached the unary expression `-`

```ruby
-self
```

#### Add(right)

```ruby
self + right
```

Adds the values. Auto called when using the operator `+`.

#### RightAdd(self)

```ruby
left + self
```

Adds the values. Auto called when using the operator `+`.

#### Sub(right)

```ruby
self - right
```

Subtracts the values. Auto called when using the operator `-`.

#### RightSub(left)

```ruby
left - self
```

Subtracts the values. Auto called when using the operator `-`.

#### Mul(right)

```ruby
self * right
```

Multiplies the values. Auto called when using the operator `*`.

#### RightMul(left)

```ruby
left * self
```

Multiplies the values. Auto called when using the operator `*`.

#### Div(right)

```ruby
self / right
```

Divides the values. Auto called when using the operator `/`.

#### RightDiv(left)

```ruby
left / self
```

Divides the values. Auto called when using the operator `/`.

#### FloorDiv(right)

```ruby
self // right
```

Floor divides the values. Auto called when using the operator `//`.

#### RightFloorDiv(left)

```ruby
left // self
```

Floor divides the values. Auto called when using the operator `//`.

#### Pow(right)

```ruby
self ** right
```

Power of the values. Auto called when using the operator `**`.

#### RightPow(left)

```ruby
left ** self
```

Power of the values. Auto called when using the operator `**`.

#### Equals(right)

```ruby
self == right
```

Equals the values. Auto called when using the operator `==`.

#### RightEquals(left)

```ruby
left == self
```

Equals the values. Auto called when using the operator `==`.

#### NotEquals(right)

```ruby
self != right
```

Not equals the values. Auto called when using the operator `!=`.

#### RightNotEquals(left)

```ruby
left != self
```

Not equals the values. Auto called when using the operator `!=`.

#### GreaterThan(right)

```ruby
self > right
```

Greater than of the values. Auto called when using the operator `>`.

#### RightGreaterThan(left)

```ruby
left > self
```

Greater than of the values. Auto called when using the operator `>`.

#### LessThan(right)

```ruby
self < right
```

Less than of the values. Auto called when using the operator `<`.

#### RightLessThan(left)

```ruby
left < self
```

Less than of the values. Auto called when using the operator `<`.

#### GreaterThanOrEqual(right)

```ruby
self >= right
```

Greater or equal than of the values. Auto called when using the operator `>=`.

#### RightGreaterThanOrEqual(left)

```ruby
left >= self
```

Greater or equal than of the values. Auto called when using the operator `>=`.

#### LessThanOrEqual(right)

```ruby
self <= right
```

Less or equal than of the values. Auto called when using the operator `<=`.

#### RightLessThanOrEqual(left)

```ruby
left <= self
```

Less or equal than of the values. Auto called when using the operator `<=`.

## Integer specific

#### Mod(right)

```ruby
self % right
```

Modulus of the values. Auto called when using the operator `%`.

#### RightMod(left)

```ruby
left % self
```

Modulus of the values. Auto called when using the operator `%`.

#### BitXor(right)

```ruby
self ^ right
```

Bitwise xor of the values. Auto called when using the operator `^`.

#### RightBitXor(left)

```ruby
left ^ self
```

Bitwise xor of the values. Auto called when using the operator `^`.

#### BitAnd(right)

```ruby
self & right
```

Bitwise and of the values. Auto called when using the operator `&`.

#### RightBitAnd(left)

```ruby
left & self
```

Bitwise and of the values. Auto called when using the operator `&`.

#### BitOr(right)

```ruby
self | right
```

Bitwise or of the values. Auto called when using the operator `|`.

#### RightBitOr(left)

```ruby
left | self
```

Bitwise or of the values. Auto called when using the operator `|`.

#### BitLeft(right)

```ruby
self << right
```

Bitwise left of the values. Auto called when using the operator `<<`.

#### RightBitLeft(left)

```ruby
left << self
```

Bitwise left of the values. Auto called when using the operator `<<`.

#### BitRight(right)

```ruby
self >> right
```

Bitwise right of the values. Auto called when using the operator `>>`.

#### RightBitRight(left)

```ruby
left >> self
```

Bitwise right of the values. Auto called when using the operator `>>`.

#### NegateBits

```ruby
~self
```

Negate bits of the value. Auto called when using the operator `~`.
