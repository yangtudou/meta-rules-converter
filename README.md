# meta-rules-converter

## Convert geo to meta-rule-set

```shell
./converter geosite -f ../geosite.dat -o ../meta-rule/geo/geosite
./converter geoip -f ../geoip.dat -o ../meta-rule/geo/geoip
./converter asn -f ../GeoLite2-ASN.mmdb -o ../meta-rule/asn
```

## Convert geo to sing-rule-set

```shell
./converter geosite -f ../geosite.dat -o ../sing-rule/geo/geosite -t sing-box
./converter geoip -f ../geoip.dat -o ../sing-rule/geo/geoip -t sing-box
./converter asn -f ../GeoLite2-ASN.mmdb -o ../sing-rule/asn -t sing-box
```

## Convert geo to Surge RULE-SET

```shell
./converter geosite -f ../geosite.dat -o ../surge-rule/geo/geosite -t surge
./converter geoip -f ../geoip.dat -o ../surge-rule/geo/geoip -t surge
./converter asn -f ../GeoLite2-ASN.mmdb -o ../surge-rule/asn -t surge
```

The converter writes policy-less external rule sets as `.list` files. Use them in
the Surge `[Rule]` section, for example:

```ini
RULE-SET,https://example.com/geosite/google.list,Proxy
RULE-SET,https://example.com/geoip/cn.list,DIRECT
```

Geosite attributes are also exported (for example, `google@cn.list`). V2Fly
regular-expression domain rules are skipped with a conversion warning because
Surge does not provide a lossless equivalent of `DOMAIN-REGEX`. IP rules are
written as `IP-CIDR` or `IP-CIDR6` and include `no-resolve`.
