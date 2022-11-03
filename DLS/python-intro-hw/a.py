from functools import reduce


class Neuron:
    def __init__(self, w, f = lambda x: x):
        self.weights = w
        self.f = f
        self.x = None

    def forward(self, x):
        self.x = x
        return reduce(lambda a, b: (a if type(a) == int else a[0] * a[1]) + b[0] * b[1], list(zip(self.weights, x)))

    def backlog(self):
        return self.x
