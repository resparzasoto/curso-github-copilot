# Requerimientos e instalación

1. Crea y activa un entorno virtual (recomendado):

```bash
python -m venv venv
source venv/bin/activate
```

2. Instala las dependencias del proyecto:

```bash
pip install -r requirements.txt
```

Si no existe el archivo `requirements.txt`, puedes crearlo ejecutando:

```bash
pip install flask
pip freeze > requirements.txt
```

# Instrucciones para ejecutar la API Flask

1. Instala Flask si no lo tienes:

```bash
pip install flask
```

2. Ve al directorio donde está el archivo `app.py`:

```bash
cd apps/flask-api
```

3. Ejecuta la API:

```bash
python app.py
```

4. Prueba el endpoint en tu navegador o con curl:

- Por defecto, la API corre en http://127.0.0.1:5000
- Ejemplo de uso:

  - http://127.0.0.1:5000/greet?name=Rubén

  - Respuesta:
    ```json
    {"message": "Hola Rubén"}
    ```