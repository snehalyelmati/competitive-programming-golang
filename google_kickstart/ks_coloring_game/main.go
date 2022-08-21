package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var N int
		fmt.Scanf("%d", &N)

		fmt.Printf("Case #%d: %v\n", t+1, 1+(N-1)/5)
	}
}

// ans = 1 + (2-1)/5 + (3-1)/5 + ...

// B * * *
// B * H *

// B * * * *
// B * H * *
// B * H * B

// B * * * *
// B * * H *

// B * * * * *
// B * * * H *
// B * B * H *

// B * * * * *
// B * H * B *

// B * * * * *
// B * * H * B

// B * * * * H
// B * B * * H

// B * * * * * *
// B * H * * * *
// B * H * B * *
// B * H * B * H

// B * * * * * *
// B * * H * * *
// B * * H * B *

// B * * * * * *
// B * * * H * *
// B * B * H * *
// B * B * H * H

// B * * * * * *
// B * * * * H *
// B * B * * H *

// B * * * * * *
// B * * * * * H
// B * B * H * H
