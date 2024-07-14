package io.zeebe.redis.exporter;

import io.zeebe.exporter.proto.RecordTransformer;
import io.zeebe.exporter.proto.Schema;
import io.camunda.zeebe.protocol.record.Record;


public class SQSSender {
    Config config = new Config();
    private String redisHost = config.redisHost;
    private int redisPort = config.redisPort;
    private String passwoString = config.redisPassword;
    Redis redis = new Redis(redisHost, redisPort, passwoString);

    boolean sendForm(Record record){
        String streamName = config.streamName;
        String recordValueType = record.getValueType().name();
        if (!config.IsOpen(recordValueType)) {
            return false;
        }
        byte[] protoRecord = recordToProtobuf(record);
        System.out.println(protoRecord);
        redis.sendToStream(streamName, protoRecord);
        return true;
    }

    private byte[] recordToProtobuf(Record record) {
        final Schema.Record dto = RecordTransformer.toGenericRecord(record);
        return dto.toByteArray();
    }

}


