# Lesson Plan for Golang - Conditionals

This lesson assumes students already have some knowledge of Go basic data types, variables, functions.

## Materials

- Lesson plan
- Text editor that is projected on a screen
- In-class example
- Coding activity

## Instructor Demo (5 mins)

The lesson plan format is divided into two parts: speaking and coding. For some parts, it is advised that the instructor will code as s/he is speaking.

#### Say:

> Conditional statements tell a program to execute different actions depending on whether a condition is true or false.
> For example, if x is equal to 5 then do this. If not, then do something else.
> That is a simple if/else statement.

> Let's put that into coding language.

#### Type the following code in the text editor:

```golang
x := 5

if x == 5 {
  fmt.Println("x is equal to 5")
} else {
  fmt.Println("x is not equal to 5")
}
```

#### As you say:

> Let's set a variable called x to the number 5.
> To start the if/else statement, we use the keyword `if`. Then we write the expression that we want to evaluate as either true or false. So in our case, we want to know if x is equal to 5.

> If it is true, then inside the curly brackets, write what should happen. For us, we want to print out the string "x is equal to 5".

> If the expression is false, meaning x does not equal 5, then we use the keyword `else`. Then in the curly brackets we execute a different action. For us, we want to print out the string "x is not equal to 5".

> Now let's execute this inside the main function

#### Type:

```golang
func main() {
  x := 5

  if x == 5 {
    fmt.Println("x is equal to 5")
  } else {
    fmt.Println("x is not equal to 5")
  }
}
```

#### Execute the code so students can see "x is equal to 5" printed in the console.

#### Say:

> Now let's change x to a different number, like 7. Then since x is not equal to 5, it should execute that else statement

#### Change the value of x to 7 and execute the code so students can see "x is not equal to 5" printed in the console.

#### Say:

> We can add as many conditional statements as we want. All we have to do is use the keyword `else if`.

> For example, let's say we want to check if x is greater than 10, less than 10, or equal to 10.

#### Type:

```golang
if x > 10 {
  fmt.Println("x is greater than 10")
} else if x < 10 {
  fmt.Println("x is less than 10")
} else if x == 10 {
  fmt.Println("x is equal to 10")
}
```

#### Say:

> Let's break this down. The first expression is checking if x is greater than 10. If true, then it will print out "x is greater than 10". If false, it will go to the next expression, which is checking if x is less than 10. Again, if true, then it will print out that statement, but if false, it will go to the last expression. The last expression is checking whether x is equal to 10. Let's execute this code and see what we get if we set x to equal 15.

#### Execute the code with x assigned to 15 so students can see "x is greater than 10" printed in the console.

#### Say:

> Okay now let's try a number that is less than 10, like 5.

#### Change x to 5 and execute the code so it prints out "x is less than 10".

#### Say:

> Now let's set x to equal 10

#### Change x to 10 and execute the code so it prints out "x is equal to 10".

#### Say:

> We could've used an `else` statement too because if x is not greater or less than 10, then it must equal 10

#### Modify the code:

```golang
if x > 10 {
  fmt.Println("x is greater than 10")
} else if x < 10 {
  fmt.Println("x is less than 10")
} else {
  fmt.Println("x is equal to 10")
}
```

#### Execute the code again with x still assigned to 10 so it prints out "x is equal to 10".

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
