# Arrays

## Static Arrays
- Contiguous bock of memory broken down into equal-size elements.
- Advantage of arrays is that random access is possible. Constant time access to read or write.
- Constant time to add/remove at the end, linear time to add/remove at the beginning.
- Example, using arrays to compute primes - Sieve of Eratosthenes. 
	```
	// print primes from 2 to n
	func printPrimes(n int) []int {
		record := make([]int, n+1)
		for i:=2; i*i<=n; i++ {
			if record[i]==0 {
				for j:=2*i; j<=n; j=j+i {
					record[i*j] = 1
				}
			}
		}

		res := make([]int, 0)
		for i:=2; i<=n; i++ {
			if record[i] == 0 {
				res = append(res, i)
			}
		}

		return res
	}
	```

## Dynamic Arrays
- When the array is filled, create a new array with the required size and update the pointer to the array with the new pointer pointing to the newly allocated array.
- Dynamic arrays have basic functions like, set(i, val), get(i), pushBack(a), etc.
- If the columns in all the rows is not the same, then those types of arrays are called as Jagged Arrays.
