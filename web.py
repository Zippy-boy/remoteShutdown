from flask import Flask, request

app = Flask(__name__)

@app.route('/')
def home():
    return '''
        <html>
            <body>
                <form action="/shutdown" method="post">
                    <label for="key">Enter the shutdown key:</label>
                    <input type="password" id="key" name="key">
                    <input type="submit" value="Shutdown">
                </form>
            </body>
        </html>
    '''

@app.route('/shutdown', methods=['POST'])
def shutdown():
    key = request.form['key']
    if key == '1234':
        # code to shutdown the system goes here
        return 'System shutting down...'
    else:
        return 'Incorrect key'

if __name__ == '__main__':
    app.run(debug=True)
