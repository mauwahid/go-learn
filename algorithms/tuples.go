package main

func powerSeries(a int) (int, int) {
	square := a * a
	cube := square * a
	return square, cube
}
