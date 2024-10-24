import json
import os
import datetime
import time
from utils import cmd, initialize_project, date_f

sale_template_file = './auctions/sale_template.json'
sale_file = './auctions/sale.json'


def set_sale_json(selling_denom, selling_amount, paying_denom, price, start_time, end_time):
    f = open(sale_template_file)
    jf = json.load(f)
    jf['selling_coin']['denom'] = selling_denom
    jf['selling_coin']['amount'] = selling_amount
    jf['paying_coin_denom'] = paying_denom
    jf['start_price'] = price
    jf['start_time'] = start_time
    jf['end_time'] = end_time
    with open(sale_file, 'w', encoding='utf-8') as newF:
        json.dump(jf, newF, ensure_ascii=False, indent=4)


if __name__ == "__main__":
    initialize_project()

    # Define auction start and end from current time
    date_now = datetime.datetime.utcnow()
    start = date_now + datetime.timedelta(seconds=15)
    end = date_now + datetime.timedelta(seconds=55)

    # Fundraising
    set_sale_json('v/0/orbit', '50000', 'uspn', '100', date_f(start), date_f(end))
    time.sleep(2)
    cmd('networkd tx fundraising create-fixed-price-auction {} --from alice --chain-id spn-1 --keyring-backend test -y'.format(
        sale_file))
    os.remove(sale_file)
    time.sleep(2)
    cmd('networkd tx participation participate 0 4 --from bob --chain-id spn-1 --keyring-backend test -y')
    time.sleep(2)
    cmd('networkd tx participation participate 0 4 --from carol --chain-id spn-1 --keyring-backend test -y')
    time.sleep(2)
    cmd('networkd tx participation participate 0 4 --from dave --chain-id spn-1 --keyring-backend test -y')

    # Wait auction start
    print("waiting for auction start...")
    time.sleep(15)

    # Place bid
    cmd('networkd tx fundraising bid 0 fixed-price 100 10000v/0/orbit --from bob --chain-id spn-1 --keyring-backend test -y')
    time.sleep(2)
    cmd('networkd tx fundraising bid 0 fixed-price 100 20000v/0/orbit --from carol --chain-id spn-1 --keyring-backend test -y')
    time.sleep(2)
    cmd('networkd tx fundraising bid 0 fixed-price 100 20000v/0/orbit --from dave --chain-id spn-1 --keyring-backend test -y')

    # Wait withdrawal delay
    print("waiting for withdrawal delay...")
    time.sleep(15)

    cmd('networkd tx participation withdraw-allocations 0 --from bob --chain-id spn-1 --keyring-backend test -y')
    time.sleep(2)
    cmd('networkd tx participation withdraw-allocations 0 --from carol --chain-id spn-1 --keyring-backend test -y')
    time.sleep(2)
    cmd('networkd tx participation withdraw-allocations 0 --from dave --chain-id spn-1 --keyring-backend test -y')
