docker build -t $1/belajar-docker:1.0.0 .
docker run -itd --name myapp -p 8000:8000 $1/belajar-docker:1.0.0