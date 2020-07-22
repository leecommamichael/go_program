# Evils of test harness
```go
package
import
() declaration block
""
func
*
() Parameter List
() Function Call
{} Function Body
```

# Zero
`//`
    These two characters form a token called a "comment". It isn't code, just a note passed from one programmer to the next.
`||`
    This is a token called "or", you need to write the pipe `|` twice. `|` is not the same as `||`. There is a boolean on both sides of "or". "or" will turn those two booleans into one. Here are the four possibilities.
    `true || true` evaluates to `true`
    `false || false` evaluates to `false`
    `true || false` evaluates to `true`
    `false || true` evaluates to `true`
`&&`
    These two characters form a token called "and", here's how it works.
    `true && true` evaluates to `true`
    `false && false` evaluates to `false`
    `true && false` evaluates to `false`
    `false && true` evaluates to `false`
`true` and `false`
    These are values. The two existing values of the boolean data type. Languages often call this type `bool`.
`!`
    This is character is a token called "not". It reverses the boolean it is to the left of. 
    `!true` evaluates to `false`
    `!false` evaluates to `true`
    `!!!!!!!!!!!!!!!!!!!!!true` evaluates to `false`... dont write that...
`if true {} else {}`
    This is an if statement. A boolean is written between the `if` and `{}` which makes a decision of whether to run the code in that `{}` or skip to the code in the `else` block `{}`.
    
    `if` is a token that always starts an "if statement". If-statements need to provide a boolean and then `{}` braces in order for the program to be valid (runnable by your computer).

    In code, braces always appear in pairs. When you see `{` and a bunch of lines, assume there is a closing `}` brace below. We call this a "block" of code.
`()`
    Parenthesis often called just "parens" control how expressions are evaluated.
    `false || true && true` evaluates to `true`
    `false || (true && true)` first evaluates the parenthesis as shown below.
    `false || true` then evaluates to `true`
# One
Numbers
`==`
`!=`
`>`
`-number`
`<=`
# Two
`""`
`"\n"`
`'J'`
    Runes/Characters
`[0]`
    Indexing
# Three
`[0]type{}`
    Sized Array Initializers
`[]type{}`
    Unsized Array Initializers
# Four
`func`
`int, float32, []int`
`return`
- Declaration vs Invocation
