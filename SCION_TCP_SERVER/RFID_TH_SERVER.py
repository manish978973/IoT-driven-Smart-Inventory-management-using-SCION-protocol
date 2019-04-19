import RPi.GPIO as GPIO
import MFRC522
import signal
from multiprocessing import Process, Value, Array
import Adafruit_DHT
import socket
import time
import json


continue_reading = True
sensor = Adafruit_DHT.DHT11
pin = 17
MIFAREReader = MFRC522.MFRC522()
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.bind(('0.0.0.0',10000))
sock.listen(1)

def readTempAndHumidity(arr):
    while continue_reading:
        arr[1], arr[0] = Adafruit_DHT.read_retry(sensor, pin)
        # if arr[0] is not None and arr[1] is not None:
        #     print('Temp={0:0.1f}*  Humidity={1:0.1f}%'.format(arr[0],arr[1]))
        # else:
        #     print('Failed to get reading. Try again!')
        # if continue_reading == False:
        #     return
        time.sleep(2)

# Capture SIGINT for cleanup when the script is aborted
def end_read(signal,frame):
    global continue_reading
   # print "Ctrl+C captured, ending read."
    continue_reading = False
    sock.close()
    #TempReadProcess.join()
    GPIO.cleanup()

# Hook the SIGINT
signal.signal(signal.SIGINT, end_read)

arr = Array('d',(0.0,0.0))

TempReadProcess = Process(target=readTempAndHumidity, args=(arr,))
TempReadProcess.start()


uid_str = "[]"
while True:
    c,a = sock.accept()

    while True:
        data = c.recv(128)
        temperature = int(arr[0])
        humidity = int(arr[1])
        
        if continue_reading:
            
            # Scan for cards    
            (status,TagType) = MIFAREReader.MFRC522_Request(MIFAREReader.PICC_REQIDL)

            # If a card is found
            
            (status,uid) = MIFAREReader.MFRC522_Anticoll()
            if status == MIFAREReader.MI_OK:
                uid_str = str(uid)
            
            (status,TagType) = MIFAREReader.MFRC522_Request(MIFAREReader.PICC_REQIDL)

               
            data = {"UID" : str(uid),
                    "Temperature" : temperature,
                    "Humidity" : humidity}
            print(str(data))
            data = json.dumps(data)
        c.send(bytes(data))
        if not data:
            c.close()
            break
        
      
