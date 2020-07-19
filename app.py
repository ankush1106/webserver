from flask import Flask, jsonify, request
import requests

app = Flask(__name__)


def paginatedTODOs(page):
    page = int(page)
    items_per_page = 10
    r = requests.get("https://jsonplaceholder.typicode.com/todos/")
    data = r.json()
    items = []
    has_next = True
    if items_per_page*page <= len(data):
        if items_per_page*page == len(data):
            has_next = False
        items = data[(page-1)*items_per_page: page*items_per_page]
    else:
        has_next=False
    return items, has_next

@app.route('/todos')
def getTODO():
    items, has_next = paginatedTODOs(request.args.get("page", 1))
    return jsonify({"data":items, "has_next":has_next})


if __name__ == '__main__':
    app.run(host="0.0.0.0", port=3001)