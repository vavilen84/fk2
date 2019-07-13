Role Nginx
=========

Role installs [nginx](http://nginx.org)

Requirements
------------

Ubuntu trusty. Probably should work on other Ubuntu versions (not tested)

Role Variables
--------------

Available variables are listed in default values (see [defaults/main.yml](defaults/main.yml) ).

### General nginx variables

| Variable | Description |
| -------- | -------- |
| **nginx_version**   | available versions listed here http://nginx.org/en/download.html . For example: **1.8.0** - stable version or **1.9.9** - mainline version |
| **nginx_user**   | sets the name of an unprivileged user whose credentials will be used by worker processes |
| **nginx_group**   | sets the name of a group whose credentials will be used by worker processes |
| **nginx_worker_processes**   | defines the number of worker processes. See [worker_processes](http://nginx.org/en/docs/ngx_core_module.html#worker_processes). Recomended value: grep processor /proc/cpuinfo | wc -l |
| **nginx_worker_connections**   | sets the maximum number of simultaneous connections that can be opened by a worker process. See [worker_connections](http://nginx.org/en/docs/ngx_core_module.html#worker_connections). Recomended value: ulimit -n  |
| **nginx_client_max_body_size**   | sets the maximum allowed size of the client request body, specified in the “Content-Length” request header field. See [client_max_body_size](http://nginx.org/en/docs/http/ngx_http_core_module.html#client_max_body_size) |
| **nginx_keepalive_requests**   | sets the maximum number of requests that can be served through one keep-alive connection. See [keepalive_requests](http://nginx.org/en/docs/http/ngx_http_core_module.html#keepalive_requests) |
| **nginx_keepalive_timeout**   | the first parameter sets a timeout during which a keep-alive client connection will stay open on the server side. See [keepalive_timeout](http://nginx.org/en/docs/http/ngx_http_core_module.html#keepalive_timeout) |
| **nginx_client_header_timeout**   | defines a timeout for reading client request header. See [client_header_timeout](http://nginx.org/en/docs/http/ngx_http_core_module.html#client_header_timeout) |
| **nginx_client_body_timeout**   | defines a timeout for reading client request body. See [client_body_timeout](http://nginx.org/en/docs/http/ngx_http_core_module.html#client_body_timeout) |
| **nginx_send_timeout**   | sets a timeout for transmitting a response to the client. See [send_timeout](http://nginx.org/en/docs/http/ngx_http_core_module.html#send_timeout)  |
| **nginx_fastcgi_connect_timeout**   | defines a timeout for establishing a connection with a FastCGI server. See [fastcgi_connect_timeout](http://nginx.org/en/docs/http/ngx_http_fastcgi_module.html#fastcgi_connect_timeout) |
| **nginx_fastcgi_send_timeout**   | sets a timeout for transmitting a request to the FastCGI server. See [fastcgi_send_timeout](http://nginx.org/en/docs/http/ngx_http_fastcgi_module.html#fastcgi_send_timeout) |
| **nginx_php_unix_socket**   | path to php unix socket. Use /var/run/php5-fpm.sock for php 5.x and /var/run/php/php7.0-fpm.sock for php 7.x    |

### nginx modules

Support for **ngx_headers_more** module https://github.com/openresty/headers-more-nginx-module

| Variable | Description |
| -------- | -------- |
| **nginx_include_headers_more_module**   | add support for module. Available values: true or false |
| **nginx_headers_more_module_version**   | available versions listed [here](https://github.com/openresty/headers-more-nginx-module/releases). Set version in format: 0.29 |
| **nginx_headers_more_module_base_url**  | url to module archive |

### HTTP Basic authentication

| Variable | Description |
| -------- | -------- |
| **nginx_http_basic_auth_file**   | path to http basic auth file |
| **nginx_http_basic_auth_users**  | associative array with users |

**nginx_http_basic_auth_users** example:
```yml
[
  { user: username, password: user_password }
]
```

### SSL

| Variable | Description |
| -------- | -------- |
| **nginx_certs_folder**   | path to ssl certs |
| **nginx_copy_ssl_certs_files**   | list with paths to files |

 **nginx_copy_ssl_certs_files** example:
```yml
[
      certs/server.cer,
      certs/server.csr,
      certs/server.key,
      certs/CA_bundle.pem
]
```

### hosts

| Variable | Description |
| -------- | -------- |
| **nginx_vhosts**   | associative array with domains to handle requests |

**nginx_vhosts** example:
```yml
[
    { server_name: '_', root: '/var/www', template: 'catch_all'},
    { server_name: 'db.example.com', root: '/usr/share/phpmyadmin', template: 'php'}
    { server_name: 'backend.example.com', root: '/var/www/backend.example.com', template: 'ssl_yii', basic_auth: true, ssl_certificate_file: 'example.com.crt', ssl_certificate_key_file: 'example.com.key', ssl_dhparam_file: 'dhparam.pem' },
    { server_name: 'storage.example.com', root: '/var/www/storage', template: 'ssl_static'},
]
```
|

Examples
------------

### Set all requests to /var/www folder by default

* add to **nginx_vhosts**

```yml
[
    { server_name: '_', root: '/var/www', template: 'catch_all'},
]
```

### How to serve static files (archives, images, videos, etc.)

* add to **nginx_vhosts**

```yml
[
    { server_name: 'storage.example.com', root: '/var/www/storage', template: 'static'}
]
```

### How to use PhpMyAdmin

* install PhpMyAdmin
* add to **nginx_vhosts**

```yml
[
    { server_name: 'db.example.com', root: '/usr/share/phpmyadmin', template: 'php'}
]
```

### How to use HTTP Basic authentication

For example, how to use PhpMyAdmin with HTTP Basic authentication

* install PhpMyAdmin
* add to **nginx_vhosts**

```yml
[
    { server_name: 'db.example.com', root: '/usr/share/phpmyadmin', template: 'php',  basic_auth: true }
]
```

* add to **nginx_http_basic_auth_users** users and passwords

```yml
[
  { user: john, password: john_secret_pass },
  { user: jack, password: jack_secret_pass },
]
```

### How to use Yii/Yii 2

* add to **nginx_vhosts**

```yml
[
    { server_name: 'yii-backend.example.com', root: '/var/www/yii_project/backend/web', template: 'yii'},
    { server_name: 'yii-frontend.example.com', root: '/var/www/yii_project/frontend/web', template: 'yii'},
    { server_name: 'yii-api.example.com', root: '/var/www/yii_project/api/web', template: 'yii'},
]
```

yii-backend.example.com domain goes to /var/www/yii_project/backend/web folder

yii-frontend.example.com domain goes to /var/www/yii_project/frontend/web folder and so on

**template** can be **yii** or **yii_ssl** (to use https)


### How to use HTTPS

* add to **nginx_vhosts**

```yml
[
    { server_name: 'db.example.com', root: '/usr/share/phpmyadmin', template: 'ssl_php', basic_auth: true, ssl_certificate_file: 'example.com.crt', ssl_certificate_key_file: 'example.com.key', ssl_dhparam_file: 'dhparam.pem' }
    { server_name: '2nd-cert.example.com', root: '/var/www/2nd-cert', template: 'ssl_php', ssl_certificate_file: '2nd-cert.example.com.crt', ssl_certificate_key_file: '2nd-cert.example.com.key', ssl_dhparam_file: 'dhparam.pem' },
    { server_name: 'yii-api.example.com', root: '/var/www/yii_project/api/web', template: 'ssl_yii', ssl_certificate_file: 'example.com.crt', ssl_certificate_key_file: 'example.com.key', ssl_dhparam_file: 'dhparam.pem' },
]
```

db.example.com runs PhpMyAdmin with HTTPS an HTTP Basic authentication

yii-api.example.com runs Yii api with HTTPS

2nd-cert.example.com runs php app with over HTTPS with another one ssl certificate


**ssl_certificate_file** - certificate from certificate authority (CA)

**ssl_certificate_key_file** - private key

**ssl_dhparam_file** - DHE parameter. We need to generate a DHE parameter with command:

```
openssl dhparam -out dhparam.pem 4096
```

SSL Server Test https://www.ssllabs.com/ssltest/analyze.html


Dependencies
------------

None.

License
-------

MIT

Author Information
------------------

[MobiDev](http://mobidev.biz/)
