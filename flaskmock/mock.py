from flask import Flask, json, request
from pprint import pprint as pp

companies = [{"id": 1, "name": "Company One"}, {"id": 2, "name": "Company Two"}]

api = Flask(__name__)

@api.route('/dummy', methods=['GET'])
def get_companies():
  return json.dumps(companies)

@api.route('/', methods=['POST'])
def post_companies():
  request_data = request.get_json()
  pp(request.headers)
  pp(request_data)  
  f = open("out.json", "r")
  return f.read(), 201

if __name__ == '__main__':
    api.run()
