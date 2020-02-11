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
  fmt.Println(addThis(3, 6))
}
```

#### Execute the code so students can see the sum, 9, printed in the console.

#### Say:

> Now you are going to write a few functions on your own. Open up the assignment and you will see instructions on what the functions will need to do and return. You have 15 minutes to complete the assignment.

As the students are coding, walk around the classroom to make sure no one is stuck.

## Instructor Review (10 mins)

#### Say:

> That was a simple assignment. How did you do? Would anyone like to share their solution code?

See if there are any volunteers. Have the student read their code aloud as you code it into your text editor for everyone to see. Execute the code to see if it works. Thank the student who volunteered.

#### Say:

> Let's go over the solution code together. First of all, we have the function called 'greeting' that accepts 3 parameters. You can name the parameters whatever you want but make sure they are strings.

#### Type:

```golang
func greeting(name string, city string, job string)
```

#### Say:

> While it is perfectly fine to write the parameters like this, I will show you a cleaner way to write them.

#### Type:

```golang
func greeting(name, city, job string)
```

#### Say:

> As you can see, you can list the parameters separated by commas as long as they are of the same type.
> Next we need to specify that the return value will be a string.

#### Type:

```golang
func greeting(name, city, job string) string
```

#### Say:

> Open the curly brackets and return the concatenated string. Make sure to have spaces in between the words!

#### Type:

```golang
func greeting(name, city, job string) string {
	return "Hello, my name is " + name + ". I live in " + city + " and I am a " + job + "."
}
```

#### Say:

> For the next function, you had to use an integer converter to change the number into a string. But first, we need to declare the function

#### Type:

```golang
func age(birthyear int) string {

}
```

#### Say:

> There are different ways you can do this, but the simplest way is to declare the variable 'age' and set it to the difference between the birth year and the current year. So that can be done with simple subtraction

#### Type:

```golang
age := 2020 - birthyear
```

#### Say:

> However, if we do this, the age is an integer and we cannot print an integer in a string. So we need to convert this integer into a string, using `strconv.Itoa()`

#### Type:

```golang
age := strconv.Itoa(2020-birthyear)
```

#### Say:

> Finally we need to return the string that includes this age.

#### Type:

```golang
return "I am " + age + " years old."
```

#### Say:

> Again you need to make sure that the string is formatted correctly with all the spaces in between the words. So the final function looks like this.

#### Type:

```golang
func age(birthyear int) string {
	age := strconv.Itoa(2020 - birthyear)
	return "I am " + age + " years old."
}
```

#### Say:

> Now to print everything out, we need to call these functions in the main function. Remember we had to use different parameters for each of the three times we call these functions.

#### Type:

```golang
func main() {
	fmt.Println(greeting("John", "New York City", "banker"))
	fmt.Println(age(1980))
	fmt.Println(greeting("Sam", "Miami", "surfer"))
	fmt.Println(age(1991))
	fmt.Println(greeting("Jane", "San Francisco", "software developer"))
	fmt.Println(age(1986))
}
```

#### Execute the code so students can see results printed in the console.

#### Say:

> The results look like three different people giving their self-introductions!
> As you can see, using functions will help us do the same task multiple times easily. We will continue to use functions all throughout this course. They will become more and more complex but this is a good start.
