FROM python:latest
COPY ./flaskapi/api2 /app
WORKDIR /app
RUN pip install flask
EXPOSE 82
CMD ["python","api2.py"]