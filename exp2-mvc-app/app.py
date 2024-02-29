from flask import Flask, render_template, request
from model.model import CalculatorModel
from view.view import CalculatorView

app = Flask(__name__)

@app.route('/', methods=['GET', 'POST'])
def index():
    if request.method == 'POST':
        x = int(request.form['x'])
        y = int(request.form['y'])
        operation = request.form['operation']

        model = CalculatorModel()
        if operation == 'add':
            model.add(x, y)
        elif operation == 'subtract':
            model.subtract(x, y)
        
        result = model.result
        return render_template('index.html', result=result)
    return render_template('index.html', result=None)

if __name__ == '__main__':
    app.run(debug=True)
