<h1>IMPLEMENTATION OF SCION IN IOT APPLICATIONS</h1>
<p1>There are several Internet of things (IoT) applications running on legacy networks which are not flexible in terms of path selections, network outages and security concerns. Thus, the current situation demands switching over to a new network platform which is highly scalable, secure, and isolated in terms of its architecture. Thus, an IoT application is being attempted which will use Scalable Control and Isolation on Next generation networks (SCION) architecture and hence, surpasses the existing concerns on the current network protocols. The chosen Internet of things (IoT) application is based on Inventory Management wherein immediate actions like stock replenishment or availing discounts can be provided based on the monitored weights of the articles available for sale. The implementation follows a simple client server architecture in SCION wherein one autonomous system (AS) acts as the client and the other autonomous system (AS) acts as the server. The key process includes (I) the client AS generate a request to the server  AS to get the current weight data and radio frequency identification (RFID) data (ii) the server AS fetches the respective weight data and the RFID data over a micro-controller and responds back to client AS with the requested data (iii) the data obtained at the client side is visualized using the visualization platform. </p1>

<div align="center">
<h2>BLOCK DIAGRAM OF THE IOT SETUP </h2>
<Image src="Images/Blockdgm.png" class="center" style="width:50%">
</div>

<div><h3>HARDWARE DEVICE SPECIFICATIONS</h3>
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


<div>
  <h5><font size="6">CONFIGURING SCION CLIENT AND SERVER</font></h5>
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
  <div>







</div>
