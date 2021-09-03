*** Settings ***
Documentation     A test suite that tests different Device creation scenarios.
Library           RabbitMq
Library           RequestsLibrary
Library           String
Library           json
Resource          ../resources/fiware.robot

Suite Setup       suite setup
Suite Teardown    suite teardown


*** Test Cases ***
Connect To RMQ

    ${event}=         Create Dictionary  type=Device    id=someid  body={}
    ${json_str}=      Json.Dumps  ${event}

    Publish Message	exchange_name=iot-msg-exchange-topic	routing_key=notify	payload=${json_str}


Get Subscriptions

    ${resp}=          GET On Session  diwise  /subscriptions


Create Subscription And Trigger Notification

    ${sub}=           Create Subscription  WaterQualityObserved  http://quantumleap:8668/v2/notify
    ${resp}=          POST On Session  diwise  /subscriptions  json=${sub}

    ${wqo}=           Create WaterQualityObserved  temp=20.2
    ${json}=          Json.Dumps  ${wqo}
    ${event}=         Create Dictionary  type=WaterQualityObserved    id=someid  body=${json}

    ${json_str}=      Json.Dumps  ${event}

    Publish Message	exchange_name=iot-msg-exchange-topic	routing_key=notify	payload=${json_str}


*** Keywords ***
suite setup
    Create Rabbitmq Connection	localhost	15672	5672	user	bitnami  alias=rmq  vhost=/

    ${headers}=       Create Dictionary   Content-Type=application/json
    Create Session    diwise    http://127.0.0.1:9090  headers=${headers}


suite teardown
    Close All Rabbitmq Connections
