### 自动部署静态活动 

### 初始化
1.go mod init main_web

2.run and go get -t -v github.com/gin-gonic/gin

### 直接运行
http://127.0.0.1:8093/auto_deploy


### nginx 反向代理配置
    server {
        keepalive_requests 120; #单连接请求上限次数。
        listen       80;
        server_name  localhost;
     
        location  /admin/ {
            proxy_pass http://127.0.0.1:8093/auto_deploy;
            index   index.html;
        }

        location /auto_deploy/ {
               proxy_pass http://127.0.0.1:8093;
        }

        location ^~ /auto_deploy_api {
            proxy_pass http://127.0.0.1:8093;
        }

        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
    }
http://127.0.0.1/admin/
