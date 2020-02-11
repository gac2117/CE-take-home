# Lesson Plan for Golang - Functions

This lesson assumes students already have some knowledge of Go basic data types and variables.

## Materials

- Lesson plan
- Text editor that is projected on a screen
- In-class example
- Coding activity

## Instructor Demo (5 mins)

The lesson plan format is divided into two parts: speaking and coding. For some parts, it is advised that the instructor will code as s/he is speaking.

#### Say:

> Functions are blocks of code or statements that give the user the ability to reuse the same code.
> These statements will perform a specific task and then return the results to the user.
> Sometimes a function may not return anything.

> To create a function, you must first declare it.

#### Type the following code in the text editor:

```golang
func function_name(parameterName type)
```

#### As you say:

> To declare a function, use the keyword `func` and then the _function name_ followed by a set of parentheses.
> If the function requires any parameters, list them inside the parentheses by first typing the _name of the parameter_ and then the _data type_ of the parameter.

> After the closing parentheses, if the function returns any values, you need to write the _data type of the return value_.

#### Add the word _type_ to the code already written:

```golang
func nameOfFunction(parameters type) type
```

#### Say:

> Then open a set of curly brackets. Inside the brackets, you will write the code for the function.

#### Add {} to the code:

```golang
func nameOfFunction(parameters type) type {
  // function body code
}
```

#### Say:

> If the function has a return value, make sure you use the keyword `return` and then the _value_ that is being returned.

#### Add the word _return_ to the code:

```golang
func nameOfFunction(parameters type) type {
  return value
}
```

#### Say:

> Let's do an example now.
> If I have a function called 'birthday' and it takes in a parameter of someone's name which is a string, then I would type

#### Type:

```golang
func birthday(name string)
```

#### Say:

> This function will return a string that will say "Happy Birthday" and whatever the name is. So after the parentheses I will type the word 'string' and then open a set of curly brackets.

#### Add on to the code:

```golang
func birthday(name string) string {

}
```

#### Say:

> Now to return a string along with the value of a parameter, you can concatenate it using the plus sign (+).
> So in our case, we will type the keyword `return`, and then the phrase "Happy birthday" in quotations including a space so that we have some space between the words 'birthday' and the person's name, and then the plus sign, and finally the parameter we are using, which is "name".

#### As you type:

```golang
func birthday(name string) string {
  return "Happy birthday " + name
}
```

#### Say:

> Now in the main function, we will print out on to the screen what this 'birthday' function returns. So we will write

#### Type:

```golang
func main() {
	fmt.Println(birthday("John"))
}
```

#### Execute the code so students can see "Happy Birthday John" printed in the console.

#### Say:

> We can also write functions to do math for us.
> Let's say we want to create a function that will do simple addition.

> First we will declare the function by using the keyword `func` and then the name of the function, "addThis". The function will take in two parameters which are numbers, so we will type

#### As you type:

```golang
func addThis(num1, num2 int)
```

#### Say:

> This function will return the sum of these two numbers, so we need to add the return type of an integer and then open up the curly brackets.

#### Add on to the code:

```golang
func addThis(num1, num2 int) int {

}
```

#### Say:

> We want to return the sum of _num1_ and _num2_, so we can type _num1_ and then the plus sign and then _num2_.

#### As you type:

```golang
func addThis(num1, num2 int) int {
  return num1 + num2
}
```

#### Say:

> In order to execute this code, we need to add it to the main function, so let's go ahead and do that

#### Add on to the code:

```golang
func main() {
	fmt.Println(birthday("John"))
  fmt.Println(addThis(3, 6))
}
```

#### Execute the code so students can see the sum, 9, printed in the console.

#### Say:

> Now you are going to write a few functions on your own. Open up the assignment and you will see instructions on what the functions will need to do and return. You have 15 minutes to complete the assignment.
