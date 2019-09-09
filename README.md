<h1>IoT driven Smart Inventory management using SCION protocol</h1>
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

`$ cd $SC`

The scion server should be automatically started , else it can be started manually with the following command.

`$ ./scion.sh start`

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
ensure accurate weight readings after which real time weight data is acquired. Run the following command to calibrate the load cell after placing 2 objects of known weights.

                                            `go run calibrate.go`
                                            
When asked, put the first weight on the scale. Then when asked, put the second weight on the scale. It will print out the AdjustZero and AdjustScale values. These values are updated in the main code **weight_server_full_rv2.go**



 ### CONFIGURING SCION CLIENT AND SERVER
<div>

  <p>
  <ul>
    <li><u>CONFIGURING SCION SERVER</u> <p>
      •	Act as a TCP Socket client and fetch UID values from Python TCP socket server whenever the card is swiped with RFID sensor. 
•	Fetch appropriate weight values from load cell [4]-HX711 [5] sensor.
•	The product data like name, expiry date, capacity and so on are also hard coded for each UID in the SCION [1]  server go script.
•	Amalgamate product data associated with respective UID and their corresponding weight as one compact JSON object and send it to requesting client SCION [1] AS node.

This script upon execution produces a JSON object with product UID, product data and the respective real time weight readings at the SCION [1] client AS node. The script is executed by running the below command as shown.

•	“go run weight_server_full.go -s 19-ffaa:1:161,[192.168.137.185]:30102” 

Where “30102” indicates the port number, “[192.168.137.185]” is the detected dynamic IP of the Raspberry Pi [2] and “19-ffaa:1:161” is the SCION [1] AS node installed on the Pi. The complete script can be found in the appendix section.

 </p></li>
  
  <li>
  <u>CONFIGURING SCION CLIENT</u>
  <p>This chapter discusses the slight modification in the SCION [1] client go script to receive the JSON object encapsulating product data, weight values and UID from the hosted SCION [1] server AS node. Apart from that, the script also performs the hosting of a Hyper Text Transfer Protocol (HTTP) socket which is utilized for exposing the JSON object over a HTTP connection so that the visualization tool Node-Red [6] can use a http get request to access the JSON object and then visualize the same. The script is executed by running the below command as shown.

•	“go run Scion_client.go -c 19-ffaa:1:bfa,[192.168.1.130]:30102 -s 19-ffaa:1:161,[192.168.137.185]:30102 ” 

Where “30102” indicates the port number,”19-ffaa:1:bfa,” indicates the SCION [1] client AS node, “[192.168.1.130]” indicates the client IP address(Personal Computer), “19-ffaa:1:161” indicates the SCION [1] server AS node and “[192.168.137.185]” indicates the server IP address(Raspberry Pi [2]). The complete script can be found in the appendix section.
</p>
  </li>
  
  </ul>
  </p>
</div>


<div align="center>
  <h><b>IMPLEMENTATION OF NODE_RED</b> </h>
  <br>
  <Image src="Images/noderednode.jpg" alt="Nodes" height="400px" width="600px">
</div>

<div><h align="center"><b>HARDWARE IMPLEMENTATION AND DATA VISUALIZATION USING NODE-RED</b></h>
   <br>
  <br>
    <Image src="Images/hardwaresetup.jpg" alt="Nodes" height="300px" width="400px">
   <Image src="Images/visualize.jpg" alt="Nodes" height="300px" width="400px">
    
  </div>



















