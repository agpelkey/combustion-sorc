# combustion-sorc

### The purpose of this project was to become more familiar with NGINX, and practice creating a microservice backend. 
### When reading through the NGINX docs I found this photo <img width="1024" alt="deploy-NGINX-API-gateway-pt1_topology" src="https://user-images.githubusercontent.com/105041614/231921611-cfca4b48-5bb9-4f40-b415-1776166a6576.png">

### I decided to use this image as inspiration for a project to learn more about configuring NGINX as an API Gateway. So, the goal is to bring this image to life.

### My current plan for the backend services are as follows:
    - Have the inventory service perform CRUD operations. Connect this service to a Postgres db that is running as a docker container.
    - The pricing service will contain the prices for every item within the Postgres table of the inventory service. 
    - I have decided not to create an "other" service, and leave it as just the two services. 
    
    
