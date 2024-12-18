import argparse
import time
from start_testnet import start_testnet
from initialize_rewards import initialize_rewards
from utils import cmd

parser = argparse.ArgumentParser(description='Start a testnet and connect it for reward')
parser.add_argument('--spn_chain_id',
                    help='Chain ID on SPN',
                    default='spn-1')
parser.add_argument('--orbit_chain_id',
                    help='Chain ID on Orbit',
                    default='orbit-1')
parser.add_argument('--spn_unbonding_period',
                    type=int,
                    default=1814400,
                    help='Unbonding period on spn',
                    )
parser.add_argument('--spn_revision_height',
                    type=int,
                    default=10,
                    help='Revision height for SPN IBC client',
                    )
parser.add_argument('--last_block_height',
                    type=int,
                    default=30,
                    help='Last block height for monitoring packet forwarding',
                    )
parser.add_argument('--max_validator',
                    type=int,
                    default=10,
                    help='Staking max validator set',
                    )
parser.add_argument('--self_delegation_1',
                    default='10000000uspn',
                    help='Self delegation for validator 1',
                    )
parser.add_argument('--self_delegation_2',
                    default='10000000uspn',
                    help='Self delegation for validator 2',
                    )
parser.add_argument('--self_delegation_3',
                    default='10000000uspn',
                    help='Self delegation for validator 3',
                    )
parser.add_argument('--unbonding_time',
                    default=1814400,  # 21 days = 1814400 seconds
                    type=int,
                    help='Staking unbonding time (unbonding period)',
                    )

if __name__ == "__main__":
    # Parse params
    args = parser.parse_args()
    spnChainID = args.spn_chain_id
    chainID = args.orbit_chain_id
    spnUnbondingPeriod = args.spn_unbonding_period
    revisionHeight = args.spn_revision_height
    lastBlockHeight = args.last_block_height
    maxValidator = args.max_validator
    selfDelegationVal1 = args.self_delegation_1
    selfDelegationVal2 = args.self_delegation_2
    selfDelegationVal3 = args.self_delegation_3
    unbondingTime = args.unbonding_time

    # Initialize rewards
    print('intialize rewards')
    initialize_rewards(
        lastBlockHeight,
        selfDelegationVal1,
        selfDelegationVal2,
        selfDelegationVal3,
    )
    print('rewards initialized')

    cmd('networkd q ibc client self-consensus-state --height {} > spncs.yaml'.format(revisionHeight))

    # Start the testnet
    print('start network')
    start_testnet(
        spnChainID,
        chainID,
        spnUnbondingPeriod,
        revisionHeight,
        lastBlockHeight,
        maxValidator,
        selfDelegationVal1,
        selfDelegationVal2,
        selfDelegationVal3,
        unbondingTime,
        True,
    )
    print('network started')

    time.sleep(10)

    # Create verified IBC client on SPN
    print('create verified client')
    cmd('networkd q tendermint-validator-set 2 --node "tcp://localhost:26659" > vs.yaml')
    cmd('networkd q ibc client self-consensus-state --height 2 --node "tcp://localhost:26659" > cs.yaml')
    time.sleep(2)
    cmd('networkd tx monitoringc create-client 0 cs.yaml vs.yaml --unbonding-period {} --revision-height 2 --from alice --chain-id spn-1 --keyring-backend test -y'.format(
        unbondingTime))
    time.sleep(2)

    # Perform IBC connection
    cmd('hermes --config ./hermes/config.toml create connection --a-chain spn-1 --a-client 07-tendermint-0 --b-client 07-tendermint-0')
    cmd('hermes --config ./hermes/config.toml create channel --a-port monitoringc --b-port monitoringp --a-chain spn-1 --a-connection connection-0 --order ordered --channel-version monitoring-1')

    # hermes --config ./hermes/config.toml start
