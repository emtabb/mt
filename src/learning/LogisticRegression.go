package learning

import (
	"math"
	"github.com/emtabb/field/interaction"
	//. "github.com/emtabb/field/src/calculation/space"
	//"github.com/emtabb/field/src/calculation/operator"
)

type LogisticRegression struct {
	interaction.Interaction
}


func (log *LogisticRegression) logisticFunction(x float64) float64 {
	x = -x
	return 1.0 / ( 1 + math.Exp(x) )
}

// https://www.kaggle.com/kanncaa1/deep-learning-tutorial-for-beginners#Logistic-Regression
// # Forward propagation steps:
// # find z = w.T*x+b
// # y_head = sigmoid(z)
// # loss(error) = loss(y,y_head)
// # cost = sum(loss)
// def forward_propagation(w,b,x_train,y_train):
//     z = np.dot(w.T,x_train) + b
//     y_head = sigmoid(z) # probabilistic 0-1
//     loss = -y_train*np.log(y_head)-(1-y_train)*np.log(1-y_head)
//     cost = (np.sum(loss))/x_train.shape[1]      # x_train.shape[1]  is for scaling
//     return cost 


/*

# In backward propagation we will use y_head that found in forward progation
# Therefore instead of writing backward propagation method, lets combine forward propagation and backward propagation
def forward_backward_propagation(w,b,x_train,y_train):
    # forward propagation
    z = np.dot(w.T,x_train) + b
    y_head = sigmoid(z)
    loss = -y_train*np.log(y_head)-(1-y_train)*np.log(1-y_head)
    cost = (np.sum(loss))/x_train.shape[1]      # x_train.shape[1]  is for scaling
    # backward propagation
    derivative_weight = (np.dot(x_train,((y_head-y_train).T)))/x_train.shape[1] # x_train.shape[1]  is for scaling
    derivative_bias = np.sum(y_head-y_train)/x_train.shape[1]                 # x_train.shape[1]  is for scaling
    gradients = {"derivative_weight": derivative_weight,"derivative_bias": derivative_bias}
    return cost,gradients

*/