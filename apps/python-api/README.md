# Requirements and Installation

1. Create and activate a virtual environment (recommended):

```bash
python -m venv venv
source venv/bin/activate
```

2. Install the project dependencies:

```bash
pip install -r requirements.txt
```

If the `requirements.txt` file does not exist, you can create it by running:

```bash
pip install flask
pip freeze > requirements.txt
```

# Instructions to Run the Flask API

1. Install Flask if you don't have it:

```bash
pip install flask
```

2. Go to the directory where the `app.py` file is located:

```bash
cd apps/flask-api
```

3. Run the API:

```bash
python app.py
```

4. Test the endpoint in your browser or with curl:

- By default, the API runs at http://127.0.0.1:5000
- Example usage:

  - http://127.0.0.1:5000/greet?name=Rubén

  - Response:
    ```json
    {"message": "Hola Rubén"}
    ```