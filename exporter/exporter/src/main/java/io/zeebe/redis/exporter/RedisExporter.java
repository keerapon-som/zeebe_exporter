package io.zeebe.redis.exporter;

import io.camunda.zeebe.exporter.api.Exporter;

import io.camunda.zeebe.protocol.record.Record;

import io.camunda.zeebe.exporter.api.context.Context;
import io.camunda.zeebe.exporter.api.context.Controller;


import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;
// import redis.clients.jedis.Jedis;
import io.zeebe.redis.exporter.Redis;



public class RedisExporter implements Exporter {
    Config config = new Config();
    
    private boolean redisConnected = false;
    
    private Controller controller;
    private SQSSender sqsSender = new SQSSender();
    private long recordPosition;
    private long oldRecordPosition;
    private long lastPosition;

    private long period = config.periodExportPosition;
    private String redisPositionKeyName = config.redisPositionKeyName;
    private String redisHost = config.redisHost;
    private int redisPort = config.redisPort;
    private String passwoString = config.redisPassword;
    Redis redis = new Redis(redisHost, redisPort, passwoString);

    ScheduledExecutorService scheduler = Executors.newScheduledThreadPool(1);

    public RedisExporter() {
       
    }

    @Override
    public void configure(Context context) {
    }

    public void startSavingRecordPosition() {
        if (redisConnected) {
            final Runnable saveRecordPosition = () -> {
                if (recordPosition == 0 || recordPosition == oldRecordPosition) {
                    return;
                }
                Boolean setResult = redis.set(redisPositionKeyName, String.valueOf(recordPosition));
                if (setResult) {
                    System.out.println("Saved recordPosition to Redis: " + recordPosition);
                    oldRecordPosition = recordPosition;
                } else {
                    // Handle the case where the set operation was not successful
                    System.err.println("Failed to save recordPosition to Redis.");
                    try {
                        redis.close(); // Close the current connection
                        redis = new Redis(redisHost, redisPort, passwoString); // Reconnect
                        setResult = redis.set(redisPositionKeyName, String.valueOf(recordPosition)); // Retry the set operation
                        if (setResult) {
                            System.out.println("Saved recordPosition to Redis after reconnect: " + recordPosition);
                            oldRecordPosition = recordPosition;
                        } else {
                            System.err.println("Failed to save recordPosition to Redis after reconnect.");
                        }
                    } catch (Exception e) {
                        System.err.println("Error reconnecting to Redis: " + e.getMessage());
                    }
                }
            };
            scheduler.scheduleAtFixedRate(saveRecordPosition, 0, period, TimeUnit.SECONDS);
        }
    }

    @Override
    public void open(Controller controller) {
        this.controller = controller;
        try {
            Redis redis = new Redis(config.redisHost, config.redisPort, config.redisPassword);
            
            String lastExportedRecordPosition = redis.get(redisPositionKeyName);

            if (lastExportedRecordPosition != null) {
                System.out.println("Fetched last exported record position from Redis: " + lastExportedRecordPosition);
                lastPosition = Long.parseLong(lastExportedRecordPosition);
                redisConnected = true;
            }

            if (lastExportedRecordPosition != null && !lastExportedRecordPosition.isEmpty()) {
                try {
                    lastPosition = Long.parseLong(lastExportedRecordPosition);
                    redisConnected = true;
                } catch (NumberFormatException e) {
                    System.err.println("Error parsing last exported record position from Redis: " + e.getMessage());
                    redisConnected = false;
                }
            }
        } catch (Exception e) {
            System.err.println("Error Init connecting to Redis: " + e.getMessage());
            redisConnected = false;
        }

        if (redisConnected) {
            System.out.println("Connected to Redis.");
            controller.updateLastExportedRecordPosition(lastPosition);
            startSavingRecordPosition();
            System.out.println("Start Scheduler to save recordPosition to Redis every "+period+ " seconds.");
        } else {
            System.err.println("Failed to connect to Redis.");
        }
    }
    
    @Override
    public void export(Record record) {
        if (redisConnected) {
            System.out.println(record.toJson());
            recordPosition = record.getPosition();
            sqsSender.sendForm(record); // ฟังก์ชั่นส่งข้อมูลไป SQS
        }
    }
}
