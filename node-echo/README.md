
Build:

    $ docker build -t echo-node:v1 .

Run:
    $ docker run -d -p 8001:8001 --name echo-node echo-node:v1
