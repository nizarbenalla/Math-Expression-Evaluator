# Math expression evaluator

## Program logic and limitations

- The main algorithms used in the program are the [Shunting yard](https://en.wikipedia.org/wiki/Shunting_yard_algorithm) and [Reverse Polish Notation](https://en.wikipedia.org/wiki/Reverse_Polish_notation)/[AST](https://en.wikipedia.org/wiki/Abstract_syntax_tree) evaluation 
to transform the mathematical infix expression into a postfix notation, making easier for computers to work with
- I split the expression given by the user into many tokens - pass those to my infix to postfix function, then evaluate the AST
### Limitations
- one limitation is the need for spaces between numbers, functions and operators in your expression (see example below)
### Dependencies
- the `github.com/golang-collections/collection` repo is used for the Stack and Queue Data Structure

## How to Build the program
#### clone the repo

```bash
git clone https://github.com/nizarbenalla/Math-Expression-Evaluator.git 
````
#### cd into the repo

```bash
cd MathExpressionEvaluator 
```
#### build the program

```bash
go build
```

## Running the Program
<h4> the `-math` option is necessary</h4>

```bash 
./MathExpressionEvaluator_ -math <expression>
```

## Examples
```
./MathExpressionEvaluator_ -math "( cos ( 1 ) + 3 ) * ( 10 - 1 + 1 ) "
Your answer is :
35.4030230586814

./MathExpressionEvaluator_ -math "( 1 - 8 ) + ( 5 / 2 )"
Enter the expression you want to compute
-4.5

./MathExpressionEvaluator_ -math " cos ( 1 ) "
Your answer is :
0.5403023058681397
```
