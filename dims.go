package numpygo

import (
	"github.com/praveenpenumaka/numpygo/domain"
	"strconv"
	"strings"
)

func getDim(ident string,size int) *domain.Tuple{
	if ident  ==  ":"{
		return &domain.Tuple{
			X: 0, Y: size-1,
		}
	}else{
		if v, err := strconv.Atoi(ident); err == nil {
			return &domain.Tuple{
				X: v,
				Y: v,
			}
		}else{
			x:=strings.Split(ident,":")
			start:=0
			if v, err := strconv.Atoi(x[0]); err == nil {
				start=v
			}
			end:=size-1
			if v, err := strconv.Atoi(x[1]); err == nil {
				end=v
			}
			return &domain.Tuple{
				X: start,
				Y: end,
			}
		}
	}
}

func ParseDim(sDims string, shape *domain.IVector) *domain.Dimensions {
	if sDims == "" || shape == nil{
		return nil
	}
	if shape.Mult() == 0 {
		return nil
	}
	if sDims == ":" {
		if shape == nil {
			return nil
		}
		dims := &domain.Dimensions{Dims: nil}
		for _, d := range shape.Values {
			tuple := &domain.Tuple{
				X: 0, Y: d-1,
			}
			dims.Dims = append(dims.Dims, *tuple)
		}
		return dims
	}else if strings.Contains(sDims,","){
		dimIdents := strings.Split(sDims,",")
		dims := &domain.Dimensions{Dims: nil}
		if len(dimIdents) == len(shape.Values){
			for i, dimIdent := range dimIdents {
				tup := getDim(dimIdent,shape.Values[i])
				dims.Dims=append(dims.Dims, *tup)
			}
		}
		return dims
	}
	return nil
}
