# Golox

Go implementation of Lox a Java interpreter from the book [Crafting Interpreters](https://craftinginterpreters.com/)


# TODO List 

## Scanning

- [x] Chapter 4.1 The Interpreter Framework
- [x] Chapter 4.2 Lexemes and Tokens
- [x] Chapter 4.3 Regular Languages and Expressions
- [x] Chapter 4.4 The Scanner Class
- [x] Chapter 4.5 Recognizing Lexemes
- [x] Chapter 4.6 Longer Lexemes
- [x] Chapter 4.7 Reserved Words and Identifiers
- [x] Challange 4 Block comments

## Representing Code

- [x] Chapter 5.1 Context-Free Grammars
- [x] Chapter 5.2 Implementing Syntax Trees
- [x] Chapter 5.3 Working with Trees
- [x] Chapter 5.4 A (Not Very) Pretty Printer

## Parsing Expressions

- [x] Chapter 6.1 Ambiguity and the Parsing Game
- [x] Chapter 6.2 Recursive Descent Parsing
- [x] Chapter 6.3 Syntax Errors
- [x] Chapter 6.4 Wiring up the Parser

## Evaluating Expressions

- [x] Chapter 7.1 Representing Values
- [x] Chapter 7.2 Evaluating Expressions
- [x] Chapter 7.3 Runtime Errors
- [x] Chapter 7.4 Hooking Up the Interpreter

## Statements and State

- [x] Chapter 8.1 Statements
- [x] Chapter 8.2 Global Variables
- [x] Chapter 8.3 Environments
- [x] Chapter 8.4 Assignment
- [x] Chapter 8.5 Scope

## Control Flow

- [x] Chapter 9.1 Turing Machines (Briefly)
- [x] Chapter 9.2 Conditional Execution
- [x] Chapter 9.3 Logical Operators
- [x] Chapter 9.4 While Loops
- [x] Chapter 9.5 For Loops

## Functions

- [x] Chapter 10.1 Function Calls
- [x] Chapter 10.2 Native Functions
- [x] Chapter 10.3 Function Declarations
- [x] Chapter 10.4 Function Objects
- [x] Chapter 10.5 Return Statements
- [ ] Chapter 10.6 Local Functions and Closures

```
For the return statements chapter, it works fine until we try to make two recursive calls
inside the return statement:
fun fib(n) {
  if (n <= 1) return n;
  print n;
  return fib(n-1) + fib(n-2);
}
```


## Resolving and Binding

- [ ] Chapter 11.1 Static Scope
- [ ] Chapter 11.2 Semantic Analysis
- [ ] Chapter 11.3 A Resolver Class
- [ ] Chapter 11.4 Interpreting Resolved Variables
- [ ] Chapter 11.5 Resolution Errors


## Classes

- [ ] Chapter 12.1 OOP and Classes
- [ ] Chapter 12.2 Class Declarations
- [ ] Chapter 12.3 Creating Instances
- [ ] Chapter 12.4 Properties on Instances
- [ ] Chapter 12.5 Methods on Classes
- [ ] Chapter 12.6 This
- [ ] Chapter 12.7 Constructors and Initializers


## Inheritance

- [ ] Chapter 13.1 Superclasses and Subclasses
- [ ] Chapter 13.2 Inheriting Methods
- [ ] Chapter 13.3 Calling Superclass Methods
- [ ] Chapter 13.4 Conclusion
