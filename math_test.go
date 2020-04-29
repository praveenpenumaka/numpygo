package numpygo

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMathLambdaExp(t *testing.T){
	ones := Ones("FLOAT64",10,10,10)
	logOnes := Exp(ones)
	assert.Equal(t,math.Exp(1),logOnes.Elements.Values[0])
}

func TestMathLambdaLog2(t *testing.T){
	ones := Ones("FLOAT64",10,10,10)
	logOnes := Log2(ones)
	assert.Equal(t,math.Log2(1),logOnes.Elements.Values[0])
}

func TestMathLambdaLog(t *testing.T){
	ones := Ones("FLOAT64",10,10,10)
	logOnes := Log(ones)
	assert.Equal(t,float64(0),logOnes.Elements.Values[0])
}

func TestMathLambdaMult(t *testing.T)  {
	ones := Ones("FLOAT64",10,10,10)
	logOnes := Mult(ones,5.2)
	assert.Equal(t, 5.2,logOnes.Elements.Values[0])
}

func TestMathLambdaDiv(t *testing.T) {
	ones := Ones("FLOAT64",10,10,10)
	logOnes := Div(ones,0.5)
	assert.Equal(t,float64(2),logOnes.Elements.Values[0])
}

func TestMathLambdaDivByZero(t *testing.T) {
	ones := Ones("FLOAT64",10,10,10)
	logOnes := Div(ones,0)
	assert.Equal(t,math.Inf(1),logOnes.Elements.Values[0])
}

func TestMathLambdaPow(t *testing.T) {
	ones := Ones("FLOAT64",10,10,10)
	logOnes := Pow(ones,10)
	assert.Equal(t,float64(1),logOnes.Elements.Values[0])
}

func TestMathLambdaEqual(t *testing.T) {
	ones := Ones("FLOAT64",10,10,10)
	logOnes := Equal(ones,0.5)
	assert.Equal(t,float64(0),logOnes.Elements.Values[0])
	logOnes = Equal(ones,float64(1))
	assert.Equal(t,float64(1),logOnes.Elements.Values[0])
}

func TestAdditionFailForNilInput(t *testing.T)  {
	a := Addition(nil,nil)
	assert.Nil(t,a)
}

func TestAdditionFailForIncompatibleSize(t *testing.T)  {
	a:= Ones("FLOAT64",10,5)
	b:= Ones("FLOAT64",10,6)
	c:= Addition(a,b)
	assert.Nil(t,c)
}

func TestAdditionFailForIncompatibleShape(t *testing.T)  {
	a:= Ones("FLOAT64",10,2)
	b:= Ones("FLOAT64",10,1,2)
	c:= Addition(a,b)
	assert.Nil(t,c)
}

func TestAdditionSuccess(t *testing.T)  {
	a:= Ones("FLOAT64",10,5)
	b:= Ones("FLOAT64",10,5)
	c:= Addition(a,b)
	assert.Equal(t,int(10),c.Shape.Values[0])
	assert.Equal(t,int(5),c.Shape.Values[1])
	assert.Equal(t,float64(2),c.Elements.Values[0])
}

func TestSubtractFailForNilInput(t *testing.T)  {
	a := Subtract(nil,nil)
	assert.Nil(t,a)
}

func TestSubtractFailForIncompatibleSize(t *testing.T)  {
	a:= Ones("FLOAT64",10,5)
	b:= Ones("FLOAT64",10,6)
	c:= Subtract(a,b)
	assert.Nil(t,c)
}

func TestSubtractFailForIncompatibleShape(t *testing.T)  {
	a:= Ones("FLOAT64",10,2)
	b:= Ones("FLOAT64",10,1,2)
	c:= Subtract(a,b)
	assert.Nil(t,c)
}

func TestSubtractSuccess(t *testing.T)  {
	a:= Ones("FLOAT64",10,5)
	b:= Ones("FLOAT64",10,5)
	c:= Subtract(a,b)
	assert.Equal(t,int(10),c.Shape.Values[0])
	assert.Equal(t,int(5),c.Shape.Values[1])
	assert.Equal(t,float64(0),c.Elements.Values[0])
}

func TestMultiplyFailForNilInput(t *testing.T)  {
	a := Multiply(nil,nil)
	assert.Nil(t,a)
}

func TestMultiplyFailForIncompatibleSize(t *testing.T)  {
	a:= Ones("FLOAT64",10,5)
	b:= Ones("FLOAT64",10,6)
	c:= Multiply(a,b)
	assert.Nil(t,c)
}

func TestMultiplyFailForIncompatibleShape(t *testing.T)  {
	a:= Ones("FLOAT64",10,2)
	b:= Ones("FLOAT64",10,1,2)
	c:= Multiply(a,b)
	assert.Nil(t,c)
}

func TestMultiplySuccess(t *testing.T)  {
	a:= Ones("FLOAT64",10,5)
	b:= Ones("FLOAT64",10,5)
	c:= Multiply(a,b)
	assert.Equal(t,int(10),c.Shape.Values[0])
	assert.Equal(t,int(5),c.Shape.Values[1])
	assert.Equal(t,float64(1),c.Elements.Values[0])
}

func TestDivisionFailForNilInput(t *testing.T)  {
	a := Division(nil,nil)
	assert.Nil(t,a)
}

func TestDivisionFailForIncompatibleSize(t *testing.T)  {
	a:= Ones("FLOAT64",10,5)
	b:= Ones("FLOAT64",10,6)
	c:= Division(a,b)
	assert.Nil(t,c)
}

func TestDivisionFailForIncompatibleShape(t *testing.T)  {
	a:= Ones("FLOAT64",10,2)
	b:= Ones("FLOAT64",10,1,2)
	c:= Division(a,b)
	assert.Nil(t,c)
}

func TestDivisionSuccess(t *testing.T)  {
	a:= Ones("FLOAT64",10,5)
	b:= Ones("FLOAT64",10,5)
	c:= Division(a,b)
	assert.Equal(t,int(10),c.Shape.Values[0])
	assert.Equal(t,int(5),c.Shape.Values[1])
	assert.Equal(t,float64(1),c.Elements.Values[0])
}