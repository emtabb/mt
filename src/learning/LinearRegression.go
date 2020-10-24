package learning

import (
	"math"
	"github.com/emtabb/field/interaction"
	. "github.com/emtabb/field/src/calculation/space"
	"github.com/emtabb/field/src/calculation/operator"
)

type LinearRegression struct {
	interaction.Interaction
	vectors []Vector
	expect Vector
	coef_ float64
	intercept_ []float64
	err []float64

	mae float64
	mse float64
	rmse float64
	mape float64
	mpe float64

	isGenerate bool
}

func (lin *LinearRegression) generate() {
	if !lin.isGenerate {
		lin.coef_ = 0
		lin.intercept_ = make([]float64, 0)
		lin.err = make([]float64, 0)
		lin.vectors = make([]Vector, 0)
		lin.isGenerate = true
	}
}

func (lin *LinearRegression) Init(vectors []Vector) {
	lin.vectors = vectors[:]
}

func (lin *LinearRegression) estimate() {
	// len_data := len(data)
	// len_expect := len(expect)
	// if len_data != len_expect {
	// 	return 0.0, 0.0
	// }
	spaceDimension := len(lin.vectors)
	m_x := make([]float64, spaceDimension)
	expect := lin.expect.Coordinate()
	m_y := operator.Mean(expect, 4)
	SS_xy := make([]float64, spaceDimension)
	SS_xx := make([]float64, spaceDimension)
	for i := 0; i < spaceDimension; i++ {
		data := lin.vectors[i].Coordinate()
		SS_xy[i] = operator.Multi(data, expect, 4) - float64(len(data)) * m_y * m_x[i]
		SS_xx[i ]= operator.Multi(data, data, 4) - float64(len(data)) * m_x[i] * m_x[i]
		lin.intercept_[i] = SS_xy[i] / SS_xx[i]
	}
	
	for i := 0; i < spaceDimension; i++ {
		lin.coef_ += m_y - lin.intercept_[i] * m_x[i] 
	}
}

func (lin *LinearRegression) Predict(predictData []float64) (float64, float64) {
	/*	*NOTE* 
	* - MAE : Mean Absolute Error
	* - MSE : Mean Squared Error
	* - RMSE : Root Mean Squared Error
	* - MAPE : Mean Absolute Percentage Error
	* - MPE : Mean Percentage Error
	*/
	
	expect := lin.expect.Coordinate()
	y_predict := 0.0
	for i := 0; i < len(predictData); i++ {
		y_predict += predictData[i] * lin.intercept_[i]
	}
	y_predict += lin.coef_

	mae := 0.0
	for i := 0; i < len(expect); i++ {
		mae += math.Abs(expect[i] - y_predict)
	}
	mae = mae / float64(len(expect))
	return y_predict, mae
}

func (lin *LinearRegression) Generate() []Vector {

	return lin.vectors
}