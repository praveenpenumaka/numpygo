package benchmark

/*
func getRandomDims(size,x,y int) (dims []domain.Tuple){
	dims = make([]domain.Tuple,size)
	for i, _ := range dims {
		dims[i] =  domain.Tuple{
			X: x,
			Y: y,
		}
	}
	return dims
}

func BenchmarkNdArrayGet(t *testing.B){
	arr := numpygo.Zeros("FLOAT64",10,10,10,10,10,10,10)
	dims := getRandomDims(arr.Dims,0,9)
	currentTime := time.Now()
	for i:=0;i<t.N;i++{
		arr.Get(dims)
	}
	finalTime := time.Now()
	diff := currentTime.Sub(finalTime)
	fmt.Sprintf("Time consumed:%d",diff.Milliseconds())
}

func init(){

}
*/
