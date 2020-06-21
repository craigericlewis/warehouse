FROM python:3

ENV APP_HOME /usr/src/app
WORKDIR $APP_HOME

COPY . ./

RUN pip3 install -r requirements.txt

# Development purposes
# CMD ["python", "-u", "-m", "flask", "run", "--host", "0.0.0.0"]
CMD ["gunicorn", "builder:app", "-w", "2", "--threads", "2", "-b 0.0.0.0:5000"]
