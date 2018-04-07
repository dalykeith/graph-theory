package assets

import (
	"fmt" // formats
)

type state struct {
	symbol rune // letter to determine epsilon
	edge1 *state // arrow from a state to point at other state
	edge2 *state
}

// helper struct
type nfa struct {
	// keeps track of fragment of NFA
	initial *state
	accept  *state
}

// Post Reg Expression to NFA, returns a pointer to an NFA struct
func Poregtonfa(pofix string) *nfa {
	nfastack := []*nfa{} // keeps track of fragments on a stack (pointer to nfastack)

	// loop through pofix expression 1 char(rune) at a time
	for _, r := range pofix {
		switch r {

		// concatenate character is represented by an Full stop (.)
		case '.':
			// pop frag2 of nfastack
			frag2 := nfastack[len(nfastack)-1] // index of last item on nfastack
			// gives program everything (up to but not including the last item) in nfastack
			nfastack = nfastack[:len(nfastack)-1]

			// pop frag1 of nfastack
			frag1 := nfastack[len(nfastack)-1] //Index of last item on nfastack
			// gives program everything (up to but not including the last item) in nfastack
			nfastack = nfastack[:len(nfastack)-1]

			//frag1 and frag2 are pointers to 2 nfa fragments

			// link (concatenate) frag1 and frag2 - accept state of frag1's edge1
			// points to frag2's initial state
			frag1.accept.edge1 = frag2.initial

			// append new pointer to nfa struct that represents new bigger nfa fragment which is constructed from frag1 and frag2
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})

		// union character is represented by an Vertical bar (|)
		case '|':
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			// new state which just points at the initial state of the 2 fragments that were popped off the stack
			// new initial state where edge1 points at frag1.initial and edge2 points at frag2.initial order does not matter because its an 'or'
			initial := state{edge1: frag1.initial, edge2: frag2.initial}

			// create a normal new state
			accept := state{}

			// the fragments edges has to point at the accept state
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			// still a pointer to an NFA but we want the initial state of the fragment we are pushing to be
			// the new initial state we created above and we want the new accept state to be the accept state
			// of the fragment that we are pushing to the nfa stack nfa stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		// kleene star is represented by an asterisk aka star (*)
		case '*':
			// pop 1 fragment of nfastack (Kleene star only works on one fragment of the nfa)
			frag := nfastack[len(nfastack)-1]
			// gives program everything (up to but not including the last item) in nfastack
			nfastack = nfastack[:len(nfastack)-1]
			// create a normal new state
			accept := state{}
			// edge1 needs to be in iintial  state of fragment we popped off and edge2 needs to point at the new accept state
			initial := state{edge1: frag.initial, edge2: &accept}
			// join accept state (edge1) for the fragment to initial state of fragment that was popped off
			frag.accept.edge1 = frag.initial
			// join edge2 of old fragment to the accept state
			frag.accept.edge2 = &accept

			// push new fragment to nfastack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
			// plus symbol
		case '+':
			// pop 1 fragment of nfastack (Kleene star only works on one fragment of the nfa)
			frag := nfastack[len(nfastack)-1]
			// create a normal new state
			accept := state{}
			// edge1 needs to be in iintial  state of fragment we popped off and edge2 needs to point at the new accept state
			initial := state{edge1: frag.initial, edge2: &accept}
			// the fragment edge has to point at the initial state
			frag.accept.edge1 = &initial
			// push new fragment to nfastack
			nfastack = append(nfastack, &nfa{initial: frag.initial, accept: &accept})

			//question mark symbol
		case '?':
			// pop 1 fragment of nfastack (Kleene star only works on one fragment of the nfa)
			frag := nfastack[len(nfastack)-1]
			initial := state{edge1: frag.initial, edge2: frag.accept}

			accept := state{edge1: frag.initial, edge2: frag.accept}
			frag.accept.edge1 = &accept

			// the fragment edge has to point at the initial state
			frag.accept.edge1 = &initial

			// push new fragment to nfastack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: frag.accept})

		// any other character inputted 
		default:
			// create a new accept state
			accept := state{}
			// create a new initial state and set symbol to r and its only edge points to the accept state
			initial := state{symbol: r, edge1: &accept}

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept}) 	// push new fragment to nfastack
		}
	}

	// if the length of the nfastack is anything but 1
	if len(nfastack) != 1 {
		fmt.Println("Unfortunatly there was ", len(nfastack), " NFA's found, they are: ", nfastack)
	}
	return nfastack[0] //Return only item (actual nfa) that is left on the stack
}

// helper function, works recursively
// accepts a list of pointers to states, a single pointer to a state, and the accept state
// returns a list of pointers to state
func addState(l []*state, s *state, a *state) []*state {
	// append the state that has been passed in to the list
	l = append(l, s)

	// if 'a' does not, and the rune(s) equals 0 then it contains 'e' arrows
	if s != a && s.symbol == 0 {

		l = addState(l, s.edge1, a)
		// if edge2 is not null, do the same
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}
	return l 	// return the list of pointers to a state
}

// takes in postfix reg expression and a regular string, returns true if they match, false if they dont
func Pomatch(po string, s string) bool {

	// initialise boolean
	ismatch := false

	// variable which passes in the postfix expression, which is ran through poregtonfa
	ponfa := Poregtonfa(po)

	// creates an array of pointers to state (basically a linked list of states)
	current := []*state{}

	// an array of pointers to state that I can get through from current
	next := []*state{}

	// pass current, initial and accept state to addState
	current = addState(current[:], ponfa.initial, ponfa.accept)

	// loop through one s, one char at a time
	for _, r := range s {

		// everytime a char is read, loop through current array
		for _, c := range current {

			// if current symbol is the same is the symbol I am currently reading from s
			if c.symbol == r {

				// returns the list of states and puts them into next
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		// swap current with next, and then create a new array for next
		current, next = next, []*state{}
	}

	for _, c := range current { // loop through the current array
		if c == ponfa.accept { // if state in current array = accept state of ponfa follow set
			ismatch = true // set ismatch to true
			break
		}
	}
	return ismatch // true or false return
}
