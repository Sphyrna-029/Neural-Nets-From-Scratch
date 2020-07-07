package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
    inputs := [4]float64{1, 2, 3, 2.5}
	
    weights := [][]float64{ {0.2, 0.8, -0.5, 1.0}, 
			    {0.5, -0.91, 0.26, -0.5}, 
			    {-0.26, -0.27, 0.17, 0.87} }

    bias := [3]float64{2, 3, 0.5}

	//                                            {              A single neuron            }
	// Inputs times the weights plus the bias (x = i1 * w1 + i2 * w2 + i3 * w3 + i4 * w4 + b).
	// This is a single layer comprising of 3 neurons, 4 inputs, 4 weights specific to each neuron, and 3 bias' for each neuron
        //
	// array := []int{1, 2, 3, 4}                                 Shape = 4 - 1D (AKA Vector)
	//
	// array := [][]int{ {1, 2, 3, 4}, 
	//                   {5, 6, 7, 8} }                           Shape = (2, 4) - 2D
	//
	// array := [][][]int{ { {1, 2, 3, 4}, {5, 6, 7, 8} }, 
	//                     { {0, 1, 2, 3}, {4, 5, 6, 7} }, 
	//                     { {1, 2, 3, 4}, {5, 6, 7, 8} }, }      Shape = (3, 2, 4) - 3D
	//
	// Tensor - An object that can be represented as an array.
	//
	// Dot Product -
	// a := []int{1, 2, 3}
	// b := []int{4, 5, 6}
	//
	// dotproduct := a[0] * b[0] + a[1] * b[1] + a[2] * b[2]
	//
	// fmt.Println(dotproduct) 
	// "32"

    output := [3]float64{inputs[0]*weights[0][0] + inputs[1]*weights[0][1] + inputs[2]*weights[0][2] + inputs[3]*weights[0][3] + bias[0],
                         inputs[0]*weights[1][0] + inputs[1]*weights[1][1] + inputs[2]*weights[1][2] + inputs[3]*weights[1][3] + bias[1],
                         inputs[0]*weights[2][0] + inputs[1]*weights[2][1] + inputs[2]*weights[2][2] + inputs[3]*weights[2][3] + bias[2]}

    fmt.Println(output)
}
