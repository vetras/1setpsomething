
Build:

    $ docker build -t echo-go:v1 .

Run:
    $ docker run -d -p 8000:8000 --name echo-go echo-go:v1
