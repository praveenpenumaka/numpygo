package numpygo

import (
	"github.com/praveenpenumaka/numpygo/domain"
	"io/ioutil"
	"path"
	"runtime"
	"strconv"
	"strings"
)

type DataSet struct {
	Data         NDArray
	Target       NDArray
	TargetNames  []string
	FeatureNames []string
	DatasetPath  string
	Description  string
}

//TODO: Add tests
func loadCSV(data string) (NDArray, NDArray, []string) {
	rows := strings.Split(data, "\n")
	rowCount := len(rows)
	if rowCount < 2 {
		return NDArray{}, NDArray{}, nil
	}
	header := strings.Split(rows[0], ",")
	headerLen := len(header)
	if headerLen <= 1 {
		return NDArray{}, NDArray{}, nil
	}
	nSamples, sampleError := strconv.ParseInt(header[0], 10, 64)
	if sampleError != nil {
		return NDArray{}, NDArray{}, nil
	}

	var targetNames []string
	if headerLen > 2 {
		for i := 2; i < headerLen; i++ {
			targetNames = append(targetNames, header[i])
		}
	}

	rowLength := len(strings.Split(rows[1], ","))
	if rowLength <= 0 {
		return NDArray{}, NDArray{}, nil
	}

	sampleShape := &domain.IVector{Values: []int{int(nSamples), rowLength - 1}}
	featureShape := &domain.IVector{Values: []int{int(nSamples), 1}}
	samples := NewNDArray("FLOAT64", sampleShape.Values)
	features := NewNDArray("FLOAT64", featureShape.Values)

	sampleIndex := 0
	featureIndex := 0
	for i := 1; i < rowCount; i++ {
		cols := strings.Split(rows[i], ",")
		if len(cols) != rowLength {
			return NDArray{}, NDArray{}, nil
		}
		for index, col := range cols {
			v, e := strconv.ParseFloat(col, 64)
			if e != nil {
				return NDArray{}, NDArray{}, nil
			}
			if index == rowLength-1 {
				features.Elements.Values[sampleIndex] = v
				sampleIndex++
			} else {
				samples.Elements.Values[featureIndex] = v
				featureIndex++
			}
		}
	}

	return samples, features, targetNames
}

//TODO: Add tests
func LoadDataFromFile(path string, featureNames []string) DataSet {
	if path == "" {
		return DataSet{}
	}
	CSVdata, fileReadError := ioutil.ReadFile(path)
	if fileReadError != nil {
		return DataSet{}
	}
	data, target, targetNames := loadCSV(string(CSVdata))
	return DataSet{
		Data:         data,
		Target:       target,
		TargetNames:  targetNames,
		FeatureNames: featureNames,
		DatasetPath:  path,
		Description:  "",
	}
}

//TODO: Add tests
func getRootPath() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return d
}

//TODO: Add tests
func LoadData(datasetName string) DataSet {
	var featureNames []string
	currentDir := getRootPath()
	basePath := currentDir + "/data/"
	dataSetBasePath := basePath + datasetName
	if datasetName == "iris" {
		featureNames = []string{"sepal length (cm)", "sepal width (cm)",
			"petal length (cm)", "petal width (cm)"}
	}
	dataPath := dataSetBasePath + "/data.csv"
	return LoadDataFromFile(dataPath, featureNames)
}
