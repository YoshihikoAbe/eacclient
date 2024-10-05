# Common Services

This document covers the "common" services that eacclient implements.

## getServices

### Description:

Returns a list of service URLs. It functions similarly to e-amusement's `services.get`.

### Request:

```xml
<eacnet>
    <request>
        <method __type="str">getServices</method>
        <data>
            <dummy __type="u8">0</dummy>
        </data>
    </request>
    <info>
        <game_id __type="str">bomber-girl</game_id>
        <soft_version __type="str"></soft_version>
        <token __type="str"></token>
        <retry_count __type="s32">0</retry_count>
    </info>
</eacnet>
```

### Response:

```xml
<eacnet>
    <status __type="s32">0</status>
    <error_code __type="s32">0</error_code>
    <response>
        <service_num __type="s32">18</service_num>
        <service>
            <service_name __type="str">acRelay</service_name>
            <url __type="str">https://p.eagate.573.jp/game/konasteapp/APIEX/servletex/need_auth/AcRelay</url>
        </service>
        <service>
            <service_name __type="str">cancelReserveConsumeItem</service_name>
            <url __type="str">https://p.eagate.573.jp/game/konasteapp/APIEX/servletex/need_auth/CancelReserveConsumeItem</url>
        </service>
        ...
```

### CLI:

`eacclient common getServices -g "bomber-girl"`

## getServerClock

### Description:

Returns the server's current time.

### Request:

```xml
<eacnet>
    <request>
        <method __type="str">getServerClock</method>
        <data>
            <dummy __type="u8">0</dummy>
        </data>
    </request>
    <info>
        <game_id __type="str">bomber-girl</game_id>
        <soft_version __type="str"></soft_version>
        <token __type="str"></token>
        <retry_count __type="s32">0</retry_count>
    </info>
</eacnet>
```

### Response:

```xml
<eacnet>
    <status __type="s32">0</status>
    <error_code __type="s32">0</error_code>
    <response>
        <server_clock __type="u64">795700800</server_clock>
    </response>
</eacnet>
```

### CLI:

`eacclient common getServerClock -g "bomber-girl"`

## getServerState

### Description:

Returns the server's current state and the maintenance schedule.

### Request:

```xml
<eacnet>
    <request>
        <method __type="str">getServerState</method>
    </request>
    <info>
        <game_id __type="str">bomber-girl</game_id>
        <soft_version __type="str"></soft_version>
        <token __type="str"></token>
        <retry_count __type="s32">0</retry_count>
    </info>
</eacnet>
```

### Response:

```xml
<eacnet>
    <status __type="s32">0</status>
    <error_code __type="s32">0</error_code>
    <response>
        <server_state __type="s32">1</server_state>
        <mainte_start_clock __type="u64">0</mainte_start_clock>
        <mainte_end_clock __type="u64">0</mainte_end_clock>
    </response>
</eacnet>
```

### CLI:

`eacclient common getServerState -g "bomber-girl"`