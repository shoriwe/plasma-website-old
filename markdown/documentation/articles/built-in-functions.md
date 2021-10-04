# Functions

## `expand(RECEIVER, SOURCE) -> None`

Assigns all the symbols and their values of the **`SOURCE`** object to the **`Receiver`**.

## `dir(OBJECT) -> Array[String]`

List all the symbols of the provided **`OBJECT`**

## `set(OBJECT, SYMBOL, VALUE) -> None`

Sets to the **`OBJECT`**.**`SYMBOL`** the **`VALUE`** provided.

## `get_from(SOURCE, SYMBOL) -> Object`

Tries to get the **`SYMBOL`** from **`SOURCE`** object.

## `delete_from(OBJECT, SYMBOL) -> None`

Tries to delete the **`SYMBOL`** from **`OBJECT`** provided.

## `input(OBJECT) -> String`

First prints the supplied **`OBJECT`** to the stdout, then reads a line from the stdin and return it.

## `print(OBJECT) -> None`

Prints the **`OBJECT`** to the stdout.

## `println(OBJECT) -> None`

Prints the **`OBJECT`** followed by a new line to the stdout.

## `id(OBJECT) -> Integer`

Returns the unique id of the provided **`OBJECT`**.

## `range(START, END, STEP) -> Generator[Integer]`

Creates a generator that starts with the value of **`START`** and finish when the counter reaches **`END`** or is
greater than it, Every step is done by adding the counter with the **`STEP`**.

## `super(SELF, CLASS) -> Object`

Returns a proxy object to operate **`SELF`** as it was and instance of **`CLASS`**.

## `map(FUNCTION, ITERABLE) -> Generator[Any]`

Pass every element from **`ITERABLE`** to **`FUNCTION`** as its argument, returns a generator of the results of the
function calls.

## `filter(FUNCTION, ITERABLE) -> Generator[Any]`

Pass every element from **`ITERABLE`** to **`FUNCTION`**  as its argument, if the output of **`FUNCTION`** is **True**,
returns a generator of those values that produced those **`True`** results.