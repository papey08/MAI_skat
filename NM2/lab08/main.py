import math

import app.plotting as plotting
from app.analytical import analytial
from app.variable_directions import variable_directions
from app.fractional_steps import fractional_steps

x_begin = 0
x_end = math.pi

y_begin = 0
y_end = math.pi

t_begin = 0
t_end = 1

a = 1
mu1 = 1
mu2 = 1

h_x = 0.01
h_y = 0.01
tau = 0.01

def phi_0(y, t, a=a, mu1=mu1, mu2=mu2):
    return math.cos(mu2*y) * math.exp(-(mu1**2 + mu2**2) * a * t)

def phi_1(y, t, a=a, mu1=mu1, mu2=mu2):
    return (-1)**mu1 * math.cos(mu2*y) * math.exp(-(mu1**2 + mu2**2) * a * t)

def phi_2(x, t, a=a, mu1=mu1, mu2=mu2):
    return math.cos(mu1*x) * math.exp(-(mu1**2 + mu2**2) * a * t)

def phi_3(x, t, a=a, mu1=mu1, mu2=mu2):
    return (-1)**mu2 * math.cos(mu1*x) * math.exp(-(mu1**2 + mu2**2) * a * t)

def psi(x, y, mu1=mu1, mu2=mu2):
    return math.cos(mu1*x) * math.cos(mu2*y)

def solution(x, y, t, a=a, mu1=mu1, mu2=mu2):
    return math.cos(mu1*x) * math.cos(mu2*y) * math.exp(-(mu1**2 + mu2**2) * a * t)


if __name__ == '__main__':

    analytical_solution = analytial(
        (x_begin, x_end),
        (y_begin, y_end),
        (t_begin, t_end),
        h_x, h_y, tau,
        solution,
    )

    variable_directions_solution = variable_directions(
        (x_begin, x_end),
        (y_begin, y_end),
        (t_begin, t_end),
        h_x, h_y, tau, a, mu1, mu2,
        phi_0, phi_1, phi_2, phi_3, psi,
    )

    fractional_steps_solution = fractional_steps(
        (x_begin, x_end),
        (y_begin, y_end),
        (t_begin, t_end),
        h_x, h_y, tau, a, mu1, mu2,
        phi_0, phi_1, phi_2, phi_3, psi,
    )

    plotting.plot_results(
        analytical_solution,
        fractional_steps_solution,
        variable_directions_solution,
        0.5, 0.5,
        (x_begin, x_end),
        (y_begin, y_end),
        (t_begin, t_end),
        h_x, h_y, tau
    )

    plotting.plot_errors_from_time(
        analytical_solution,
        fractional_steps_solution,
        variable_directions_solution,
        (t_begin, t_end),
        tau,
    )
