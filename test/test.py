import numpy as np
from scipy.stats import norm

data_points = np.array([0, 0.5, 1, 2, 3, 4])
params = {
    # Calculate mean for cluster 1
    'mean1': np.mean(data_points[:2]),
    # Calculate mean for cluster 2
    'mean2': np.mean(data_points[2:]),
    # Calculate variance for cluster 1
    'variance1': np.var(data_points[:2]),
    # Calculate variance for cluster 2
    'variance2': np.var(data_points[2:]),
    # Assume equal distribution for cluster 1 and cluster 2
    'mixing_coefficient1': 0.5,
    'mixing_coefficient2': 0.5
}

# Calculate the Gaussian Mixture Model probability density function for each data point
def gmm_pdf(x, params):
    return (params['mixing_coefficient1'] * norm.pdf(x, params['mean1'], np.sqrt(params['variance1'])) +
            params['mixing_coefficient2'] * norm.pdf(x, params['mean2'], np.sqrt(params['variance2'])))

# Calculate the log likelihood
def log_likelihood(data_points, params):
    return -np.sum(np.log(gmm_pdf(data_points, params)))

# Calculate initial E_in
value = log_likelihood(data_points, params)
print(value)
