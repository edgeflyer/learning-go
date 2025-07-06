package main

import (
	"bytes"
	"fmt"
)

// IntSet是一个包含非负整数的集合
// 零值代表空的集合
type IntSet struct {
	words []uint64
}

// Has方法的返回值表示是否存在非负数x
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add 添加非负数x到集合中
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith将会对s和t做并集并将结果存在s中
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String方法以字符串{1 2 3}的形式返回集合
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
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// practice
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			count += int(word & (1 << j))
		}
	}
	return count
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, x%64
	s.words[word] &= ^(1 << uint64(bit))
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	copySet := &IntSet{}
	copySet.words = append([]uint64(nil), s.words...)
	return copySet
}

func (s *IntSet) AddAll(nums ...int) error {
	for _, x := range nums {
		word, bit := x/64, x%64

		for word > len(s.words) {
			s.words = append(s.words, 0)
		}
		if s.words[word]&(1<<uint64(bit)) != 0 {
			return fmt.Errorf("%d already added", x)
		}
		s.words[word] |= 1 << uint64(bit)
	}
	return nil
}
