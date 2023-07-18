# combustion-sorc

### The purpose of this project was to become more familiar with NGINX, and practice creating a microservice backend. 
### When reading through the NGINX docs I found this photo <img width="1024" alt="deploy-NGINX-API-gateway-pt1_topology" src="https://user-images.githubusercontent.com/105041614/231921611-cfca4b48-5bb9-4f40-b415-1776166a6576.png">

### I decided to use this image as inspiration for a project to learn more about configuring NGINX as an API Gateway. So, the goal is to bring this image to life.

### My current plan for the backend services are as follows:
    - Have the inventory service perform CRUD operations. Connect this service to a Postgres db that is running as a docker container.
    - The pricing service will contain the prices for every item within the Postgres table of the inventory service. 
    - I have decided not to create an "other" service, and leave it as just the two services. 
    

### Post-mortem:
#### I took two different approaches for the services. 
#### For the inventory service, I kept things simple. I did not abrstact any files into seperate directories, and overall kept the whole code "lean and mean". I also coupled the database logic into the service itself, allowing the db.go file to create the DB connection and create the tables.
#### For the pricing service, I did some deeper research into project/directory structure. As a result, you will see that this service is set up very different from the inventory serivce. Additionally, I incorporated more features into this service such a logging and graceful server shutdown. For a service that is only performing one task, this might be a bit overkill but I like the extra information (makes the service feel more sexy). Lastly, instead of incorporating the DB logic into the application itself, I used the golang "migrate" tool.
#### Overall, the pricing service looks "cleaner" with the seperate directories but I feel like the inventory service is more intuitive. The program ran into less dependancy and import issues becuase everything was housed under the same directory. Additionally, I like to keep my main function straightforward, and I felt like it became rather bloated in the pricing service. I do like the migrate tool a lot, and I will most likely continue to use it for any DB table creations. All in all, I think there were a lot of techniques I learned which were incorporated into the pricing service and not the inventory service. I would like to keep the style the same moving forward, but incorporate more of these techniques into that style (such as the structured logging, graceful shutdown, error handling in the readJSON function, etc. etc.).  
