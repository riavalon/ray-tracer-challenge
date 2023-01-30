package tuples

import (
	"fmt"
	"math"
)

// EPSILON constant is used for checking equivalency when doing floating point arithmetic.
// Equivalency is determined by checking if the difference of two tuples is less than this constant.
const EPSILON = 0.00001

// Tuple struct defines the basic shape of a point or vector
type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

// IsPoint is a utility function for tuples to check if it is representing a point or vector
// returns `true` if tuple is a point.
func (t *Tuple) IsPoint() bool {
	if t.W == 1.0 {
		return true
	}
	return false
}

// IsEquivalentTo compares the tuple instance against another tuple to see if the values are each
// equivalent by checking the difference and asserting it's less than the EPSILON value.  If tuples
// are of differing types, they will not be considered equivalent (point vs vector)
func (t *Tuple) IsEquivalentTo(b *Tuple) bool {
	if t.W != b.W {
		return false
	}
	if math.Abs(t.X - b.X) > EPSILON {
		return false
	}
	if math.Abs(t.Y - b.Y) > EPSILON {
		return false
	}
	if math.Abs(t.Z - b.Z) > EPSILON {
		return false
	}
	return true
}

// Add two tuples together combining the sum of each tuple and returns a brand new tuple.
func (t *Tuple) Add(b *Tuple) (*Tuple, error) {
	if t.IsPoint() && b.IsPoint() {
		return nil, fmt.Errorf("can't add two points")
	}
	newTuple := CreateTuple(
		t.X + b.X,
		t.Y + b.Y,
		t.Z + b.Z,
		t.W + b.W,
	)
	return newTuple, nil
}

// Subtract two tuples and return a new tuple with the difference.  This works with
// two points, two vectors, or a vector and a point.
func (t *Tuple) Subtract(b *Tuple) (*Tuple, error) {
	newTuple := CreateTuple(
		t.X - b.X,
		t.Y - b.Y,
		t.Z - b.Z,
		t.W - b.W,
	)
	if newTuple.W < 0 {
		return nil, fmt.Errorf("cannot subtract a point from a vector")
	}
	return newTuple, nil
}

// Negate a tuple's place values.  Operates on all values including W.
// returns a brand new tuple.
func (t *Tuple) Negate() *Tuple {
	return CreateTuple(
		0 - t.X,
		0 - t.Y,
		0 - t.Z,
		0 - t.W,
	)
}

// Multiply a tuples place values uniformly by a given value.
func (t *Tuple) Multiply(scalar float64) *Tuple {
	return CreateTuple(
		t.X * scalar,
		t.Y * scalar,
		t.Z * scalar,
		t.W * scalar,
	)
}

// Divide a tuples place values uniformly by a give value
func (t *Tuple) Divide(scalar float64) *Tuple {
	return CreateTuple(
		t.X / scalar,
		t.Y / scalar,
		t.Z / scalar,
		t.W / scalar,
	)
}

// GetMagnitude returns the total distance of a given vector.
func (t *Tuple) GetMagnitude() float64 {
	return math.Sqrt(
		math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2) + math.Pow(t.W, 2),
	)
}

// Normalize converts vector into a unit vector
func (t *Tuple) Normalize() *Tuple {
	return t.Divide(t.GetMagnitude())
}

// CreateTuple returns a new tuple with x, y, z, and w values
func CreateTuple(x, y, z, w float64) *Tuple {
	return &Tuple{X: x, Y: y, Z: z, W: w}
}

// CreatePoint returns a tuple with a W value of 1.0
func CreatePoint(x, y, z float64) *Tuple {
	return &Tuple{X: x, Y: y, Z: z, W: 1.0}
}

// CreateVector returns a tuple with a W value of 0.0
func CreateVector(x, y, z float64) *Tuple {
	return &Tuple{X: x, Y: y, Z: z, W: 0.0}
}