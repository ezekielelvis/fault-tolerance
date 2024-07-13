package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Define the Gate struct with fields for name, function, inputs, output, and stuck-at fault value
type Gate struct {
	name     string
	function func(int, int) int
	input1   *Gate
	input2   *Gate
	output   int
	stuckAt  *int
}

// Define logic gate functions: AND, OR, and XOR
func AND(a, b int) int {
	return a & b
}

func OR(a, b int) int {
	return a | b
}

func XOR(a, b int) int {
	return a ^ b
}

// Function to simulate the circuit by calculating the output of each gate
func simulateCircuit(inputA, inputB int, gates []*Gate) {
	for _, gate := range gates {
		// If the gate has both inputs, calculate its output using the function
		if gate.input1 != nil && gate.input2 != nil {
			gate.output = gate.function(gate.input1.output, gate.input2.output)
		} else {
			// Set the output of input gates A and B directly
			if gate.name == "A" {
				gate.output = inputA
			} else if gate.name == "B" {
				gate.output = inputB
			}
		}

		// Override the output if there's a stuck-at fault
		if gate.stuckAt != nil {
			gate.output = *gate.stuckAt
		}
	}
}

// Function to diagnose the circuit and print any stuck-at faults
func diagnoseCircuit(gates []*Gate) {
	fmt.Println("\nFault Diagnosis:")
	for _, gate := range gates {
		if gate.stuckAt != nil {
			fmt.Printf("Gate %s is stuck at %d\n", gate.name, *gate.stuckAt)
		}
	}
}

// Main function to set up and run the circuit simulation
func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	// Define the gates
	inputA := &Gate{name: "A"}
	inputB := &Gate{name: "B"}
	andGate := &Gate{name: "AND", function: AND, input1: inputA, input2: inputB}
	orGate := &Gate{name: "OR", function: OR, input1: inputA, input2: inputB}
	xorGate := &Gate{name: "XOR", function: XOR, input1: andGate, input2: orGate}

	// Create a slice containing all gates
	gates := []*Gate{inputA, inputB, andGate, orGate, xorGate}

	// Introduce a stuck-at fault at a random gate
	stuckValue := rand.Intn(2)
	faultyGate := gates[rand.Intn(len(gates))]
	faultyGate.stuckAt = &stuckValue

	// Print the circuit simulation results
	fmt.Println("Circuit Simulation with Stuck-at Fault:")
	fmt.Println("A | B | AND | OR | XOR")
	fmt.Println("-------------------------")

	// Loop through all possible input combinations (00, 01, 10, 11)
	for i := 0; i < 4; i++ {
		inputAValue := (i >> 1) & 1 // Extract bit 1 (A)
		inputBValue := i & 1        // Extract bit 0 (B)

		simulateCircuit(inputAValue, inputBValue, gates)

		// Print the outputs for the current input combination
		fmt.Printf("%d | %d | %d   | %d  | %d\n",
			inputA.output, inputB.output, andGate.output, orGate.output, xorGate.output)
	}

	// Diagnose the circuit to find any stuck-at faults
	diagnoseCircuit(gates)
}
