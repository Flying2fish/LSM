package main

import (
	"Ylsm"
	"Ylsm/config"
	"log"
	"time"
)

type TestValue struct {
	A int64
	B int64
	C int64
	D string
}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	Ylsm.Start(config.Config{
		DataDir:       `E:\Codes\DB`,
		Level0Size:    100,
		PartSize:      4,
		Threshold:     5000000,
		CheckInterval: 1000,
	})
	d1 := insert()
	d2 := queryAll()
	d3 := deleteAll()
	d4 := deleteAbsent()
	d5 := queryAbsent()
	log.Println("func insert, time consumption", d1)
	log.Println("func queryAll, time consumption", d2)
	log.Println("func deleteAll, time consumption", d3)
	log.Println("func deleteAbsent, time consumption", d4)
	log.Println("func queryAbsent, time consumption", d5)
}

func insert() (duration time.Duration) {
	testV := TestValue{1, 2, 3, "abcdefghijklmnopqrstuvwxyz"}
	count := 0
	start := time.Now()
	defer func() { duration = time.Since(start) }()
	key := []byte{'a', 'a', 'a', 'a', 'a'}
	for a := 0; a < 26; a++ {
		for b := 0; b < 26; b++ {
			for c := 0; c < 26; c++ {
				for d := 0; d < 26; d++ {
					for e := 0; e < 26; e++ {
						key[0] = 'a' + byte(a)
						key[1] = 'a' + byte(b)
						key[2] = 'a' + byte(c)
						key[3] = 'a' + byte(d)
						key[4] = 'a' + byte(e)
						testV.D = string(key) + "abcdefghijklmnopqrstuvwxyz"
						Ylsm.Set[TestValue](string(key), testV)
						count++
					}
				}
			}
		}
	}
	log.Println("insert count:", count)
	return
}

func queryAll() (duration time.Duration) {
	start := time.Now()
	count := 0
	defer func() { duration = time.Since(start) }()
	key := []byte{'a', 'a', 'a', 'a', 'a'}
	for a := 0; a < 26; a++ {
		for b := 0; b < 26; b++ {
			for c := 0; c < 26; c++ {
				for d := 0; d < 26; d++ {
					for e := 0; e < 26; e++ {
						key[0] = 'a' + byte(a)
						key[1] = 'a' + byte(b)
						key[2] = 'a' + byte(c)
						key[3] = 'a' + byte(d)
						key[4] = 'a' + byte(e)
						want := string(key) + "abcdefghijklmnopqrstuvwxyz"
						if ans, ok := Ylsm.Get[TestValue](string(key)); !ok || ans.D != want {
							count++
						}
					}
				}
			}
		}
	}
	log.Println("error count for queryAll:", count)
	return
}

func queryAbsent() (duration time.Duration) {
	start := time.Now()
	count := 0
	defer func() { duration = time.Since(start) }()
	key := []byte{'a', 'a', 'a', 'a', 'a'}
	for a := 0; a < 26; a++ {
		for b := 0; b < 26; b++ {
			for c := 0; c < 26; c++ {
				for d := 0; d < 26; d++ {
					for e := 0; e < 26; e++ {
						key[0] = 'a' + byte(a)
						key[1] = 'a' + byte(b)
						key[2] = 'a' + byte(c)
						key[3] = 'a' + byte(d)
						key[4] = 'a' + byte(e)
						if _, ok := Ylsm.Get[TestValue](string(key)); ok {
							count++
						}
					}
				}
			}
		}
	}
	log.Println("error count for queryAbsent:", count)
	return
}

func deleteAll() (duration time.Duration) {
	start := time.Now()
	defer func() { duration = time.Since(start) }()
	key := []byte{'a', 'a', 'a', 'a', 'a'}
	for a := 0; a < 26; a++ {
		for b := 0; b < 26; b++ {
			for c := 0; c < 26; c++ {
				for d := 0; d < 26; d++ {
					for e := 0; e < 26; e++ {
						key[0] = 'a' + byte(a)
						key[1] = 'a' + byte(b)
						key[2] = 'a' + byte(c)
						key[3] = 'a' + byte(d)
						key[4] = 'a' + byte(e)
						Ylsm.Delete(string(key))
					}
				}
			}
		}
	}
	return
}

func deleteAbsent() (duration time.Duration) {
	start := time.Now()
	defer func() { duration = time.Since(start) }()
	Ylsm.Delete("abcdefg")
	return
}
