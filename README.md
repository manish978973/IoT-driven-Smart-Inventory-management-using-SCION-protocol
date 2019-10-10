## IoT driven Smart Inventory management using SCION protocol
<p1>There are several Internet of things (IoT) applications running on legacy networks which are not flexible in terms of path selections, network outages and security concerns. Thus, the current situation demands switching over to a new network platform which is highly scalable, secure, and isolated in terms of its architecture. Thus, an IoT application is being attempted which will use <a href="https://www.scion-architecture.net">Scalable Control and Isolation on Next generation networks </a> (SCION) architecture and hence, surpasses the existing concerns on the current network protocols. The chosen Internet of things (IoT) application is based on Inventory Management wherein immediate actions like stock replenishment or availing discounts can be provided based on the monitored weights of the articles available for sale. The implementation follows a simple client server architecture in SCION wherein one autonomous system (AS) acts as the client and the other autonomous system (AS) acts as the server. The key process includes (I) the client AS generate a request to the server  AS to get the current weight data and radio frequency identification (RFID) data (ii) the server AS fetches the respective weight data and the RFID data over a micro-controller and responds back to client AS with the requested data (iii) the data obtained at the client side is visualized using the visualization platform. The SCION network can be set up in the local machine based on the tutorials and referrences in the SCION portal. </p1>

<div align="center">
<h2>BLOCK DIAGRAM OF THE IOT SETUP </h2>
<Image src="Images/Blockdgm.png" class="center" style="width:50%">
</div>

<div><h3>HARDWARE AND SOFTWARE SPECIFICATIONS</h3>
<ul>
  <li>Raspberry 3B+ <p>The application uses Raspberry Pi [2] as the micro-controller unit and around two Raspberry Pi [2] has been used upon which two separate SCION [1] AS nodes are installed respectively, one acts as the server interfaced with the RFID module and load [4] cell and the other acts as the client interfaced with the LCD unit where the weight data is being displayed (not within the scope of this report).</p></li>  
  
  <li>
  Load cell and Amplifier HX711 <p>As the main objective of application is to monitor and fetch the real time weight, load cell [4] of weight range up to 1 kg has been used. It is essentially a transducer that detects an applied load (force weight) and then changes it into an electrical signal output that is proportional to the load. The load causes a deformation in the dimensions of the underlying strain gauge; this deformation changes the electrical resistance which further changes the output electrical voltage as the wheat stone bridge is not balanced. The Load cell [4] consists of four pin outlets: red, black, green and white. 
  <br>
  
  The voltage signals which comes as output from the load cell [4] is then amplified with a HX711 [5] , a 24 bit analog to digital converter. This amplifier has got four pin outlets namely E+ (excitation) or VCC, E- (excitation) or ground and two outputs, that is A- and A+.
  
  </p>
  </li>
  
  <li>
  MFRC522 RFID reader <p>The MFRC522 [9] is a highly integrated reader/writer for contact-less communication. It uses the concept of radio-frequency identification and electromagnetic fields to transfer data over short distances. RFID employed in our project is useful to identify product data for further processing and analyzing. An RFID system uses RFID tags are attached to the object which is to be identified. Key chain and an electromagnetic card has been used in our project which has their own unique identification ids. It has got the following pins which are being utilized namely SDA, SCK, MOSI, MISO, GND, RST.</p>
  </li>
</ul>
</div>

### INSTALLATION OF SCION IN VIRTUAL MACHINE AND RASPBERRY PI

* A new SCION [1] Lab AS is generated with the a unique id is generated.
* The Virtual machine(VM) configuration file of this generated AS node is then downloaded. 
* This SCION [1] AS configuration file is available for download at the site [www.scionlab.org](www.scionlab.org) where registration of a user account is mandatory. 
* Following this the downloaded file is extracted and the virtual environment for SCION is set using Virtual Box and vagrant.
* Following this the vagrant is started and we navigate to the path where SCION AS is installed.

           $ cd $SC

The scion server should be automatically started , else it can be started manually with the following command.

           $ ./scion.sh start

* The `$ tail -f $SC/logs/bs*.DEBUG` command can be executed to check if the SCION server is working fine and we are able to receive the beacons.
* A similar procedure can be repeated in Raspberry pi to get SCION installed on top of Ubuntu Mate on Raspberry Pi 3B+.

### INSTALLATION OF UBUNTU MATE IN RASPBERRY PI

* Download a copy of Raspbian OS for Pi 3B+ and Ubuntu MATE for raspberry Pi 3B.
* Copy and then replace the following files from Raspbian [10] OS to Ubuntu MATE: bootcode.bin,fixup.dat, start.elf, bcm270-rpi-3-b-plus.dtb,kernel17.img, all contents from /lib/modules/4.9.80-v7+ and /lib/firmware/brcm/
* Flash the new OS to the SD card via Etcher and perform the similar steps for Ubuntu MATE as described on [www.scionlab.org](www.scionlab.org).


`Instead of command prompt, Putty has been used to connect with the Raspberry Pi and then to run and host the SCION [1]
network from the Raspberry Pi [2].`

### INTERFACING LOAD CELL, HX711 AMPLIFIER AND MFRC522 READER WITH PI

The HX711 [5] weight amplifier is interfaced with the load cell [4] using it’s following connections

• Excitation (E+) or VCC is red

• Excitation (E-) or ground is black

• Output (A+) is white

• Output (A-) is green

This load cell HX711 [5] integrated sensor unit is then interfaced with the Raspberry Pi serially as follows:

• Vcc of HX711 to Pin 2 (5V)

• GND to Pin 6 (GND)

• DT to Pin 29 (GPIO 5)

• SCK to Pin 31 (GPIO 6)

This sensor unit comprising of load cell [4] and HX711 [5] amplifier needs to be tested and calibrated accordingly to
ensure accurate weight readings after which real time weight data is acquired. Run the following command to **calibrate** the load cell after placing 2 objects of known weights.

                                            go run calibrate.go
                                            
When asked, put the first weight on the scale. Then when asked, put the second weight on the scale. It will print out the AdjustZero and AdjustScale values. These values are updated in the main code **weight_server_full_rv2.go**

RFID tags are attached to the product or pallet which is to be identified. The application primarily uses a card and a key
chain as the RFID tag. Each tag is associated with a unique id.The RFID reader sends a signal to the tag and read it’s
response. There are 8 hardware connections for RFID sensor with the Raspberry Pi [2] as follows:

• SDA connects to Pin 24

• SCK connects to Pin 23

• MOSI connects to Pin 19

• MISO connects to Pin 21

• GND connects to Pin 6

• RST connects to Pin 22

• 3.3v connects to Pin 1 

By default the Pi has the SPI (Serial Peripheral Interface) disabled, which is a prerequisite for the RFID reader
to function and therefore, needs to be enabled using rasp iconfig tool as follows:

• Run the command `sudo raspi-config`

• Select “Interfacing options”

• Select “P4 SPI” and then, select “Yes”

• Run the command “sudo reboot” to reboot the Pi

• Run the command `lsmod | grep spi` to check.

• And ensure if spi_bcm2835 is listed.



 ### CONFIGURING SCION CLIENT AND SERVER
<div>

  <p>
  <ul>
    <li><u>CONFIGURING SCION SERVER</u> <p>
      •	Scion server act as a TCP Socket client and fetch UID values from Python TCP socket server whenever the card is swiped with RFID sensor. 
      
      Run the command python RFID_TH_SERVER.py to start the Python TCP socket server and fetch values from RFID to the server side.
      
•	Fetch appropriate weight values from load cell [4]-HX711 [5] sensor.

•	The product data like name, expiry date, capacity and so on are also hard coded for each UID from RFID in the SCION [1]  server go script.

•	Amalgamate product data associated with respective UID and their corresponding weight as one compact JSON object and send it to requesting client SCION [1] AS node.

This script upon execution produces a JSON object with product UID, product data and the respective real time weight readings at the SCION [1] client AS node. The script is executed by running the below command as shown.

•		`go run weight_server_full_rv2.go -s 19-ffaa:1:161,[192.168.137.185]:30102` 

Where “30102” indicates the port number, “[192.168.137.185]” is the detected dynamic IP of the Raspberry Pi [2] and “19-ffaa:1:161” is the SCION [1] AS node installed on the Pi. 

 </p></li>
  
  <li>
  <u>CONFIGURING SCION CLIENT</u>
  <p>This chapter discusses the slight modification in the SCION [1] client go script to receive the JSON object encapsulating product data, weight values and UID from the hosted SCION [1] server AS node. Apart from that, the script also performs the hosting of a Hyper Text Transfer Protocol (HTTP) socket which is utilized for exposing the JSON object over a HTTP connection so that the visualization tool Node-Red [6] can use a http get request to access the JSON object and then visualize the same. The script is executed by running the below command as shown.

•	`go run weight_client_retry.go -c 19-ffaa:1:bfa,[192.168.1.130]:30102 -s 19-ffaa:1:161,[192.168.137.185]:30102`

Where “30102” indicates the port number,”19-ffaa:1:bfa,” indicates the SCION [1] client AS node, “[192.168.1.130]” indicates the client IP address(Personal Computer), “19-ffaa:1:161” indicates the SCION [1] server AS node and “[192.168.137.185]” indicates the server IP address(Raspberry Pi [2]). 
</p>
  </li>
  
  </ul>
  </p>
</div>

### CLIENT SERVER COMMUNICATION FLOW

* The Raspberry Pi [2] initiates the RFID reader to read the UID of the scanned RFID tag if detected.

* The Raspberry Pi [2] hosts a TCP socket server bound to a specific port. The data packet containing the UID is converted to JSON 
format and then send as bytes to this TCP connection and starts listening for any incoming request. Whenever this TCP server receives a request, it responds with this data packet.

* Then, the Raspberry Pi [2] runs and connects to the SCION [1] network as a SCION [1] AS server node by
identifying the SCION [1] server address, scion path and dispatcher path. Once identified, the SCION [1] server starts listening over a UDP connection for any incoming requests. Then a TCP connection is established with that specifc port by the SCION [1] server and the HX711 [5] sensor is also initialized. Whenever any incoming request is encountered from the SCION [1] AS client node, it sends a sample data packet as request to the TCP socket over a TCP connection. It gets back the UID as the TCP response from the TCP server. 

* If the UID obtained is a valid existing one, then the real time weight reading from the HX711[5]-load cell [4] unit is fetched and then appended to the data packet.Moreover the UID, current number, current time are also appended to the data packet which is then converted into JSON format.

* This data packet is then sent to the respective client SCION[1] AS node over the SCION [1] network.

* The client SCION [1] AS node now receives the data packet as a JSON object encapsulating the UID, product
name, product expiry date, current time, unit weight, capacity, current number and weight.

* The client SCION [1] AS hosts a HTTP socket on a specific port and establish a HTTP connection and starts listening. Then, it exposes this data packet as JSON object over this defined port. Whenever, this HTTP server encounters a request from any browser or user, it sends with this JSON object as response.

* The visualization tool, Node-Red [6] dashboard makes a HTTP GET request to this defined port, and gets this JSON object as response which is then processed, parsed and displayed using several node functions available in the Node-Red [6].



### IMPLEMENTATION OF GUI : NODE RED

The flowbased, visual programming development tool, Node-Red [6]has been employed for this purpose.[Node-RED](https://nodered.org).
is a programming tool for wiring together hardware devices, APIs and online services in new and interesting ways.It provides a browser-based editor that makes it easy to wire together flows using the wide range of nodes in the palette that can be deployed to its runtime in a single-click.

Node-Red can be installed in Windows, Linux and Raspberry Pi. Please refer the [installation procedure](https://nodered.org/docs/getting-started/windows)for getting started with Node-Red.

Once installed follow up the steps: 

![alt text](https://github.com/manish978973/SCION_IOT/blob/master/Images/noderednode.jpg "Logo Title Text 1")

* A “http” input type node is used to access the JSON object (data packet) hosted on a HTTP server at port 4000.The node is configured accordingly to use a GET method at the Uniform Resource Locator (URL) (port 4000). The JSON object is received as a message payload at the NodeRed [6] end.

* Several change type nodes are used to parse several data information from the entire message payload (JSONobject). Change type nodes are used to process and parse the temperature, humidity, product name, expiry date, capacity, current number, weight information from the message payload. 

* Several dashboard nodes like “text”, “gauge”, “chart” are used to visualize all these parsed separate information as a gauge, wave, level default text types of visualization.

Node-Red acts as a Graphical User Interface for the users so that they can monitor the inventory status of products in real
time which eliminates the need to do frequent physical inventory inspection and take business decisions accordingly. This also facilitates the users to take business decisions either to replenish the decreasing stock or to grant discounts in real time. 

### RESULTS

* Rigid assembly with the appropriate placement of the load cell [4], RFID reader module and the Raspberry Pi [2] to form a stand-alone “smart pallet” unit. The assembly comprises of a mounting plate which acts as the platform for keeping the RFID tag attached product pallet and a placeholder unit which contains the RFID, load cell [4] connected Raspberry Pi[2].

<div><h align="center"><b>HARDWARE IMPLEMENTATION AND DATA VISUALIZATION USING NODE-RED</b></h>
   <br>
  <br>
  <div align="center>
    <Image src="Images/hardwaresetup.jpg" alt="Nodes" height="300px" width="400px">
   <Image src="Images/visualize.jpg" alt="Nodes" height="300px" width="400px">
                                                                             </div>
  </div>

### CONCLUSION AND FURTHER WORK

The application which we have realized is just an initial development and posses wide scope of developmental progress in several application contexts. This application can find it’s usage in wholesale, retail, manufacturing, warehouses and logistics domains wherever an inventory status and management of products are required on it’s further development like using load cell [4] of wide range. The application can extend it’s functionalities to automatic discount generation in real time, automatic order generation
for stock replenishment in real time, monitor the sales data of the product which can further used for reconciliation with the
point of sales report. 

Another possible extension is to incorporate other sensors like temperature, humidity and pressure sensors which would make the use case more realistic and even applicable for monitoring and transportation of cold storage and fresh foodproducts where the real time temperature, pressure and humidity of the environment can be known and hence the temperature of the cold storage area can be remotely
controlled and changed accordingly.

### ABBRREVIATIONS

| Form      | Definition           |
| ------------- |:-------------:| 
|SCION          |Scalability, Control , And Isolation On Next Generation Networks| 
|IoT            |Internet of Things|  
|RFID           |Radio Frequency Identification|   
|AS             |Autonomous System|      
|GUI            |Graphical User Interface| 
|SPI            |Serial Peripheral Interface| 
|TCP            |Transmission Control Protocol| 














