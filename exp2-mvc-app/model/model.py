class CalculatorModel:
    def __init__(self):
        self.result = 0
    
    def add(self, x, y):
        self.result = x + y
    
    def subtract(self, x, y):
        self.result = x - y
