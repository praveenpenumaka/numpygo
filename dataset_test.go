package numpygo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadDataWithEmptyName(t *testing.T) {
	dataSet := LoadData("")
	assert.Equal(t, 0, dataSet.Data.Size)
}

func TestLoadDataWithIris(t *testing.T) {
	dataSet := LoadData("iris")
	assert.Equal(t, 600, dataSet.Data.Size)
	assert.Equal(t, 150, dataSet.Target.Size)
}

func TestLoadDataWithBostonHousePrices(t *testing.T) {
	dataSet := LoadData("boston_house_prices")
	assert.Equal(t, 6578, dataSet.Data.Size)
	assert.Equal(t, 506, dataSet.Target.Size)
}

func TestGetRootPath(t *testing.T) {
	path := getRootPath()
	assert.NotEmpty(t, path)
}

func TestLoadCSVWithEmptyData(t *testing.T) {
	a1, a2, e := loadCSV("")
	assert.Equal(t, 0, a1.Size)
	assert.Equal(t, 0, a2.Size)
	assert.Equal(t, 0, len(e))
}

func TestLoadCSVWithInvalidRows(t *testing.T) {
	a1, a2, e := loadCSV("abc,def")
	assert.Equal(t, 0, a1.Size)
	assert.Equal(t, 0, a2.Size)
	assert.Equal(t, 0, len(e))
}
