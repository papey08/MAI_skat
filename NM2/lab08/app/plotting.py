import numpy as np

import matplotlib.pyplot as plt

def max_abs_error(A, B):
    return abs(A - B).max()

def mean_abs_error(A, B):
    return abs(A - B).mean()


def plot_results(analytical, fractional_steps, variable_directions, cur_time, cur_y, x_range, y_range, t_range, h_x, h_y, tau):
    x = np.arange(*x_range, h_x)
    y = np.arange(*y_range, h_y)
    t = np.arange(*t_range, tau)
    cur_t_id = abs(t - cur_time).argmin()
    cur_y_id = abs(y - cur_y).argmin()

    plt.figure(figsize=(15, 9))
    # for method_name, solution in solutions.items():
    plt.plot(x, analytical[cur_t_id][:, cur_y_id], label='anaytical')
    plt.plot(x, fractional_steps[cur_t_id][:, cur_y_id], label='fractional steps')
    plt.plot(x, variable_directions[cur_t_id][:, cur_y_id], label='variable directions')

    plt.legend()
    plt.grid()
    plt.show()


def plot_errors_from_time(analytical, fractional_steps, variable_directions, t_range, tau):
    t = np.arange(*t_range, tau)

    plt.figure(figsize=(15, 9))
    # for method_name, solution in solutions.items():
    #     if method_name == analytical_solution_name:
    #         continue
    #     max_abs_errors = np.array([
    #         max_abs_error(solution[i], solutions[analytical_solution_name][i])
    #         for i in range(len(t))
    #     ])
    #     plt.plot(t, max_abs_errors, label=method_name)
    fractional_steps_max_abs_errors = np.array([
        max_abs_error(fractional_steps[i], analytical[i])
        for i in range(len(t))
    ])
    plt.plot(t, fractional_steps_max_abs_errors, label='fractional steps')

    variable_directions_max_abs_errors = np.array([
        max_abs_error(variable_directions[i], analytical[i])
        for i in range(len(t))
    ])
    plt.plot(t, variable_directions_max_abs_errors, label='variable directions')

    plt.xlabel('time')
    plt.ylabel('Max abs error')

    plt.legend()
    plt.grid()
    plt.show()
