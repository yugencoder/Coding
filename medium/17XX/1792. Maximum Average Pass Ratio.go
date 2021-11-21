package _7XX

import (
	"container/heap"
)

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	res := 0.0
	heapData := &Heap{}

	for _, c := range classes {
		heap.Push(heapData, []float64{float64(c[0]), float64(c[1])})
	}

	for i := 0; i < extraStudents; i++ {
		c := heap.Pop(heapData).([]float64)
		c[0] += 1
		c[1] += 1

		heap.Push(heapData, c)
	}

	for _, v := range *heapData {
		res += v[0] / v[1]
	}
	return res / float64(len(classes))
}

type Heap [][]float64

func (p Heap) Len() int { return len(p) }
func (p Heap) Less(i, j int) bool {
	pi := p[i][0]
	ti := p[i][1]

	pj := p[j][0]
	tj := p[j][1]

	return (pi+1)/(ti+1)-pi/ti > (pj+1)/(tj+1)-pj/tj
	//return p[i][0] < p[i][1]
}
func (p Heap) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *Heap) Push(i interface{}) { *p = append(*p, i.([]float64)) }
func (p *Heap) Pop() interface{}   { v := (*p)[len(*p)-1]; *p = (*p)[:len(*p)-1]; return v }
