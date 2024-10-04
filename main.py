import os

import requests
from flask import Flask, jsonify, Response

app = Flask(__name__)

def runShell(c:str)->str:
    return os.popen(c).read()

def getSystemStats() -> dict[str, str]:
    return  {
        "ip":  runShell("hostname -I").split(sep=" ")[0].rstrip(),
        "ps": runShell("ps -ax").split(sep='\n')[1:],
        "df": runShell("df -h \\/").split(sep='\n')[1].rstrip(),
        "uptime": runShell("uptime -p").rstrip()
}

def reqService2SystemStats() -> any:
    try:
        r = requests.get("http://service_2_go:8220")
    except Exception as e:
        return {"error": str(e)}
    return r.json()

@app.route('/', methods=['GET'])
def GET_SystemStats() -> Response:
    return jsonify({
        "service_1_python": getSystemStats(),
        "service_2_go": reqService2SystemStats()
    })

if __name__ == '__main__':
    app.run(host="0.0.0.0", port=8199, debug=False)
