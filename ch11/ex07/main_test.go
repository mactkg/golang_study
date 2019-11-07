package main

import (
	"math/rand"
	"testing"
)

// Add
func benchmarkIntSetAdd(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		set := IntSet{}
		for index := 0; index < size; index++ {
			set.Add(rand.Intn(3000))
		}
	}
}
func BenchmarkIntSetAdd1(b *testing.B)     { benchmarkIntSetAdd(b, 1) }
func BenchmarkIntSetAdd10(b *testing.B)    { benchmarkIntSetAdd(b, 10) }
func BenchmarkIntSetAdd100(b *testing.B)   { benchmarkIntSetAdd(b, 100) }
func BenchmarkIntSetAdd1000(b *testing.B)  { benchmarkIntSetAdd(b, 1000) }
func BenchmarkIntSetAdd10000(b *testing.B) { benchmarkIntSetAdd(b, 10000) }

func benchmarkMapIntSetAdd(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		set := MapIntSet{}
		for index := 0; index < size; index++ {
			set.Add(rand.Intn(3000))
		}
	}
}
func BenchmarkMapIntSetAdd1(b *testing.B)     { benchmarkMapIntSetAdd(b, 1) }
func BenchmarkMapIntSetAdd10(b *testing.B)    { benchmarkMapIntSetAdd(b, 10) }
func BenchmarkMapIntSetAdd100(b *testing.B)   { benchmarkMapIntSetAdd(b, 100) }
func BenchmarkMapIntSetAdd1000(b *testing.B)  { benchmarkMapIntSetAdd(b, 1000) }
func BenchmarkMapIntSetAdd10000(b *testing.B) { benchmarkMapIntSetAdd(b, 10000) }

// Union
func benchmarkIntSetUnionWith(b *testing.B, size int) {
	x := IntSet{}
	y := IntSet{}
	for index := 0; index < size; index++ {
		x.Add(rand.Intn(3000))
		y.Add(rand.Intn(3000))
	}
	for i := 0; i < b.N; i++ {
		x.UnionWith(&y)
	}
}
func BenchmarkIntSetUnionWith1(b *testing.B)     { benchmarkIntSetUnionWith(b, 1) }
func BenchmarkIntSetUnionWith10(b *testing.B)    { benchmarkIntSetUnionWith(b, 10) }
func BenchmarkIntSetUnionWith100(b *testing.B)   { benchmarkIntSetUnionWith(b, 100) }
func BenchmarkIntSetUnionWith1000(b *testing.B)  { benchmarkIntSetUnionWith(b, 1000) }
func BenchmarkIntSetUnionWith10000(b *testing.B) { benchmarkIntSetUnionWith(b, 10000) }

func benchmarkMapIntSetUnionWith(b *testing.B, size int) {
	x := MapIntSet{}
	y := MapIntSet{}
	for index := 0; index < size; index++ {
		x.Add(rand.Intn(3000))
		y.Add(rand.Intn(3000))
	}
	for i := 0; i < b.N; i++ {
		x.UnionWith(&y)
	}
}
func BenchmarkMapIntSetUnionWith1(b *testing.B)     { benchmarkMapIntSetUnionWith(b, 1) }
func BenchmarkMapIntSetUnionWith10(b *testing.B)    { benchmarkMapIntSetUnionWith(b, 10) }
func BenchmarkMapIntSetUnionWith100(b *testing.B)   { benchmarkMapIntSetUnionWith(b, 100) }
func BenchmarkMapIntSetUnionWith1000(b *testing.B)  { benchmarkMapIntSetUnionWith(b, 1000) }
func BenchmarkMapIntSetUnionWith10000(b *testing.B) { benchmarkMapIntSetUnionWith(b, 10000) }

// Has
func benchmarkIntSetHas(b *testing.B, size int) {
	x := IntSet{}
	for index := 0; index < size; index++ {
		x.Add(rand.Intn(3000))
	}
	for i := 0; i < b.N; i++ {
		for index := 0; index < size/2; index++ {
			x.Has(rand.Intn(3000))
		}
	}
}
func BenchmarkIntSetHas1(b *testing.B)     { benchmarkIntSetHas(b, 1) }
func BenchmarkIntSetHas10(b *testing.B)    { benchmarkIntSetHas(b, 10) }
func BenchmarkIntSetHas100(b *testing.B)   { benchmarkIntSetHas(b, 100) }
func BenchmarkIntSetHas1000(b *testing.B)  { benchmarkIntSetHas(b, 1000) }
func BenchmarkIntSetHas10000(b *testing.B) { benchmarkIntSetHas(b, 10000) }

func benchmarkMapIntSetHas(b *testing.B, size int) {
	x := MapIntSet{}
	for index := 0; index < size; index++ {
		x.Add(rand.Intn(3000))
	}
	for i := 0; i < b.N; i++ {
		for index := 0; index < size/2; index++ {
			x.Has(rand.Intn(3000))
		}
	}
}
func BenchmarkMapIntSetHas1(b *testing.B)     { benchmarkMapIntSetHas(b, 1) }
func BenchmarkMapIntSetHas10(b *testing.B)    { benchmarkMapIntSetHas(b, 10) }
func BenchmarkMapIntSetHas100(b *testing.B)   { benchmarkMapIntSetHas(b, 100) }
func BenchmarkMapIntSetHas1000(b *testing.B)  { benchmarkMapIntSetHas(b, 1000) }
func BenchmarkMapIntSetHas10000(b *testing.B) { benchmarkMapIntSetHas(b, 10000) }
