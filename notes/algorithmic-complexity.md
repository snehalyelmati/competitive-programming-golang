# Algorithmic Complexity

## BigO Notation
- Analyzing how fast the program grows asymtotically.
- O(n) isn't always faster than O(n^2), initially algorithms with O(n^2) might be faster than O(n) but O(n^2) eventually eclipse the runtime of O(n).
- Three common notations,
	- **Omega** - for best case time complexity.
	- **Big O** - for worst case time complexity.
	- **Theta** - where the best case and worst case time complexities are the same.
- For binary search, 
	- Omega(1) - where the middle element is the target element.
	- O(logn) - first element is the target element.
- For linear search,
	- Omega(1) - the first element is the target element.
	- O(n) - the last element is the target element.
- Example for theta, we store the length of the array to look up its size instead of traversing through the entire array. In this case the best and worst case time complexities are the same, we represent it as theta(1).

- [Cheatsheet for time complexities](https://www.bigocheatsheet.com/)

