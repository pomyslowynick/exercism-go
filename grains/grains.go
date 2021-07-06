package grains

import (
	"errors"
	"math"
)

// Community solution, I benchmarked it and it's so much faster than math.Pow
// I thought compiler would optimize for the powers of two and the performance would be pretty much the same
// On my machine it's a difference between 294ns adnd 12ns, more than twenty times as much time.
// func Square(n int) (uint64, error) {
// 	if 1 > n || n > 64 {
// 		return uint64(0), errors.New("Invalid")
// 	}

// 	return uint64(1 << uint(n-1)), nil
// }

func Square(chessboardSquare int) (result uint64, err error) {
	if chessboardSquare == 0 {
		return 0, errors.New("number cannot be zero")
	}

	if chessboardSquare < 0 {
		return 0, errors.New("number cannot be negative")
	}

	if chessboardSquare > 64 {
		return 0, errors.New("number cannot be higher than 64")
	}
	result = uint64(math.Pow(2, float64(chessboardSquare-1)))
	// The code below results in a way better performance
	// return uint64(1 << uint(n-1)), nil

	// I think the reason why mat.Pow is less efficient is because it operates on floats?
	return result, nil
}

func Total() (sum uint64) {
	for i := 1; i <= 64; i++ {
		val, err := Square(i)
		if err != nil {
			panic("squaring operation threw an error")
		}
		sum += val
	}

	return sum
}
