import time
import json
import adafruit_dht
from board import <pin>
from datetime import datetime

def save_data(data):
    
    timestamp = datetime.now().strftime("%Y-%m-%d")
    filename = f"data_{timestamp}.txt"

    with open(filename, "w") as file:
        json.dump(data)


while True:
    
    
    dht_device = adafruit_dht.DHT11(<pin>)
    temperature = dht_device.temperature
    humidity = dht_device.humidity

    data = {"Temperature": temperature, "Humidity": humidity}
  
    save_data(data)

   
    time.sleep(86400) # 24 hours