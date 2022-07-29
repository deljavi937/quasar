#!/bin/sh

cp ~/quasar-demo/quasar/demos/orion-manual-demo/hermes_config.toml ~/.hermes/config.toml

hermes keys restore --mnemonic "old cinnamon boy hurry pipe upset exhibit title copy squirrel grit eye love toy cotton connect inhale cost quarter mistake ahead endless bless license" quasar

hermes keys restore --mnemonic "ready hundred phrase theme bar breeze zone system bitter double flush deposit sugar swap burger outside primary nature attend caught wire ticket depth cycle" cosmos

hermes keys restore --mnemonic "rabbit garlic monitor wish pony magic budget someone room torch celery empower word assume digital rack electric weapon urban foot sketch jelly wet myself" osmosis

## Checking balance
quasarnoded q bank balances quasar143wwmxhsd8nkwu7j8gzpv9ca503g8j55h059ew --node tcp://localhost:26659
gaiad q bank balances cosmos1lrelhs37akgz2wht0y377uerxjm9fh33ke3ksc  --node tcp://localhost:26669
osmosisd q bank balances osmo194580p9pyxakf3y3nqqk9hc3w9a7x0yrnv7wcz --node tcp://localhost:26679


# Create connection
hermes create connection quasar cosmos

hermes create connection quasar osmosis

hermes create connection osmosis cosmos

# Create channel

hermes create channel --port-a transfer --port-b transfer cosmos connection-0

hermes create channel --port-a transfer --port-b transfer cosmos connection-1

hermes create channel --port-a transfer --port-b transfer quasar connection-1

# start
hermes start
