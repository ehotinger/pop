# pop

[![Build Status](https://travis-ci.com/ehotinger/pop.svg?branch=master)](https://travis-ci.com/ehotinger/pop)

C# style expression parsing and evaluation in Go.

## Operator Precedence

|               Symbol                |        Operation        | Associativity |
|-------------------------------------|-------------------------|---------------|
| `[ ] ( ) . ++ --` (postfix)         | Expression              | Left to right |
| `& * + - ~ ! ++ --` (prefix)        | Unary                   | Right to left |
| `* / %`                             | Multiplicative          | Left to right |
| `+ -`                               | Additive                | Left to right |
| `<< >>`                             | Bitwise-shift           | Left to right |
| `< > <= >=`                         | Relational              | Left to right |
| `== !=`                             | Equality                | Left to right |
| `&`                                 | Bitwise-AND             | Left to right |
| `^`                                 | Bitwise-XOR             | Left to right |
| `\|`                                | Bitwise-OR              | Left to right |
| `&&`                                | Logical-AND             | Left to right |
| `\|\|`                              | Logical-OR              | Left to right |
| `? :` (Ternary)                     | Conditional             | Right to left |

- Operators are listed in order of highest to lowest precedence. Multiple symbols on the same line indicate equal precedence.