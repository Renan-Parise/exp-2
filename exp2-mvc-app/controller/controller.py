from model import CalculatorModel

class CalculatorController:
    def __init__(self, model):
        self.model = model
    
    def perform_operation(self, operation, x, y):
        if operation == 'add':
            self.model.add(x, y)
        elif operation == 'subtract':
            self.model.subtract(x, y)
