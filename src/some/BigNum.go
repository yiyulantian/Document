package some

import "fmt"

func BigNum() {
	BigNumSub()
}

func BigNumSub() {
	a := "100087654321"
	b := "123456789"
	//123456789 864197532
	abyte := []byte(a)
	bbyte := []byte(b)
	alen := len(abyte)
	blen := len(bbyte)
	cover := 0
	var i int
	for i = 0; i < blen; i++ {
		at := alen - 1 - i
		bt := blen - 1 - i
		anbyte := int(abyte[at]) - cover
		bnbyte := int(bbyte[bt])
		cover = 0
		if anbyte < bnbyte {
			cover = 1
			anbyte += 10
			abyte[at] = byte(anbyte - bnbyte + 48)
		} else {
			abyte[at] = byte(anbyte - bnbyte + 48)
		}
	}
	for ; cover > 0; i++ {
		at := alen - i - 1

		anbyte := int(abyte[at]) - cover
		cover = 0
		if anbyte < 48 {
			cover = 1
			anbyte += 10
			abyte[at] = byte(anbyte)
		} else {
			abyte[at] = byte(anbyte)
		}
	}
	for i = range abyte {
		if abyte[i] != '0' {
			break
		}
	}
	abyte = abyte[i:]
	fmt.Printf("%s", abyte)
}
