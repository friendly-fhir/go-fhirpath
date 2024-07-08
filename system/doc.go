/*
Package system contains the definitions of FHIRPath "system"-namespace types.
This includes all operations for parsing, converting to/from, and comparing
these types:
  - Boolean: A boolean type
  - Integer: A 32-bit integer type
  - Decimal: A decimal type with precision equivalent to an IEEE double-precision
    floating-point value.
  - Quantity: A numerical type that also contains a unit value
  - String: A type capable of holding a UTF-8 sequence of characters
  - Date: A representation of a (possibly partial) Date value
  - Time: A representation of a wall-clock time, disconnected from a date
  - DateTime: A (possibly partial) representation of a moment in time

Due to some of the weirder requirements of FHIRPath regarding comparison
semantics of certain types, some of these types define a `TryEqual` or `TryCompare`
instead of `Equal` or `Compare`, since in FHIRPath some operations may
optionally _not_ yield a result at all. This is all abstracted in the top-level
equality and comparison functions.
*/
package system
