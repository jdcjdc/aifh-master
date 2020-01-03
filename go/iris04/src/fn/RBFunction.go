package fn

import (
	//"math/rand"
	"fmt"
	//"time"
	// "math"
	"math"
)

//public double evaluate(final double[] x) {
//double value = 0;
//final double width = getWidth();
//
//for (int i = 0; i < getDimensions(); i++) {
//final double center = this.getCenter(i);
//value += Math.pow(x[i] - center, 2) / (2.0 * width * width);
//}
//return Math.exp(-value);
//}
var Dimension int

func Evaluate(x []float64) float64 {
	// x's are indeed the weighted inputs.
	/*
	x[0] = 2.04561472162559
	x[1] = 3.320753780398749
	x[2] = -0.5332015505639599
	x[3] = -0.05290297481441039
	*/
	// var width float64 = 7.1213  // een willikeurig getal tussen -10 en 10
	println(Dimension)
	var totalResult float64
	var width = 4.350749831546924

	for i, xi := range x {
		fmt.Printf("nbr: %f\n", xi)
		var center = center(i)
		fmt.Printf("rnd: %f\n", center)
		//var floatResult = math.Pow(x, float64(2)) //type cast int to float64
		//floatResult += math.Pow(float64(nbr), float64(2)) //type cast int to float64
		// var tmpResult = math.Pow(float64(xi) - center, float64(2)) / (2.0 * width * width) //type cast int to float64
		var tmp01 = math.Pow(float64(xi) - center, float64(2))
		var tmp02 = (2.0 * width * width)
		var result = tmp01 / tmp02  // todo test whether result is always between 0 and 1
		fmt.Printf("tmpResult: %f\n", result)
		totalResult += result
	}
	fmt.Printf("totalResult: %f\n",  totalResult)
	// totalResult = 6.699799596854386  // todo remove
	var endResult = math.Pow(math.E, -totalResult)
	// endResult : 0.011251108668042896
	fmt.Printf("endResult: %f\n", endResult)
	return 0.0
}

func center(nbr int) float64 {
	//rand.Seed(time.Now().UTC().UnixNano())
	//var rndNbr = rand.Float64()
	//rndNbr *= 20
	//rndNbr -= 10
	if nbr == 0 {
		return 0.43307431896961646
	} else if nbr == 1 {
		return -2.424059093762918
	} else if nbr == 2 {
		return -9.967270700171543
	} else if nbr == 3 {
		return 6.675779370897246
	}

	return 0.0
}

//public void performRandomize(final double[] memory) {
//for (int i = 0; i < memory.length; i++) {
//memory[i] = this.rnd.nextDouble(this.lowRange, this.highRange);  // with lowRange: -10.00 highRange: +10
//}
//}
