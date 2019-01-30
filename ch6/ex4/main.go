package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := &IntSet{}
	a.Add(2)
	a.Add(1)
	a.Remove(1)
	fmt.Println(a.Len())
	a.Add(100)
	a.Clear()
	fmt.Println(a.Len())
}

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", i*64+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, x%64
	if word < len(s.words) {
		s.words[word] &^= 1 << uint(bit)
	}
}

func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

func (s *IntSet) Copy() *IntSet {
	return &IntSet{
		words: append([]uint64(nil), s.words...),
	}
}

func (s *IntSet) IsInclude(t *IntSet) bool {
	for i, tword := range t.words {
		for j := 0; j < 64; j++ {
			if tword&(1<<uint(j)) != 0 {
				if s.words[i]&(1<<uint(j)) == 0 {
					return false
				}
			}
		}
	}
	return true
}

func (s *IntSet) AddAll(add ...int) {
	for _, a := range add {
		s.Add(a)
	}
}

// 共通部分
func (s *IntSet) IntersectWith(t *IntSet) {
	for i := range s.words {
		if i >= len(t.words) {
			s.words[i] = 0
			continue
		}
		s.words[i] &= t.words[i]
	}
}

// sにしかない要素
func (s *IntSet) DiffereceWith(t *IntSet) {
	for i := range s.words {
		if i >= len(t.words) {
			break
		}
		s.words[i] &^= t.words[i]
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 対照差
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i >= len(s.words) {
			s.words = append(s.words, t.words[i:]...)
			break
		}
		s.words[i] ^= tword
	}
}

func (s *IntSet) Elems() []int {
	elems := []int{}
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, 64*i+j)
			}
		}
	}
	return elems
}
