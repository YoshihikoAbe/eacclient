# Konasute Services

This document covers the Konasute specific services that eacclient implements.

## getResourceInfo

### Description:

Returns the URL to the latest resource list.

### Request:

```xml
<eacnet>
    <request>
        <method __type="str">getResourceInfo</method>
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
        <hash __type="str">546b1bc8ff0a2eb0af240791107c72c08ae88eada9cf4a20e872fadf54d73eaa</hash>
        <url __type="str">https://d1q1iflay6y5jj.cloudfront.net/ResourceInfo/2024091800_acf6dc22cf47/qbg_2024091800_pro.xml</url>
    </response>
</eacnet>
```

### CLI:

`eacclient konasute getResourceInfo -g "bomber-girl"`

## getUserIDs

### Description:

Returns the IDs associated with a token.

### Request:

```xml
<eacnet>
    <request>
        <method __type="str">getUserIDs</method>
        <data>
            <dummy __type="u8">0</dummy>
        </data>
    </request>
    <info>
        <game_id __type="str">bomber-girl</game_id>
        <soft_version __type="str"></soft_version>
        <token __type="str">TOKEN</token>
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
        <card_num __type="str">0000000000000000</card_num>
        <ref_id __type="str">0000000000000000</ref_id>
        <data_id __type="str">0000000000000000</data_id>
        <sns_id __type="str">0</sns_id>
    </response>
</eacnet>
```

### CLI:

`eacclient konasute getUserIDs -g "bomber-girl" -t "TOKEN"`

## acRelay

### Description:

Forwards a request to an XRPC server.

### Request:

```xml
<eacnet>
    <request>
        <service __type="str">SERVICE</service>
        <module __type="str">MODULE</module>
        <method __type="str">METHOD</method>
        <data>
            ...
        </data>
    </request>
    <info>
        <game_id __type="str">bomber-girl</game_id>
        <soft_version __type="str"></soft_version>
        <token __type="str">TOKEN</token>
        <retry_count __type="s32">0</retry_count>
    </info>
</eacnet>
```

### Response:

```xml
<eacnet>
    <status __type="s32">0</status>
    <error_code __type="s32">0</error_code>
    <xrpc_status_code __type="s32">0</xrpc_status_code>
    <xrpc_fault_code __type="s32">0</xrpc_fault_code>
    <response>
        ...
    </response>
</eacnet>
```

### CLI:

`eacclient konasute acRelay -g "bomber-girl" -t "TOKEN" -v "QBG:J:A:A" "SERVICE" "MODULE" "METHOD" "FILENAME"`
