#remove unused object (free cash)
docker system prune -a
#build in inamge 
docker image build -f Dockerfile -t ascii-art-web-image .
#run a container with the previous image
docker container run -p 8080:3000 --detach --name dockeriez ascii-art-web-image
#execute the bash in the running container
docker exec -it dockeriez /bin/bash