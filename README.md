# WebSocket behind Nginx Example
Just a reminder of how configure Nginx so that it supports WebSocket connections.

### The problem
From the [nginx blog](https://www.nginx.com/blog/websocket-nginx):  

_There are some challenges that a reverse proxy server faces in supporting WebSocket. One is that WebSocket is a hop‑by‑hop protocol, so when a proxy server intercepts an Upgrade request from a client it needs to send its own Upgrade request to the backend server, including the appropriate headers. Also, since WebSocket connections are long lived, as opposed to the typical short‑lived connections used by HTTP, the reverse proxy needs to allow these connections to remain open, rather than closing them because they seem to be idle._  

### The solution

To make WebSocket connections are properly forwarded, `Upgrade` and `Connection` headers must be set explicitly:

``` nginx
# file: nginx/nginx.conf
# inside the http section:
map $http_upgrade $connection_upgrade {
    default upgrade;
    '' close;
}

# inside the server section:
location /ws {
    proxy_pass http://wsserver:8000;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection $connection_upgrade;
}
```

### Launch the example
```
docker-compose build
docker-compose up
open http://localhost
```
