## Sprint 1

### User stories

1.  As someone looking for a job, I want to be able to quickly and easily share my Linkedin, Resume, Personal Website, Cover Letter, Phone, Email, etc. with others in order to advertise myself as effectively as possible.
2.  As a social media personality, I want to be able to select from the following social platforms to display them on my page: Instagram, Snapchat, TikTok, Twitter, YouTube, Email, etc. to allow me to connect with as many people as possible across all social platforms.
3.  As a college student looking to make friends, I want to be able to share my social media accounts simply by having a sticker, so that it is easier and faster to connect with people.
4. As a businessperson, I want to be able to upload my own profile picture and bio to allow a more personal connection when networking with others.
5. As someone concerned with their online security, I want to be able to login to the platform from a choice of authentication methods to allow me to choose the option that works best with my pre-existing authentication patterns.

### Issues Planned to Address

1. Storing the user information they want to display.
2. Create a SQL database with all the users' information.
3. Define handler functions for http requests.
4. Create a mock user profile in Angular

#### Completed
On the backend, we were able to create a MySQL database and create a server connection so that the data from the requests gets written in our tables. We tested the functionality using Postman. With that said, the handler functions turn out to work well also. Now we have to extend the functionality of our database and have more data stored for each user.

On the frontend, we were able to create a mock user profile with a name, bio, picture, and clickable social media icons. After that, we began work on routing, creating a home, registration, and error pages. 

#### Not Completed
On the backend, we weren't able to store every information the user wants to display, we still have to figure out a way to store the social media accounts and link them to their actual social media page.

On the frontend, we weren't able to connect with the backend, so that is our primary goal for the next sprint. Additionally, we didn't finish the home and registration pages, so we need to fill those out. 
