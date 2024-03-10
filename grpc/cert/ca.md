

``` shell
## 生成私钥文件
openssl genrsa -des3 -out ca.key 2048

## 创建证书请求，
openssl req -new -key ca.key -out ca.csr

## 生成 ca.crt 
openssl x509 -req -days 365 -in ca.csr -signkey ca.key -out ca.crt




```