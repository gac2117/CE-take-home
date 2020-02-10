# Lesson Plan for Golang - Functions

This lesson assumes students already have some knowledge of Go basic data types and variables.

## Instructor Demo (5 mins)

#### Say:

> Functions are blocks of code or statements that give the user the ability to reuse the same code.
> These statements will perform a specific task and then return the results to the user.
> Sometimes a function may not return anything.

> To create a function, you must first declare it.

#### Write the following code on the board or on the screen shared with students:

```
func nameOfFunction(parameterName type)
```

#### As you say:

> To declare a function, use the keyword `func` and then the _function name_ followed by a set of parentheses.
> If the function requires any parameters, list them inside the parentheses by first typing the _name of the parameter_ followed by a space and then the _data type_.

> After the closing parentheses, if the function returns any values, you need to write the _data type of the return value_.

#### Add the word _type_ to the code already written:

```
func nameOfFunction(parameters type) type
```

#### Say:

> Then open a set of curly brackets. Inside the brackets, you will write the code for the function.

#### Add {} to the code:

```
func nameOfFunction(parameters type) type {

}
```

#### Say:

> If the function has a return value, make sure you use the keyword `return` and then the _value_.

#### Add the word _return_ to the code:

```
func nameOfFunction(parameters type) type {
  return value
}
```

#### Say:

> Let's do an example now.
> If I have a function called 'birthday' and it takes in a parameter of someone's name, then I would type

#### Write:

```
func birthday(name string)
```

#### Say:

> This function will return a string that will say "Happy Birthday" and whatever the name is. So after the parentheses I will type the word 'string' and then open a set of curly brackets.

#### Write:

```
func birthday(name string) string {

}
```

#### Say:

> Now to return a string along with the value of a parameter, you can concatenate it using the plus sign (+).
> So in our case we will type the keyword `return`, and then the phrase "Happy birthday" in quotations, and then the plus sign, and finally the parameter we are using, which is "name".

#### Write:

```
func birthday(name string) string {
  return "Happy birthday " + name
}
```

#### Say:

> We can also write functions to do math for us.
> Let's say we want to create a function that will do simple addition.

> First we will declare the function by using the keyword `func` and then the name of the function, "addThis". The function will take in two parameters which are numbers, so we will type

#### Write:

```
func addThis(num1, num2 int)
```

#### Say:

> This function will return the sum of these two numbers, so we need to add the return type and open up the curly brackets.

#### Write:

```
func addThis(num1, num2 int) int {

}
```

#### Say:

> We want to return the sum of _num1_ and _num2_, so we can return the sum like this

#### Write:

```
func addThis(num1, num2 int) int {
  return num1 + num2
}
```

#### Say:

> Now you are going to write a few functions on your own. Open up the assignment and you will see instructions on what the function will need to do and return. You have 15 minutes to complete the assignment.
