# go_weather

raspberry pi project for smart weather monitoring


HARDWARE:
-raspberry pi 3, 1gb ram
-temperature and humidity sensor (DHT11)


SOFTWARE:
sensor.py reads data from sensors and saves data to json file every 24hours (daily). 
every json file contains data about weather for each day
rest of the code is a backend and frontend for the web monitoring app (can be started using a command go run main.go)
its basically the web app with a chart that shows weather data throughout the days 
it was made for a school project where I needed to create an IoT device

I decided to use golang templating system for frontend (server side rendering)
to enhance the simplicity of the project. And also because of hardware limitations (1gb ram).

Also, I do know how to connect web server to proper sql database but in this case I was limited by hardware and
using docker to compose up some postgres database was really painfull. That is why I decided to use json files as my database.
I am 'querying' them using dates they were created 

Also nr 2, there is an led.go file that does not do anything. But if an led would be connected to the 13 pin, it would work. 
I was testing library, beacause I wanted to add a handler that reads the device data to show real-time weather, but I prefered 
to leave sensor management to python for easier development. 

