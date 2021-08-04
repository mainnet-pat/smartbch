# Docker

To run smartBCH via `docker-compose` you can execute the commands below! Note, the first time you run docker-compose it will take a while, as it will need to build the docker image.

```
# Generate a set of 10 test keys.
docker-compose run smartbch gen-test-keys -n 10 > test-keys.txt

# Init the node, include the keys from the last step as a comma separated list.
docker-compose run smartbch init mynode --chain-id 0x2711 \
    --init-balance=10000000000000000000 \
    --test-keys=`paste -d, -s test-keys.txt` \
    --home=/root/.smartbchd --overwrite

# Generate consensus key info
CPK=$(docker-compose run -w /root/.smartbchd/ smartbch generate-consensus-key-info)
docker-compose run --entrypoint mv smartbch /root/.smartbchd/priv_validator_key.json /root/.smartbchd/config

# Generate genesis validator
K1=$(head -1 test-keys.txt)
VAL=$(docker-compose run smartbch generate-genesis-validator $K1 \
  --consensus-pubkey $CPK \
  --staking-coin 10000000000000000000000 \
  --voting-power 1 \
  --introduction "tester" \
  --home /root/.smartbchd
  )
docker-compose run smartbch add-genesis-validator --home=/root/.smartbchd $VAL

# Start it up, you are all set!
# Note that the above generated 10 accounts are not unlocked, you have to operate them through private keys
docker-compose up
```
