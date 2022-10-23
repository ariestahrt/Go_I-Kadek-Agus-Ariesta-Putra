docker stop /myapp
docker rm /myapp
docker build -t ariestadocker/belajar-docker:1.0.0 .
# docker run -itd -v $PWD/.env:/.env --name myapp -p 3000:3000 $1/belajar-docker:1.0.0
docker run -v $PWD/.env:/.env -p 3000:3000 ariestadocker/belajar-docker:1.0.0