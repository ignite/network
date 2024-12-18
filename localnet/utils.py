import subprocess

val_address = 'spnvaloper15rz2rwnlgr7nf6eauz52usezffwrxc0muf4z5n'


def cmd(command):
    subprocess.run([command], shell=True, check=True)


def cmd_devnull(command):
    subprocess.run([command], shell=True, check=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)


def date_f(d):
    return d.isoformat("T") + "Z"


def initialize_project():
    cmd_devnull('networkd tx staking delegate {} 100000uspn --from bob -y'.format(val_address))
    cmd_devnull('networkd tx staking delegate {} 100000uspn --from carol -y'.format(val_address))
    cmd_devnull('networkd tx staking delegate {} 100000uspn --from dave -y'.format(val_address))
    cmd_devnull('networkd tx profile create-coordinator --from alice -y')
    cmd_devnull('networkd tx project create-project orbit 1000000orbit --from alice -y')
    cmd_devnull('networkd tx project mint-vouchers 1 100000orbit --from alice -y')
