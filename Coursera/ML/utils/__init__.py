import numpy as np
import matplotlib.pyplot as plt

def draw_J_history(history):
    iterations = np.arange(history.shape[0])

    plt.plot(iterations,history)
    plt.xlabel("Iterations")
    plt.ylabel("J")
    plt.title(label="Yo")
    plt.show()
    return

def draw_hypothesis(X,y,theta):
    print(X.shape,theta.shape)
    
    plt.plot(X,y,"bx", label = "Training examples")
    plt.plot(X, X @ theta, "r", label = "Hypothesis")
    plt.legend()
    plt.title(label = "Dataset representation with hypothesis")
    plt.show()
    return

def feature_scaling(X):
    X_norm = X
    
    mu = X_norm.mean(0)
    sigma = X_norm.std(0)
    
    X_norm = (X_norm - mu) / sigma
    
    return X_norm, mu, sigma
def add_bias(matrix, direction = 1):
    m,n = matrix.shape
    
    shape = (m,1)
    if direction == 0:
        shape = (1,n)

    return np.concatenate((np.ones(shape),matrix),direction) 

def optimize_gradient(X, y, init_theta, 
                      cost_func,
                      learning_rate, 
                      max_iterations = 1000,
                      regularization = None,
                      scaling=False):
    J = 0
    J_history = []
    theta = init_theta
    theta_history = np.zeros((max_iterations,theta.shape[0]))
    
    if scaling == True:
        X_no_bias, mu, sigma = feature_scaling(X[:,1:])
        X_local = add_bias(X_no_bias)
    else:
        X_local = X
        
    for i in range(max_iterations):
        J, dJ = cost_func(X_local, y, theta, regularization = regularization)
        J_history.append(J)
        
        theta = theta - learning_rate * dJ
        theta_history[i,:] = theta.reshape(init_theta.shape[0])
    
    return (theta, J, J_history, theta_history)