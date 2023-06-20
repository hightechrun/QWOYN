#!/bin/sh

rm -rf ~/.qwoynd

echo "silly slab oxygen reflect hawk wasp peace omit carbon pause turkey organ relax sing youth since fence increase record thing trial alien render begin" > validator.txt
echo "weather leader certain hard busy blouse click patient balcony return elephant hire mule gather danger curious visual boy estate army marine cinnamon snake flight" > mnemonic.txt;
echo "never chuckle bird almost jacket veteran weekend original rare habit point scorpion place gadget net train more plug upon pear renew mule material dynamic" > mnemonic2.txt;

# Build genesis
qwoynd50 init --chain-id=qwoyn-1 test
qwoynd50 keys add validator --keyring-backend="test" < validator.txt;
qwoynd50 keys add maintainer --recover --keyring-backend=test < mnemonic.txt;
qwoynd50 keys add user1 --recover --keyring-backend=test < mnemonic2.txt;

VALIDATOR=$(qwoynd50 keys show validator -a --keyring-backend="test")
MAINTAINER=$(qwoynd50 keys show maintainer -a --keyring-backend="test")
USER1=$(qwoynd50 keys show user1 -a --keyring-backend="test")
# VALIDATOR=qwoyn1hzqg4r2e789930hs88wqle25ef94xajuqay93r
# MAINTAINER=qwoyn1h9krsew6kpg9huzcqgmgmns0n48jx9yd5vr0n5
# USER1=qwoyn13tqzdukugulllnk3p5js3w7hzw8gclkeenzp6e
qwoynd50 genesis add-genesis-account $VALIDATOR 1000000000000uqwoyn,1000000000000ucoho,1000000000000stake
qwoynd50 genesis add-genesis-account $MAINTAINER 1000000000000uqwoyn,1000000000000ucoho,1000000000000stake
qwoynd50 genesis add-genesis-account $USER1 1000000000000uqwoyn,1000000000000ucoho,1000000000000stake
qwoynd50 genesis gentx validator 100000000stake --keyring-backend="test" --chain-id=qwoyn-1
qwoynd50 genesis collect-gentxs
sed -i '' 's/"voting_period": "172800s"/"voting_period": "20s"/g' $HOME/.qwoynd/config/genesis.json

# sed -i 's/stake/uqwoyn/g' $HOME/.qwoynd/config/genesis.json

# Start node
qwoynd50 start --pruning=nothing --minimum-gas-prices="0stake"
