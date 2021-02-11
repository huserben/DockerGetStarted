import os
from urllib.request import urlopen

from flask import Flask
app = Flask(__name__)


@app.route('/')
def hello_world():
    rest_api = os.environ.get('REST_API')
    if (rest_api == None):
        return "<h1>No REST API set!</h1>"
    else:
        print(rest_api)
        response = urlopen(rest_api)
        return "<h1>{0}</h1>".format(response.read().decode('utf-8'))

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')
