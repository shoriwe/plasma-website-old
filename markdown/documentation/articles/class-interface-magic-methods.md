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

#### Mul

#### RightMul

#### Div

#### RightDiv

#### FloorDiv

#### RightFloorDiv

#### Pow

#### RightPow

#### Equals

#### RightEquals

#### NotEquals

#### RightNotEquals

#### GreaterThan

#### RightGreaterThan

#### LessThan

#### RightLessThan

#### GreaterThanOrEqual

#### RightGreaterThanOrEqual

#### LessThanOrEqual

#### RightLessThanOrEqual

## Integer specific

#### Mod

#### RightMod

#### BitXor

#### RightBitXor

#### BitAnd

#### RightBitAnd

#### BitOr

#### RightBitOr

#### BitLeft

#### RightBitLeft

#### BitRight

#### RightBitRight

#### NegateBits
