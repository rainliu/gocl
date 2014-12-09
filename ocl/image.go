// +build cl11 cl12

package ocl

type Image interface {
	Memory
}

type image struct {
	memory
}
