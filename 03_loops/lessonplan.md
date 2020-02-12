# Lesson Plan for Golang - Loops

This lesson assumes students already have some knowledge of Go basics, such as basic data types, variables, functions, logical operators, conditionals.

## Materials

- Lesson plan
- Text editor that is projected on a screen
- In-class examples
- Coding activities

## Instructor Demo (5 mins)

The lesson plan format is divided into two parts: speaking and coding. For some parts, it is advised that the instructor will code as s/he is speaking.

### **For Loop**

#### Say:

> There are many types of loops in programming languages, but in Golang there is only one: the `for` loop. Basically, the `for` loop will repeatedly execute code inside the block for a specific number of times.

### **Syntax of `for` loop**

#### Type the following code in the text editor:

```golang
for initialization; condition; post{

}
```

#### As you say:

> We start by using the keyword `for` and then the initialization statement. This is a simple statement, like variable declarations. It is executed before the first loop or iteration.

> Next is the condition statement. This is a boolean expression that is executed before every iteration. It is checking whether we should run the loop again or not. If the boolean expression is true, then the loop executes. If it is false, then it stops.

> Finally there is the post statement. This is executed after every iteration.

> We separate these statements using a semi-colon (;). Let's build a simple `for` loop now.

### **In-class Example**

#### Change the initialization statement to:

```golang
for i := 0; condition; post{

}
```

#### As you say:

> For the initialization statement, we are going to declare a variable 'i' and set it to 0. That will serve as our counter.

#### Change the condition statement to:

```golang
for i := 0; i < 5; post{

}
```

#### As you say:

> For the condition statement, we are going to run the loop as long as 'i' is less than 5. So basically it will run 5 times.

#### Change the post statement to:

```golang
for i := 0; i < 5; i++{

}
```

#### Say:

> For the post statement, we are going to increment 'i' by 1 at the end of every loop.

> Now, for the body of the loop, what do we want to do 5 times? Let's say we want to print out "Coding is fun!".

#### Add on to the code:

```golang
for i := 0; i < 5; post{
  fmt.Println("Coding is fun!")
}
```

#### Say:

> Let's execute this code to see what happens.

#### Execute the code so the students can see the output

The output should be:

Coding is fun!

Coding is fun!

Coding is fun!

Coding is fun!

Coding is fun!

#### Say:

> Okay well let's try to see exactly how this is working. Instead of printing "Coding is fun" 5 times, let's print what the variable 'i' is in every loop.

#### Change the code to:

```golang
for i := 0; i < 5; post{
  fmt.Println(i)
}
```

#### Execute the code so students can see the output

The output should be:

0

1

2

3

4

#### Say:

> Notice what the loop is doing. First, before the loop executes, 'i' is 0. After the loop executes, 'i' is incremented by 1 so it is 1. That is still less than 5, so the loop executes and prints out 1. Then it is incremented by 1 again, making it 2. That is still less than 5, so the loop executes.

> This keeps repeating until 'i' is no longer less than 5. So that's why it ends with printing out 4.

### **Coding Exercises**

#### Say:

> Now you are going to write a couple of `for` loops on your own. One of them is a variation of a very popular coding challenge called "FizzBuzz".

> Open up the assignment and you will see instructions on what kind of `for` loop you will need to write. You have 20 minutes to complete the assignment. If you get stuck, please raise your hand and I will go help you.

As the students are coding, walk around the classroom to make sure no one is struggling.

## Instructor Review (10 mins)

#### Say:

> How was it? I know the coding exercises are becoming more challenging as we progress, but that shows how much you've learned so far and what you're capable of doing now!

> Would anyone like to share their code for these exercises?

See if there are any volunteers. Have the student read his/her code aloud as you code it into your text editor for everyone to see. Execute the code to see if it works. Thank the student who volunteered.

### **Sum Exercise**

#### Say:

> For the first exercise, the hint tells us to declare a variable called 'sum' outside of the loop. So let's do that first.

#### Type:

```golang
sum := 0

fmt.Println(sum)
```

#### Say:

> We will set the sum to zero because nothing has been added yet. Now we want to add the numbers starting with 15. So we will set the initialization statement to `i := 15` in our `for` loop.

#### Add on to the code:

```golang
sum := 0

for i := 15

fmt.Println(sum)
```

#### Say:

> Next we want to keep looping until we hit 25. That means 'i' needs to be less than 26.

#### Add on to the code:

```golang
sum := 0

for i := 15; i < 26

fmt.Println(sum)
```

#### Say:

> So as long as 'i' is less than 26, this loop will continue to run. Finally, after every iteration we will increment 'i' by 1.

#### Add on to the code:

```golang
sum := 0

for i := 15; i < 26; i++ {

}

fmt.Println(sum)
```

#### Say:

> Inside the loop, we want to add all the numbers in order. So in the first iteration, 'i' is equal to 15. We will add that to the variable 'sum' which will keep track of our sum.

> In the second iteration, 'i' is equal to 16. When we add that to 'sum' it will become 15 + 16, which is 31. The way we code that is this:

#### Add on to the code:

```golang
sum := 0

for i := 15; i < 26; i++ {
  sum += i
}

fmt.Println(sum)
```

#### Say:

> Basically it will continue to add all the numbers until we hit 26. We will NOT add 26 to the sum. Let's see what the sum of all the numbers between 15 and 25 are.

### **Execute the Coding Exercise**

The output should be 220.

### **FizzBuzz Exercise**

#### Say:

> The next exercise was a variation of a popular coding challenge that is actually given in technical interviews. This is an easier version of it. Usually it involves printing out numbers 1 to 100. Every number divisible by 3 will print out "Fizz", every number divisible by 5 will print out "Buzz", and every number divisible by 15 will print out "FizzBuzz".

> But for our exercise, we only have to print out "FizzBuzz" if the number was divisible by 4.
> First, let's set up our `for` loop.

#### Type:

```golang
for i := 1; i<=30; i++ {

}
```

#### As you say:

> We will set 'i' to 1 so it starts printing out the numbers starting with 1. We will keep running the loop until it equals 30. And we will increment 'i' by 1 every time the loop ends.

> Now, the hint tells us we can use an if/else statement inside of the `for` loop, so let's go ahead and set that up.

#### Add on to the code:

```golang
for i := 1; i<=30; i++ {
  if {

  } else {

  }
}
```

#### Say:

> In order to know if a number is cleanly divisible by another number, the remainder must be 0. We check for this using the modulus operator %.

#### Add on to the code:

```golang
for i := 1; i<=30; i++ {
  if i%4 == 0{

  } else {

  }
}
```

#### Say:

> We are saying, if we divide 'i' by 4 and there is no remainder, then we know this is cleanly divisible by 4. That means we need to print out "FizzBuzz", so let's add that to the body of the if statement.

#### Add on to the code:

```golang
for i := 1; i<=30; i++ {
  if i%4 == 0{
    fmt.Println("FizzBuzz")
  } else {

  }
}
```

#### Say:

> For the rest of the numbers, we want to just print out the number. So in the body of the else statement, let's do that.

#### Add on to the code:

```golang
for i := 1; i<=30; i++ {
  if i%4 == 0{
    fmt.Println("FizzBuzz")
  } else {
    fmt.Println(i)
  }
}
```

#### Say:

> Let's check if this works now!

#### **Execute the Coding Exercise**

The output should be:

1

2

3

FizzBuzz

5

6

7

FizzBuzz

9

10

11

FizzBuzz

13

14

15

FizzBuzz

17

18

19

FizzBuzz

21

22

23

FizzBuzz

25

26

27

FizzBuzz

29

30

#### Say:

> Here we see 1, 2, 3, and then where there should be 4, we have the word "FizzBuzz". Then it continues on 5, 6, 7, and then "FizzBuzz". If you check all the way down, you will see that it printed "FizzBuzz" every 4th time.

> Does anyone have any questions?

### **Conclusion**

#### Say:

> There are different variations of the `for` loop in Golang. For example, if you only give a condition statement, then it will continue to loop through the code until the condition is false. That is also known as the `while` loop. But for our purposes today, we learned about the simple `for` loop.
