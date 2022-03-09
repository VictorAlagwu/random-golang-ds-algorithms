package main


//Write a function that takes in a non-empty array of integers and returns an
//array of the same length, where each element in the output array is equal to
//the product of every other number in the input array.
func main() {

}

func arrayOfProductsBruteForce(array []int) []int {
	// Write your code here.
	arr := []int{}
	for i := 0; i < len(array); i++ {
		currentMultiple := 0
		for j := 1; j < len(array) - 1; j++ {

			if i + 1 == j {
				break
			}
			currentMultiple *= i

		}
		arr = append(arr, currentMultiple )
	}
	return nil
}