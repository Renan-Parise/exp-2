from flask import Flask, render_template
import requests

app = Flask(__name__)

def consultar_paises(query):
    url = 'https://countries.trevorblades.com/'
    headers = {'Content-Type': 'application/json'}
    data = {'query': query}

    response = requests.post(url, headers=headers, json=data)
    
    if response.status_code == 200:
        return response.json()
    else:
        print("Erro ao consultar a API:", response.status_code)
        return None

@app.route('/')
def index():
    query = '''
    query {
        countries {
            code
            name
            continent {
                name
            }
            languages {
                name
            }
        }
    }
    '''
    
    resultado = consultar_paises(query)
    
    if resultado:
        paises = resultado['data']['countries']
        return render_template('index.html', paises=paises)
    else:
        return "Erro ao consultar a API"

if __name__ == '__main__':
    app.run(debug=True)
