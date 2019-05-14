# fixed64
64-bit numeric type with 4-decimal fixed precision

This type was created to make working with financial (ie. currency) values in Go easier.
It provides methods to perform basic arithmetic, as well as conversion methods to/from
common types.

We can use float64, but floating point values are not precise for accounting purposes.
I.e. float64 may round decimals up or down, etc.

Fixed64 is implemented by scaling int64, to give it 4 decimal places.
The range is from -922.33 trillion to +922.33 trillion, which is sufficient
for normal accounting / financial records.
