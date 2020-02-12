# Lesson Plan for Golang - Conditionals

This lesson assumes students already have some knowledge of Go basics, such as basic data types, variables, functions, and logical operators.

## Materials

- Lesson plan
- Text editor that is projected on a screen
- In-class examples
- Coding activities

## Instructor Demo (5 mins)

The lesson plan format is divided into two parts: speaking and coding. For some parts, it is advised that the instructor will code as s/he is speaking.

### **If/Else Statements**

#### Say:

> Conditional statements tell a program to execute different actions depending on whether a condition is true or false.
> For example, if x is equal to 5 then do this. If not, then do something else.
> That is a simple if/else statement.

> Let's put that into coding language.

### **In-class Example**

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

> If it is true, then inside the curly brackets, we will write what should happen. For us, we want to print out the string "x is equal to 5".

> If the expression is false, meaning x does not equal 5, then we use the keyword `else`. Then in the curly brackets we execute a different action. In our case, we want to print out the string "x is not equal to 5".

> Now let's execute this inside the main function

#### Add on to the code:

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

#### Execute the code so students can see the output

The output should be "x is equal to 5"

#### Say:

> Now let's change x to a different number, like 7. Then since x is not equal to 5, it should execute that else statement

#### Change the value of x to 7:

```golang
func main() {
  x := 7

  if x == 5 {
    fmt.Println("x is equal to 5")
  } else {
    fmt.Println("x is not equal to 5")
  }
}
```

#### Execute the code so students can see the output

The output should be "x is not equal to 5"

### **Else if statements**

#### Say:

> We can add as many conditional statements as we want. All we have to do is use the keyword `else if`.

### **In-class Example**

#### Say:

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

#### As you say:

> The first expression is checking if x is greater than 10. If true, then it will print out "x is greater than 10".

> If false, it will go to the next expression, which is checking if x is less than 10. Again, if true, then it will print out that statement, but if false, it will go to the last expression.

> The last expression is checking whether x is equal to 10.

> Let's execute this code and see what we get if we set x to equal 15.

#### Add on to the code:

```golang
func main() {
  x := 15

	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x < 10 {
		fmt.Println("x is less than 10")
	} else if x == 10 {
		fmt.Println("x is equal to 10")
	}
}
```

#### Execute the code so students can see the output

The output should be "x is greater than 10"

#### Say:

> Okay now let's try a number that is less than 10, like 5.

#### Change the value of x to 5:

```golang
func main() {
  x := 5

	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x < 10 {
		fmt.Println("x is less than 10")
	} else if x == 10 {
		fmt.Println("x is equal to 10")
	}
}
```

#### Execute the code so students can see the output

The output should be "x is less than 10"

#### Say:

> Now let's set x to equal 10

#### Change the value of x to 10:

```golang
func main() {
  x := 10

	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x < 10 {
		fmt.Println("x is less than 10")
	} else if x == 10 {
		fmt.Println("x is equal to 10")
	}
}
```

#### Execute the code so students can see the output

The output should be "x is equal to 10"

#### Say:

> We could've used an `else` statement too because if x is not greater or less than 10, then it must equal 10

#### Change the last `else if` to `else`:

```golang
if x > 10 {
  fmt.Println("x is greater than 10")
} else if x < 10 {
  fmt.Println("x is less than 10")
} else {
  fmt.Println("x is equal to 10")
}
```

#### Execute the code so students can see the output

The output should still be "x is equal to 10"

### **Switch statements**

#### Say:

> But for complex conditionals, it is cleaner to use switch statements. A switch statement will run the first case equal to the condition expression. The cases are evaluated from the top down, and will stop once a case is executed. If none of the cases match, there is a default case that will be executed instead.

### **In-class Example**

#### Say:

> Let's write a switch statement now. Let's say a variable called 'month' is set to the numeric representation of a month. The switch statement will then print out the name of the month.

> To start, we will use the keyword `switch` and then the expression that will be used against the cases. So in our case, it is the month.

#### Type:

```golang
  switch month {

  }
```

#### Say:

> To create the different cases, we use the keyword `case` and then the conditional that needs to be evaluated as true.

#### Add on to the code:

```golang
	switch month {
	case 1:

	case 2:

	case 3:

	case 4:

	case 5:

	case 6:

	case 7:

	case 8:

	case 9:

	case 10:

	case 11:

	case 12:

	default:

	}
```

#### Say:

> So in our case, we will need 12 cases to check all the months of the year. And then we need a default case just in case the month variable is set to a number that is not a month, like 0 or 13.

> Now what do we want do if one of these cases are true? We want to print out the name of the month. So in each case, we need to code that.

#### Add on to the code:

```golang
	switch month {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("May")
	case 6:
		fmt.Println("June")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("August")
	case 9:
		fmt.Println("September")
	case 10:
		fmt.Println("October")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("December")
	default:

	}
```

#### Say:

> Now for the default case, if the month variable is set to a number that is not 1-12, then we can print out a message saying "This is not a month".

#### Add on to the code:

```golang
	default:
		fmt.Println("This is not a month")
```

#### Say:

> Let's now declare a variable at the top called month and set it equal to 5. And let's execute the code to see what it prints out.

#### Add on to the code:

```golang
  month := 5
```

#### Execute the code so students can see the output.

The output should be "May"

#### Change the value of month and execute the code again so it prints out a different output.

#### Say:

> If we change the value of month to a number other than 1-12, it should hit the default case and print out "This is not a month". So let's try that now.

#### Change the value of month to 13:

```golang
  month := 13
```

#### Execute the code so students can see the output

The output should be "This is not a month"

### **Coding Exercises**

#### Say:

> Now you are going to write a few conditional statements on your own. Open up the assignment and you will see instructions on what kind of conditionals you will need to write. You have 20 minutes to complete the assignment. If you get stuck, please raise your hand and I will go help you.

As the students are coding, walk around the classroom to make sure no one is struggling.

## Instructor Review (10 mins)

#### Say:

> How did you do? Would anyone like to share their solution code?

See if there are any volunteers. Have the student read his/her code aloud as you code it into your text editor for everyone to see. Execute the code to see if it works. Thank the student who volunteered.

#### Say:

> For these exercises, there could be multiple ways to express the condition, so if you have something that looks different from the way I'm coding it here, please share!

### **Even/Odd Number Exercise**

> For the first exercise, we need to determine whether the number is an even number or odd number. One way we can do this is by dividing it by 2 and checking if there is any remainder. If there isn't, then it is an even number.

#### Type:

```golang
	x := 51

	if x%2 == 0 {
		fmt.Println("This is an even number")
	} else {
		fmt.Println("This is an odd number")
	}
```

#### Say:

> Your variable doesn't have to be x; it can be anything.

> The way we can check for remainders is by using the modulus operator %, which is the percent sign.

> So here we are checking if x has any remainder when it is divided by 2. If there isn't any, then it prints out "This is an even number". Otherwise we can assume it is an odd number. You can test it out by plugging in other numbers for x.

### **Execute the Coding Exercise**

Change the value of x to an even number. The output should be "This is an even number".

Change the value of x to an odd number. The output should be "This is an odd number".

### **24-hour Clock Exercise**

#### Say:

> Next we worked on the 24-hour clock. Again you can call your variable whatever you'd like. I called it 'hour'.

> The first condition we want to check is whether the hour is between 6 and 12. In order to do this, we would need to use Logical AND Operator, which is represented by two ampersands &&, as well as the math symbols for greater than >, lesser than <, and equal to =.

#### As you type:

```golang
	hour := 7

	if hour >= 6 && hour <= 12 {
		fmt.Println("Good morning!")
	}
```

#### Say:

> I have my variable 'hour' set to 7, and my 'if statement' is checking whether 'hour' is greater than or equal to 6 as well as less than or equal to 12.

> Since my 'hour' variable is set at 7 right now, let's execute this to see if it will print out 'Good morning!"

### **Execute the Coding Exercise**

The output should be "Good morning!"

#### Say:

> Perfect! So next we want to check for another condition, so we need to use the keywords `else if` and then write the condition we are checking for. In this case, let's check for the afternoon hours.

#### Add on to the code:

```golang
	hour := 15

	if hour >= 6 && hour <= 12 {
		fmt.Println("Good morning!")
	} else if hour >= 13 && hour <= 18 {
		fmt.Println("Good afternoon!")
	}
```

#### Say:

> Here we are checking, if 'hour' wasn't between 6 and 12, then let's check if it's greater than or equal to 13 and less than or equal to 18. If it is, then it should print out "Good afternoon!"

> Let's change the 'hour' variable to 15 and execute the code.

### **Execute the Coding Exercise**

The output should be "Good afternoon!"

#### Say:

> Good! It's working so far. Now one way we can check for the other two conditions is by choosing which one is easier to check for. I think it is easier to check whether the 'hour' is invalid, meaning it is greater than 24. So that's what my next `else if` expression is checking for.

#### Add on to the code:

```golang
	hour := 30

	if hour >= 6 && hour <= 12 {
		fmt.Println("Good morning!")
	} else if hour >= 13 && hour <= 18 {
		fmt.Println("Good afternoon!")
	} else if hour >= 24 {
		fmt.Println("Please use a number between 0 and 23")
	}
```

#### Say:

> Let's go ahead an change the 'hour' variable to a bigger number, like 30, and execute the code.

### **Execute the Coding Exercise**

The output should be "Please use a number between 0 and 23

#### Say:

> Alright, finally we will check for the last condition, which is whether the hour is between 0 and 5 or 19 and 23. That is hard to code using logical operators, so let's just use the `else` statement for that.

#### Add on to the code:

```golang
	hour := 20

	if hour >= 6 && hour <= 12 {
		fmt.Println("Good morning!")
	} else if hour >= 13 && hour <= 18 {
		fmt.Println("Good afternoon!")
	} else if hour >= 24 {
		fmt.Println("Please use a number between 0 and 23")
	} else {
		fmt.Println("Good night!")
	}
```

#### Say:

> Let's test this out and change the 'hour' variable to 20 and execute the code.

### **Execute the Coding Exercise**

The output should be "Good night!"

#### Say:

> Like I said before, there are other ways you can code this and still have it run perfectly fine. Does anyone have a different way they coded it?

See if there are any volunteers. Have them share their code and see different ways of making this work.

#### Say:

> See? This is the beauty of coding. There are many ways to do the same thing. You can be creative and go about it in different ways. However, one thing we want to strive for is to be concise.

> Have you heard of DRY? It stands for "Don't Repeat Yourself" and basically you want to code so that you're not being repetitive and unnecessarily long in your coding.

### **World Greeting Exercise**

#### Say:

> Anyway for the last exercise, it required a little bit of outside knowledge regarding world languages. If you couldn't think of any off the top of your head, you were welcome to use Google Translate.

> First we need to declare a variable, which I called 'language', and set it to any language. I set mine to 'French'.

#### Type:

```golang
language := "French"
```

#### Say:

> Next we need to set up the switch statement. So first we use the keyword `switch` and then whatever we are checking for, which is the language.

#### Add on to the code:

```golang
switch language {

}
```

#### Say:

> Next we build our cases. For the first case, I'll just use what my variable is set at right now, which is French. I know that the French word for hello is "Bonjour".

#### Add on to the code:

```golang
switch language {
	case "French":
    fmt.Println("Bonjour!")
}
```

#### Say:

> Now I just need to add two more languages, so I can copy and paste that section down, and change the language and greetings.

#### Add on to the code:

```golang
switch language {
	case "French":
		fmt.Println("Bonjour!")
	case "Korean":
		fmt.Println("An-nyoung!")
	case "Japanese":
    fmt.Println("Konnichiwa!")
}
```

#### Say:

> Perfect! Finally we need to have a default case, just in case the language isn't one of the three you chose. So let's add the default case and have it return "I don't recognize that language".

#### Add on to the code:

```golang
	switch language {
	case "French":
		fmt.Println("Bonjour!")
	case "Korean":
		fmt.Println("An-nyoung!")
	case "Japanese":
		fmt.Println("Konnichiwa!")
	default:
		fmt.Println("I don't recognize that language")
	}
```

#### Say:

> Alright now let's test it out by changing the language to "Korean"

### **Execute the Coding Exercise**

The output should be "An-nyoung!"

#### Say:

> Let's try a language that isn't one of the three cases, like "German".

### **Execute the Coding Exercise**

The output should be "I don't recognize that language"

#### Say:

> Does anyone have any questions?

### **Conclusion**

#### Say:

> Conditional statements are crucial in programming and is used in any programming language you may learn in the future. What we learned today was just the basics. These can become very complex, but as long as you can understand the logical flow of information, you will be able to use it in your code.
