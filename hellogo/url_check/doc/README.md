```
pwd
	~/url_check
touch Dockerfile
docker build
docker build -t mathapp-development .
docker run -it --rm -p 3030:3030 -v $PWD/jnuho.github.io/hello/url_check:/go/src/url_check url_check
```
