from flask import Flask,request,jsonify


api = Flask(__name__)


@api.route('/send',methods=["GET"])
def send():
    getargs = request.args
    formvalid = 'body' in getargs and 'number' in getargs
    if formvalid :
        try:
            int(request.args['number'])
            return jsonify({'message':request.args['number']}),200
        except:
            return jsonify({'message':'wrong number'}),400
    return jsonify({'message':'invalid form'}),400


if __name__ == '__main__':
    api.run(host='0.0.0.0',port=82,debug=False)
