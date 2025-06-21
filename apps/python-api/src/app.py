from flask import Flask, request, jsonify
from flasgger import Swagger

app = Flask(__name__)
Swagger(app)

@app.route('/greet', methods=['GET'])
def greet():
    """
    Greet endpoint
    ---
    parameters:
      - name: name
        in: query
        type: string
        required: false
        description: The name to greet
    responses:
      200:
        description: Returns a greeting message
        schema:
          type: object
          properties:
            message:
              type: string
    """
    name = request.args.get('name', default='World')
    message = f"Hello {name}"
    return jsonify({'message': message})

if __name__ == '__main__':
    app.run(debug=True)
