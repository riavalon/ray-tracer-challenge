package tuples

import (
	"math"
	"testing"
)

func Test_CreateTuple_InitializesWithGivenValues(t *testing.T) {
	x, y, z, w := 4.3, -4.2, 3.1, 1.0
	tuple := CreateTuple(x, y, z, w)

	if tuple.W != 1.0 {
		t.Errorf("CreateTuple should initialize with proper w value.  expected %v; got %v", w, tuple.W)
	}

	if tuple.X != x {
		t.Errorf("CreateTuple should initialize with proper X value.  expected %v; got %v", x, tuple.X)
	}

	if tuple.Y != y {
		t.Errorf("CreateTuple should initialize with proper Y value.  expected %v; got %v", y, tuple.Y)
	}

	if tuple.Z != z {
		t.Errorf("CreateTuple should initialize with proper Z value.  expected %v; got %v", z, tuple.Z)
	}
}

func Test_Tuple_IsPointFunction(t *testing.T) {
	tuplePoint := CreateTuple(4.3, -3.2, 5.2, 1.0)
	tupleVector := CreateTuple(4.3, -4.3, -2, 0.0)

	if tuplePoint.IsPoint() == false && tuplePoint.W == 1.0 {
		t.Errorf("tuple.isPoint should return true if W value is 1.0 - W = %v", tuplePoint.W)
	}

	if tupleVector.IsPoint() == true && tuplePoint.W == 0.0 {
		t.Errorf("tuple.IsPoint should return false if W value is 0 - W = %v", tupleVector.W)
	}
}

func Test_CreatePoint_CreatesTupleWithPointValue(t *testing.T) {
	point := CreatePoint(4, -4, 3)

	if point.W != 1.0 {
		t.Errorf("CreatePoint should return tuple with W value of 1.0 but got %v", point.W)
	}
}

func Test_CreateVector_CreatesTupleWithVectorValue(t *testing.T) {
	vec := CreateVector(4, -4, 3)

	if vec.W != 0.0 {
		t.Errorf("CreateVector should return tuple with W value of 0.0 but got %v", vec.W)
	}
}

func Test_TupleIsEquivalentTo_ReturnsTrueForTuplesWithMatchingValues(t *testing.T) {
	pointOne := CreatePoint(1, 2, 3)
	pointTwo := CreatePoint(1, 2, 3)
	result := pointOne.IsEquivalentTo(pointTwo)

	if !result {
		t.Errorf("Point 1 should be equivalent to Point 2;\nPoint 1: %v\nPoint 2: %v", pointOne, pointTwo);
	}
}

func Test_TupleIsEquivalentTo_ReturnsFalseForTuplesWithMisMatchingValues(t *testing.T) {
	pointOne := CreatePoint(1, 2, 3)
	pointTwo := CreatePoint(3, 2, 1)
	result := pointOne.IsEquivalentTo(pointTwo)

	if result {
		t.Errorf("Point 1 should NOT be equivalent to point 2;\nPoint 1: %v\nPoint 2: %v", pointOne, pointTwo)
	}
}

func Test_TupleIsEquivalentTo_ReturnsFalseIfTuplesAreOfDifferentType(t *testing.T) {
	point := CreatePoint(1, 2, 3)
	vector := CreateVector(1, 2, 3)
	result := point.IsEquivalentTo(vector)

	if result {
		t.Errorf("Points and Vectors should not be considered equal")
	}
}

func Test_TupleIsEquivalentTo_ReturnsFalseForValuesLargerThanEpsilon(t *testing.T) {
	pointOne := CreatePoint(0.3, 0.5, 1.2)
	pointTwo := CreatePoint(0.2, 0.5, 1.2)
	result := pointOne.IsEquivalentTo(pointTwo) // 0.3 - 0.2 in floating point math -> 0.99999999998

	if result {
		t.Errorf("Difference values larger than EPSILON should not be considered equivalent")
	}
}

// ARTHIMETIC TESTS --------------------------------------------------

func Test_TupleAdd_AddsTwoTuplesTogether(t *testing.T) {
	a := CreatePoint(3, -2, 5)
	b := CreateVector(-2, 3, 1)
	expected := CreatePoint(1, 1, 6)
	result, _ := a.Add(b)

	if !result.IsEquivalentTo(expected) {
		t.Errorf("Place values of tuples should be properly added together.\nWanted x1, y1, z6;\nGot x%v, y%v, z%v", result.X, result.Y, result.Z)
	}
}

func Test_TupleAdd_AddsTwoVectors(t *testing.T) {
	a := CreateVector(1, 2, 3)
	b := CreateVector(3, 2, 1)
	expected := CreateVector(4, 4, 4)
	result, _ := a.Add(b)

	if !result.IsEquivalentTo(expected) {
		t.Errorf("Place values of tuples should be properly added together.\nWanted x4, y4, z4;\nGot x%v, y%v, z%v", result.X, result.Y, result.Z)
	}

	if result.IsPoint() {
		t.Errorf("A vector added to a vector should also be a vector, but got a point")
	}
}

func Test_TupleAdd_ReturnsAnErrorIfBothTuplesArePoints(t *testing.T) {
	a := CreatePoint(1, 2, 3)
	b := CreatePoint(3, 2, 1)
	result, err := a.Add(b)

	if err == nil {
		t.Errorf("Expected Tuple.Add to return an error if both tuples are points")
	}

	if result != nil {
		t.Errorf("Should not return a result if both tuples are points, only an error")
	}
}

func Test_TupleSubtract_SubtractsTwoVectors(t *testing.T) {
	a := CreateVector(3, 2, 1)
	b := CreateVector(6, 5, 7)
	result, _ := a.Subtract(b)
	expected := CreateVector(-3, -3, -6)

	if !result.IsEquivalentTo(expected) {
		t.Errorf("Expected new vector with place values equal to the difference of Vectors A and B.\nWant -3 -3 -6\nGot %v %v %v", result.X, result.Y, result.Z)
	}
}

func Test_TupleSubtract_SubtractsTwoPoints(t *testing.T) {
	a := CreatePoint(3, 2, 1)
	b := CreatePoint(1, 2, 3)
	expected := CreateVector(2, 0, -2)
	result, _ := a.Subtract(b)

	if !result.IsEquivalentTo(expected) {
		t.Errorf("Expected two points to return a new vector with place values equal to the difference of A and B.\nWant 2 0 -2\nGot %v %v %v", result.X, result.Y, result.Z)
	}

	if result.IsPoint() {
		t.Errorf("Expected a vector from point minus point but instead got another point")
	}
}

func Test_TupleSubtract_SubtractsPointAndVector(t *testing.T) {
	a := CreatePoint(1, 2, 3)
	b := CreateVector(3, 2, 1)
	expected := CreatePoint(-2, 0, 2)
	result, _ := a.Subtract(b)

	if !result.IsEquivalentTo(expected) {
		t.Errorf("Expected new point with values equal to the difference of Tuples A and B.\nWant 2 0 -2;\nGot %v %v %v", result.X, result.Y, result.Z)
	}

	if !result.IsPoint() {
		t.Errorf("Expected that a point minus a vector would result in a point but got a vector instead")
	}
}

func Test_TupleSubtract_ReturnsErrorIfResultIsNeitherVectorOrPoint(t *testing.T) {
	a := CreateVector(1, 2, 3)
	b := CreatePoint(3, 2, 1)
	result, err := a.Subtract(b)

	if err == nil {
		t.Errorf("Expected an error to be returned, but no error was found")
	}

	if result != nil {
		t.Errorf("Expected to return an error and a nil result, but got a result anyway")
	}
}

func Test_TupleNegate_NegatesAGivenTuplesValues(t *testing.T) {
	vector := CreateVector(1, -2, 3)
	expected := CreateVector(-1, 2, -3)
	result := vector.Negate()

	if !result.IsEquivalentTo(expected) {
		t.Errorf("Expected the result to be negated.\nGot %v %v %v;\nWant -1 2 -3", result.X, result.Y, result.Z)
	}
}

func Test_TupleMultiply_MultiplyATupleByAScalar(t *testing.T) {
	scalar := 3.5
	tuple := CreateTuple(1, -2, 3, -4)
	expected := CreateTuple(3.5, -7, 10.5, -14)
	result := tuple.Multiply(scalar)

	if !result.IsEquivalentTo(expected) {
		t.Errorf(
			"Expected result to be uniformly scaled by multiplication.\nWant 3.5 -7 10.5 -14\nGot %v %v %v %v",
			result.X,
			result.Y,
			result.Z,
			result.W,
		)
	}
}

func Test_TupleMultiply_MultiplyByAFraction(t *testing.T) {
	fraction := 0.5
	tuple := CreateTuple(1, -2, 3, -4)
	expected := CreateTuple(0.5, -1, 1.5, -2)
	result := tuple.Multiply(fraction)

	if !result.IsEquivalentTo(expected) {
		t.Errorf(
			"Expected result to have each place value halved by the fraction 0.5 -\nWant 0.5 -1 1.5 -2\nGot %v %v %v %v",
			result.X,
			result.Y,
			result.Z,
			result.W,
		)
	}
}

func Test_TestDivide_DividesByAScalar(t *testing.T) {
	var scalar float64 = 2
	tuple := CreateTuple(1, -2, 3, -4)
	expected := CreateTuple(0.5, -1, 1.5, -2)
	result := tuple.Divide(scalar)

	if !result.IsEquivalentTo(expected) {
		t.Errorf(
			"Expected result to have each place value divided by %v\nWant 0.5 -1 1.5 -2\nGot %v %v %v %v",
			scalar,
			result.X,
			result.Y,
			result.Z,
			result.W,
		)
	}
}

// MAGNITUDE TESTS --------------------------------------------------

func Test_TupleMagnitude_ComputesMagnitudeForUnitVectorByX(t *testing.T) {
	vector := CreateVector(1, 0, 0)
	var expected float64 = 1
	result := vector.GetMagnitude()

	if result != expected {
		t.Errorf("Failed to compute proper magnitude. Wanted %v; Got %v", expected, result)
	}
}

func Test_TupleMagnitude_ComputesMagnitudeForUnitVectorByY(t *testing.T) {
	vector := CreateVector(0, 1, 0)
	var expected float64 = 1
	result := vector.GetMagnitude()

	if result != expected {
		t.Errorf("Failed to compute proper magnitude. Wanted %v; Got %v", expected, result)
	}
}

func Test_TupleMagnitude_ComputesMagnitudeForUnitVectorByZ(t *testing.T) {
	vector := CreateVector(0, 0, 1)
	var expected float64 = 1
	result := vector.GetMagnitude()

	if result != expected {
		t.Errorf("Failed to compute proper magnitude. Wanted %v; Got %v", expected, result)
	}
}

func Test_TupleMagnitude_ComputesNonUnitVectorMagnitudeWithPostiveValues(t *testing.T) {
	vector := CreateVector(1, 2, 3)
	expected := math.Sqrt(14)
	result := vector.GetMagnitude()

	if result != expected {
		t.Errorf("Failed to computer proper magnitude. Wanted %v; Got %v", expected, result)
	}
}

func Test_TupleMagnitude_ComputesNonUnitVectorMagnitudeWithNegativeValues(t *testing.T) {
	vector := CreateVector(-1, -2, -3)
	expected := math.Sqrt(14)
	result := vector.GetMagnitude()

	if result != expected {
		t.Errorf("Failed to compute proper magnitude. Wanted %v; Got %v", expected, result)
	}
}

func Test_TupleNormalize_NormalizingVectors(t *testing.T) {
	// Normalizes Vector 4, 0, 0
	vector := CreateVector(4, 0, 0)
	exp1 := CreateVector(1, 0, 0)
	res1 := vector.Normalize()

	if !res1.IsEquivalentTo(exp1) {
		t.Errorf(
			"Failed to normalize vector. Want %v %v %v; Got %v %v %v",
			exp1.X, exp1.Y, exp1.Z, res1.X, res1.Y, res1.Z,
		)
	}

	// Normalizes Vector 1, 2, 3
	vector2 := CreateVector(1, 2, 3)
	exp2 := CreateVector(0.26726, 0.53452, 0.80178) // Approx. 1/√14 2/√14 3/√14
	res2 := vector2.Normalize()

	if !res2.IsEquivalentTo(exp2) {
		t.Errorf(
			"Failed to normalize vector. Want %v %v %v; Got %v %v %v",
			exp2.X, exp2.Y, exp2.Z, res2.X, res2.Y, res2.Z,
		)
	}
}

func Test_TupleNormalize_MagnitudeOfNormalizedVectorIsOne(t *testing.T) {
	vector := CreateVector(1, 2, 3)
	var expected float64 = 1
	result := vector.Normalize().GetMagnitude()

	if result != expected {
		t.Errorf("A normalized vector should have a magnitude of one but got %v", result)
	}
}
