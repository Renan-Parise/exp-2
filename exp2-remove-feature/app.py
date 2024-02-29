# main_app.py
from flask import Flask, render_template, request, redirect, url_for

app = Flask(__name__)

class TaskModel:
    def __init__(self):
        self.tasks = []
        self.task_id_counter = 1

    def add_task(self, description):
        task = {'id': self.task_id_counter, 'description': description, 'completed': False}
        self.tasks.append(task)
        self.task_id_counter += 1

    def get_tasks(self):
        return self.tasks

    def mark_task_as_completed(self, task_id):
        for task in self.tasks:
            if task['id'] == int(task_id):
                task['completed'] = not task['completed']

    def remove_task(self, task_id):
        self.tasks = [task for task in self.tasks if task['id'] != int(task_id)]

class TaskView:
    def render(self, tasks):
        return render_template('index.html', tasks=tasks)

class TaskController:
    def __init__(self, model, view):
        self.model = model
        self.view = view

    def show_tasks(self):
        tasks = self.model.get_tasks()
        return self.view.render(tasks)

    def add_task(self, description):
        self.model.add_task(description)

    def mark_task_as_completed(self, task_id):
        self.model.mark_task_as_completed(task_id)
    
    def remove_task(self, task_id):
        self.model.remove_task(task_id)

# O controlador é responsável por receber as interações do usuário, chamar métodos no modelo correspondente e atualizar a visualização.

model = TaskModel()
view = TaskView()
controller = TaskController(model, view)

@app.route('/')
def index():
    return controller.show_tasks()

@app.route('/add_task', methods=['POST'])
def add_task():
    description = request.form['task']
    controller.add_task(description)
    return redirect(url_for('index'))

@app.route('/mark_completed/<task_id>', methods=['POST'])
def mark_completed(task_id):
    controller.mark_task_as_completed(task_id)
    return redirect(url_for('index'))

@app.route('/remove_task/<task_id>', methods=['POST'])
def remove_task(task_id):
    controller.remove_task(task_id)
    return redirect(url_for('index'))

if __name__ == '__main__':
    app.run(debug=True)
