package beeku

import (
	"math/rand"
	"time"
)

type reducetype func(interface{}) interface{}
type filtertype func(interface{}) bool

func Slice_randList(min, max int) []int {
	if max < min {
		min, max = max, min
	}
	length := max - min + 1
	t0 := time.Now()
	rand.Seed(int64(t0.Nanosecond()))
	list := rand.Perm(length)
	for index, _ := range list {
		list[index] += min
	}
	return list
}

func Slice_merge(slice1, slice2 []string) (c []interface{}) {
	c = append(slice1, slice2...)
	return
}

func In_slice(val interface{}, slice []string) bool {
	for _, v := range slice {
		if val == v {
			return true
		}
	}
	return false
}

func Slice_reduce(slice []string, a reducetype) (dslice []interface{}) {
	for _, v := range slice {
		dslice = append(dslice, a(v))
	}
	return
}

func Slice_rand(a []string) (b interface{}) {
	randnum := rand.Intn(len(a))
	b = a[randnum]
	return
}

func Slice_sum(intslice []int64) (sum int64) {
	for _, v := range intslice {
		sum += v
	}
	return
}

func Slice_filter(slice []string, a filtertype) (ftslice []interface{}) {
	for _, v := range slice {
		if a(v) {
			ftslice = append(ftslice, v)
		}
	}
	return
}

func Slice_diff(slice1, slice2 []string) (diffslice []interface{}) {
	for _, v := range slice1 {
		if In_slice(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

func Slice_intersect(slice1, slice2 []string) (diffslice []interface{}) {
	for _, v := range slice1 {
		if !In_slice(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

func Slice_chunk(slice []string, size int) (chunkslice [][]interface{}) {
	if size >= len(slice) {
		chunkslice = append(chunkslice, slice)
		return
	}
	end := size
	for i := 0; i <= (len(slice) - size); i += size {
		chunkslice = append(chunkslice, slice[i:end])
		end += size
	}
	return
}

func Slice_range(start, end, step int64) (intslice []int64) {
	for i := start; i <= end; i += step {
		intslice = append(intslice, i)
	}
	return
}

func Slice_pad(slice []string, size int, val interface{}) []interface{} {
	if size <= len(slice) {
		return slice
	}
	for i := 0; i < (size - len(slice)); i++ {
		slice = append(slice, val)
	}
	return slice
}

func Slice_unique(slice []interface{}) (uniqueslice []interface{}) {
	for _, v := range slice {
		if !In_slice(v, uniqueslice) {
			uniqueslice = append(uniqueslice, v)
		}
	}
	return
}

func Slice_shuffle(slice []string) []interface{} {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
	return slice
}
