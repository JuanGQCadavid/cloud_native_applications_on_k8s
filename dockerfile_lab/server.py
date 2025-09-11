import socket

from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route("/")
def hello():
    response = "Client IP: " + request.remote_addr + "\nHostname: " + socket.gethostname() + "\n"
    return response, 200


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)