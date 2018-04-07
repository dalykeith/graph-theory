package main

import (
	"bufio" //implements buffered I/O
	"fmt"	//implements formatted I/O
	"os"	//operating system functionality

	
	assets "./src/assets" //imports connecting folder using shunt.go & rega.go
)
func main() {

	option := 0
	exit := true

	
	for exit {

		//user menu
		fmt.Print("\nPlease choose an option:\n1)Infix Expression to Postfix Expresssion\n2)Postix Regular Expresssion to NFA \n3)Exit\n")
		fmt.Scanln(&option) //read in user option

		
		if option == 1 { //if user chooses option 1 (infix to postfix)
			fmt.Print("Please enter the expression to be converted:")
			reader := bufio.NewReader(os.Stdin) //read in string & use created function to trim string

			expression, _ := reader.ReadString('\n')
			expression = assets.TrimEndString(expression)

			// inputs to outputs
			// a.b.c* -> ab.c*.	(a.(b|d))* -> abd|.*	a.(b|d).c* -> abd|.c*.		a.(b.b)+.c -> abb.+.c.

			fmt.Println("Infix:  ", expression) //output the expression
			fmt.Println("Postfix: ", assets.Intopost(expression)) //pass the expression into intopost function and output result

		} else if option == 2 { //if user chooses option 2 (postfix to NFA)

			fmt.Print("Please enter the expression to be converted:")
			reader := bufio.NewReader(os.Stdin)	//read in string & use created function to trim string

			expression, _ := reader.ReadString('\n')
			expression = assets.TrimEndString(expression)

			fmt.Println("Postfix:  ", expression) //output expression
			fmt.Println("NFA: ", assets.Poregtonfa(expression)) //pass the expression into Poregtonfa function and output result
			fmt.Print("\n")

			fmt.Print("Please enter a regular string to see if it matches the NFA:") // prompt user for basic string
			regString, _ := reader.ReadString('\n') // read in string and use created function to trim string
			regString = assets.TrimEndString(regString)
			regString = assets.Intopost(regString)

			//"ab.c*|", "ccc"

			// pass expression and basic string into Pomatch
			// if returns true, strings match
			if assets.Pomatch(expression, regString) == true {
				fmt.Println("Regular string, ", regString, " matches the expression: ", expression) // output result

				// if returns false, output strings do not match
			} else if assets.Pomatch(expression, regString) == false {
				fmt.Print("String does not match")

				// if any errors 
			} else {
				fmt.Print("An error occured please retry")
			}

			// exit program & end loop
		} else if option == 3 {
			fmt.Print("\nProgram Exited!\n")
			exit = false
		} else {
			fmt.Print("\nPlease enter a valid option!\n") //if input error occurs, catch and output
		}

	} //end While loop
} // end main
