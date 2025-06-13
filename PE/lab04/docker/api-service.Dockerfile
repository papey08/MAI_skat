FROM python:3.13.2
WORKDIR /app
COPY ../src .
RUN pip install --no-cache-dir -r requirements.txt
ENV PYTHONPATH="./"
EXPOSE 8099
CMD ["python", "./api/app.py"]
