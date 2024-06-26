#include <WiFi.h>
#include <DHT.h>
#include <PubSubClient.h>

const char* ssid = "your wifi SSID";
const char* password = "your wifi password";

const char* mqtt_server = "your mqtt broker url";
const char* mqtt_port = "mqtt port";
const char* mqtt_topic = "mqtt topic";

#define PHOTORESISTOR_PIN 36
#define DHTPIN 12
#define DHTTYPE DHT11

DHT dht(DHTPIN, DHTTYPE);

WiFiClient espClient;
PubSubClient client(espClient);

void setup() {
  Serial.begin(115200);
  
  Serial.println("Connecting to ");
  Serial.println(ssid);
  WiFi.begin(ssid, password);
  while (WiFi.status()!= WL_CONNECTED) { 
    delay(500);
    Serial.print(".");
  }
  Serial.println("");
  Serial.println("WiFi connected");
  Serial.println("IP address: ");
  Serial.println(WiFi.localIP());

  client.setServer(mqtt_server, 1883);

  pinMode(PHOTORESISTOR_PIN, INPUT);

  delay(3000);
}

void loop() {
  if (!client.connected()) {
    reconnect();
  }
  client.loop();

  float humidity = dht.readHumidity();
  float temperature = dht.readTemperature();
  int luminosity = analogRead(PHOTORESISTOR_PIN);

  String payload = "{\"temperature\": ";
  payload += String(temperature);
  payload += ", \"humidity\": ";
  payload += String(humidity);
  payload += ", \"luminosity\": ";
  payload += String(luminosity);
  payload += "}";
  client.publish(mqtt_topic, payload.c_str());
  Serial.println(payload);

  delay(1000);
}

void reconnect() {
  while (!client.connected()) {
    Serial.print("Attempting MQTT connection...");
    if (client.connect("ESP32HeltecClient")) {
      Serial.println("connected");
    } else {
      Serial.print("failed, rc=");
      Serial.print(client.state());
      Serial.println(" try again in 5 seconds");
      delay(5000);
    }
  }
}
