import json
import os
import datetime
import time
from utils import cmd, initialize_project, date_f

auction_template_file = './auctions/auction_template.json'
auction_file = './auctions/auction.json'


def set_auction_json(selling_denom, selling_amount, paying_denom, start_price, min_bid_price, start_time, end_time):
    f = open(auction_template_file)
    jf = json.load(f)
    jf['selling_coin']['denom'] = selling_denom
    jf['selling_coin']['amount'] = selling_amount
    jf['paying_coin_denom'] = paying_denom
    jf['start_price'] = start_price
    jf['min_bid_price'] = min_bid_price
    jf['start_time'] = start_time
    jf['end_time'] = end_time
    with open(auction_file, 'w', encoding='utf-8') as newF:
        json.dump(jf, newF, ensure_ascii=False, indent=4)


if __name__ == "__main__":
    initialize_project()

    # Define auction start and end from current time
    date_now = datetime.datetime.utcnow()
    start = date_now + datetime.timedelta(seconds=15)
    end = date_now + datetime.timedelta(seconds=40)

    # Fundraising
    set_auction_json('v/1/orbit', '50000', 'uspn', '100', '50', date_f(start), date_f(end))
    cmd('networkd tx fundraising create-batch-auction {} --from alice --chain-id spn-1 --keyring-backend test -y'.format(
        auction_file))
    time.sleep(2)
    os.remove(auction_file)
    cmd('networkd tx participation participate 0 4 --from bob --chain-id spn-1 --keyring-backend test -y')
    time.sleep(2)
    cmd('networkd tx participation participate 0 4 --from carol --chain-id spn-1 --keyring-backend test -y')
    time.sleep(2)
    cmd('networkd tx participation participate 0 4 --from dave --chain-id spn-1 --keyring-backend test -y')

    # Wait auction start
    print("waiting for auction start...")
    time.sleep(15)

    # Place bid
    cmd('networkd tx fundraising place-bid 0 batch-many 120 10000v/1/orbit --from bob --chain-id spn-1 --keyring-backend test -y')
    time.sleep(2)
    cmd('networkd tx fundraising place-bid 0 batch-many 80 20000v/1/orbit --from carol --chain-id spn-1 --keyring-backend test -y')
    time.sleep(2)
    cmd('networkd tx fundraising place-bid 0 batch-many 140 20000v/1/orbit --from dave --chain-id spn-1 --keyring-backend test -y')

    # Wait withdrawal delay
    print("waiting for withdrawal delay...")
    time.sleep(5)

    cmd('networkd tx participation withdraw-allocations 0 --from bob --chain-id spn-1 --keyring-backend test -y')
    time.sleep(2)
    cmd('networkd tx participation withdraw-allocations 0 --from carol --chain-id spn-1 --keyring-backend test -y')
    time.sleep(2)
    cmd('networkd tx participation withdraw-allocations 0 --from dave --chain-id spn-1 --keyring-backend test -y')
