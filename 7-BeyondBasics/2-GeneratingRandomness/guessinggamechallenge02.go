package main

import (
    "fmt" // import the fmt package which allows you to use text
          //formatting, reading input & printing output functions

    "math/rand" //import the rand package which allows you to
                //generate random numbers

    "time" // import the package which will provide time functionality to measure time
)

func main() {

    // Print some basic information on the screen for the user
    fmt.Println("Game: Guess a number between 0 and 10")
    fmt.Println("How many tries would you like? Please enter a number:")
    var tries int
    fmt.Scan(&tries)

    maxNum := 10
    if tries > 5 {
        fmt.Println("Since you have", tries, "tries, let's make this a little more challenging. You will now be guissing a number between 0 and 100!")
        maxNum = 100
    } else {
        fmt.Println("You have", tries, "tries.")
    }

    // generating random numbers
    source := rand.NewSource(time.Now().UnixNano())
    // The default number generator is predictable, so it will produce the same sequence of numbers each time. To 
    //  produce varying range of numbers, give it a seed that changes (in this case: time would ensure it changes ). 
    //  Note that this is not safe to use for random numbers you want to be secret; use crpyto/rand for those.

    // now we can generate some chaos with our "source" seed
    randomizer := rand.New(source)
    secretNumber := randomizer.Intn(maxNum)  // generate numbers between 0 and 10 only

    // we declare a variable of type int called "guess"
    var guess int

    // create a for loop that begins at 1 and loops to the number that the user entered.
    for try := 1; try <= tries; try++ {
        // declaring the conditions for the for loop ; the shorthand form of declaring a variable was used here. 
        // Declare and Initialize ‘ := ‘ you declare and assign a value upon declaration. Go will automatically 
        // infer the type of the variable since you already assigned a value to it.

        // print out the number of times the player has made a guess
        fmt.Println("Round:", try)

        // the program will prompt the player to make a guess and enter a number
        fmt.Println("Please enter your number")
        fmt.Scan(&guess)   // this function makes it possible for the program to receive the input

        if guess < secretNumber {
            // if the guessed number is less than or greater than the correct number; give the player a hint

            fmt.Printf("Sorry, wrong guess ; number is too small\n ")
        } else if guess > secretNumber {
            fmt.Printf("Sorry, wrong guess ; number is too large\n ")
        } else {
            fmt.Printf("You win!\n")
            break
            // Print out "you win" message when the player guesses the correct number
        }

        if try == tries {
            // if the number of tries is equal to the number the user entered, print game over and also the correct number

            fmt.Printf("Game over!!\n ")
            fmt.Printf("The correct number is %d\n", secretNumber)
        }
    }

    // every time the game ends, print this message
    fmt.Println("Thanks for playing, see you next time!")
}

