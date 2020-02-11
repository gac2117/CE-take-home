# Lesson Plan for Golang - Loops

This lesson assumes students already have some knowledge of Go basic data types, variables, functions, logical operators, conditionals.

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

#### Type the following code in the text editor:

```golang
for initialization; condition; post{

}
```

> We start by using the keyword `for` and then the initialization statement. It is a simple statement like variable declarations. It is executed before the first loop or iteration.

> Next is the condition statement. It is a boolean expression that is executed before every iteration. It is checking whether we should run the loop again or not. If the boolean expression is true, then the loop executes. If it is false, then it stops.

> Finally there is the post statement. It is executed after every iteration. Let's build a simple `for` loop now.

#### Type:

```golang
for i := 0; condition; post{

}
```

#### Say:

> For the initialization statement, we are going to declare a variable 'i' and set it to 0. That will serve as our counter.

#### Add on to the code:

```golang
for i := 0; i < 5; post{

}
```

#### Say:

> For the condition statement, we are going to run the loop as long as 'i' is less than 5. So basically it will run 5 times.

#### Add on to the code:

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

#### Execute the code so it prints out the output, "Coding is fun" 5 times.

#### Say:

> Okay well let's try to see exactly how this is working. Instead of printing "Coding is fun" 5 times, let's print what the variable 'i' is in every loop.

#### Change the code into:

```golang
for i := 0; i < 5; post{
  fmt.Println(i)
}
```

#### Execute the code so it prints out the output, 0 1 2 3 4.

#### Say:

> Notice what the loop is doing. First, before the loop executes, 'i' is 0. After the loop executes, 'i' is incremented by 1 so it is 1. That is still less than 5, so the loop executes and prints out 1. Then it is incremented by 1 again, making it 2. That is still less than 5, so the loop executes.

> This keeps repeating until i is no longer less than 5. So that's why it ends with printing out 4.

#### Say:

> Now you are going to write a couple of `for` loops on your own. One of them is a variation of a very popular coding challenge called "FizzBuzz".

> Open up the assignment and you will see instructions on what kind of `for` loop you will need to write. You have 20 minutes to complete the assignment. If you get stuck, please raise your hand and I will go help you.

As the students are coding, walk around the classroom to make sure no one is stuck.

## Instructor Review (10 mins)

#### Say:

> How did you do? Would anyone like to share their solution code?

See if there are any volunteers. Have the student read their code aloud as you code it into your text editor for everyone to see. Execute the code to see if it works. Thank the student who volunteered.
