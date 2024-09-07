package io.zeebe.redis.exporter;

import io.camunda.zeebe.exporter.api.context.Context;
import io.camunda.zeebe.protocol.record.RecordType;
import io.camunda.zeebe.protocol.record.ValueType;

import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

public final class RecordFilter implements Context.RecordFilter {

  private final List<RecordType> enabledRecordTypes;
  private final List<ValueType> enabledValueTypes;

  public RecordFilter() {
    // EVENT((short)0),
    // COMMAND((short)1),
    // COMMAND_REJECTION((short)2),
    // SBE_UNKNOWN((short)255),
    // NULL_VAL((short)255);
    final List<String> enabledRecordTypeList = parseAsList("COMMAND,EVENT,REJECTION,RECORD");

    enabledRecordTypes =
        Arrays.stream(RecordType.values())
            .filter(
                recordType ->
                    enabledRecordTypeList.isEmpty()
                        || enabledRecordTypeList.contains(recordType.name()))
            .collect(Collectors.toList());
            // JOB((short)0),
            // DEPLOYMENT((short)4),
            // PROCESS_INSTANCE((short)5),
            // INCIDENT((short)6),
            // MESSAGE((short)10),
            // MESSAGE_SUBSCRIPTION((short)11),
            // PROCESS_MESSAGE_SUBSCRIPTION((short)12),
            // JOB_BATCH((short)14),
            // TIMER((short)15),
            // MESSAGE_START_EVENT_SUBSCRIPTION((short)16),
            // VARIABLE((short)17),
            // VARIABLE_DOCUMENT((short)18),
            // PROCESS_INSTANCE_CREATION((short)19),
            // ERROR((short)20),
            // PROCESS_INSTANCE_RESULT((short)21),
            // PROCESS((short)22),
            // DEPLOYMENT_DISTRIBUTION((short)23),
            // PROCESS_EVENT((short)24),
            // DECISION((short)25),
            // DECISION_REQUIREMENTS((short)26),
            // DECISION_EVALUATION((short)27),
            // PROCESS_INSTANCE_MODIFICATION((short)28),
            // ESCALATION((short)29),
            // SIGNAL_SUBSCRIPTION((short)30),
            // SIGNAL((short)31),
            // RESOURCE_DELETION((short)32),
            // COMMAND_DISTRIBUTION((short)33),
            // PROCESS_INSTANCE_BATCH((short)34),
            // MESSAGE_BATCH((short)35),
            // FORM((short)36),
            // USER_TASK((short)37),
            // PROCESS_INSTANCE_MIGRATION((short)38),
            // COMPENSATION_SUBSCRIPTION((short)39),
            // CHECKPOINT((short)254),
            // SBE_UNKNOWN((short)255),
            // NULL_VAL((short)255);
    final List<String> enabledValueTypeList = parseAsList("JOB,DEPLOYMENT,VARIABLE");

    enabledValueTypes =
        Arrays.stream(ValueType.values())
            .filter(
                valueType ->
                    enabledValueTypeList.isEmpty()
                        || enabledValueTypeList.contains(valueType.name()))
            .collect(Collectors.toList());
  }

  private List<String> parseAsList(String list) {
    return Arrays.stream(list.split(","))
        .map(String::trim)
        .filter(item -> !item.isEmpty())
        .map(String::toUpperCase)
        .collect(Collectors.toList());
  }

  @Override
  public boolean acceptType(RecordType recordType) {
    return enabledRecordTypes.contains(recordType);
  }

  @Override
  public boolean acceptValue(ValueType valueType) {
    return enabledValueTypes.contains(valueType);
  }
}
