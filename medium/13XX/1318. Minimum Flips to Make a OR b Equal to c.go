package t13XX

import "fmt"

func minFlips(a int, b int, c int) int {
	var res int

	got := a | b
	want := c

	for i := 0; i < 32; i++ {
		mask := 1 << i

		//fmt.Println(mask)

		aBit := a & mask
		bBit := b & mask

		gotBit := got & mask
		wantBit := want & mask

		//fmt.Printf("aBit:%v: bBit:%v wantBit:%v\n", aBit, bBit, wantBit)

		if gotBit != wantBit {
			fmt.Printf("aBit:%v: bBit:%v wantBit:%v\n", aBit, bBit, wantBit)
			if (aBit == bBit) && (aBit > 0 && wantBit == 0) {
				res += 2
				continue
			}
			res++
		}
	}

	return res
}
