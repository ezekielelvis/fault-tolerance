from tabulate import tabulate

def circuit_function(inputs, truth_table):
    # Convert the input string to a binary number
    input_number = int(inputs, 2)
    # Return the output from the truth table
    return truth_table[input_number]

def check_stuck_at_fault(circuit_function, input_size, truth_table):
    # Generate all possible inputs as binary strings
    all_inputs = [bin(i)[2:].zfill(input_size) for i in range(2**input_size)]

    # Check if the output is the same for all inputs
    output = circuit_function(all_inputs[0], truth_table)
    for input in all_inputs[1:]:
        if circuit_function(input, truth_table) != output:
            return False

    # If the output is the same for all inputs, there's a stuck-at-fault
    return True

def main():
    # Define example input size and truth table
    input_size = 3
    # Truth table for 3-input circuit (2^3 = 8 entries)
    truth_table = [0, 1, 0, 1, 1, 0, 1, 0]

    # Check for a stuck-at-fault
    has_fault = check_stuck_at_fault(circuit_function, input_size, truth_table)

    # Print the truth table
    table = [['Inputs'] + [f'Input_{i+1}' for i in range(input_size)] + ['Output']]
    for i, output in enumerate(truth_table):
        table.append([bin(i)[2:].zfill(input_size)] + list(map(str, [int(digit) for digit in bin(i)[2:].zfill(input_size)])) + [output])
    print("\nTruth Table:")
    print(tabulate(table, headers="firstrow", tablefmt="grid"))

    # Print the result
    if has_fault:
        print("\nThe circuit has a stuck-at-fault.")
    else:
        print("\nThe circuit does not have a stuck-at-fault.")

if __name__ == "__main__":
    main()
