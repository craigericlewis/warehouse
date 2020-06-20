FROM python:3

ENV APP_HOME /usr/src/app
WORKDIR $APP_HOME
COPY . ./

EXPOSE 5000

RUN pip install flask

CMD ["python", "-u", "-m", "flask", "run", "--host", "0.0.0.0"]
