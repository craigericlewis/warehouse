FROM python:3

ENV APP_HOME /usr/src/app
WORKDIR $APP_HOME

COPY . ./

RUN pip3 install -r requirements.txt

# CMD ["python", "-u", "-m", "flask", "run", "--host", "0.0.0.0"]
ENTRYPOINT ["./gunicorn_starter.sh"]
