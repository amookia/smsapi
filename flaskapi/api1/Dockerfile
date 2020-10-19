FROM python:latest
COPY ./flaskapi/api1 /app
WORKDIR /app
RUN pip install flask
EXPOSE 81
CMD ["python","api1.py"]