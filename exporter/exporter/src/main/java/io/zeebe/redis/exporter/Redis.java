package io.zeebe.redis.exporter;
import java.util.Map;

import redis.clients.jedis.Jedis;
import redis.clients.jedis.StreamEntryID;

import org.apache.commons.codec.binary.Base64;
import java.util.HashMap;
import java.util.Map;

public class Redis {
    private Jedis Jedis;

    public Redis(String host, int port, String password) {
        if (Jedis == null) {
            Jedis = new Jedis(host, port);
            Jedis.auth(password);
        }
    }

    public Boolean set(String key, String value) {
        try {
            Jedis.set(key, value);
            return true;
        } catch (Exception e) {
            System.err.println("Failed to save recordPosition to Redis.");
            return false;
        }
    }

    public String get(String key) {
        return Jedis.get(key);
    }

    public void close() {
        Jedis.close();
    }

        /**
     * Sends a message to a Redis stream.
     * @param streamName The name of the Redis stream.
     * @param messageData A map of the message fields to send.
     * @return The ID of the added stream entry if successful, null otherwise.
     */
    public StreamEntryID sendToStream(String streamName, byte[] messageData) {
        if (messageData.length > 0) {
            try {
                // Prepare the message
                // Encode the byte array to a Base64 string
                String encodedData = Base64.encodeBase64String(messageData);

                // Prepare the message
                Map<String, String> message = new HashMap<>();
                message.put("data", encodedData);
                return Jedis.xadd(streamName, StreamEntryID.NEW_ENTRY, message);
            } catch (Exception e) {
                System.err.println("Failed to send message to Redis stream: " + e.getMessage());
                return null;
            }
        } else {
            System.err.println("Message data is empty.");
            return null;
        }
    }
    
}
