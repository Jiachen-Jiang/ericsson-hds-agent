# Ericsson HDS Agent Collector Tools

Ericsson HDS Agent uses the following tools to collect inventory data from the host machine:

 - `bmc-info` or `ipmitool` or `freeipmi-tools`
 - `smartctl` (download from https://www.smartmontools.org/)
 - `rpm` or `dpkg`
 - `lspci`
 - `lsusb`
 - `ethtool`
 - `ip`
 - `dmidecode`
 - `mcelog`

If a tool is not installed on the host machine, Agent will skip the collection of data from that tool and move on. It is recommended that the host machine install the above list of tools to collect the most amount of data.


## User Scripts

When first started, Agent looks for scripts in two folders, `Inventory/` and `Metrics/`, in the same location as its `node.id` file (by default the directory containing Agent, though this can be set with the `-chdir` flag). Any executable scripts found in `Inventory/` will be run with the inventory collectors, and those found in `Metrics/` will be run with the metric collectors. 

**Note: scripts must be executable or they will not be discovered.**

The scripts must output to stdout in order for Agent to correctly send their results to DataHub. Agent handles transforming the script output into a form ingestable by DataHub.

### Inventory

Inventory script output is stored in a blob with type `inventory.user`. Inside `Content`, the key is the name of the script, minus the final extension, and the value is whatever the script outputs to stdout. Errors are logged in the blob as they are in the `inventory.all` blob (in a three-tiered format). This can make it difficult to automatically parse the blob unless you use a dynamic language or make sure that user scripts also output three-tiered JSON strings.

In the examples directory, inventory collected with the script `folder_list.py` and `usbinfo.sh` will be stored in a blob. The keys in `Content` are `folder_list` and `usbinfo`. A sample output looks like:

```
{"Type":"inventory.user","Id":2,"Digest":"f1227b5a85eb513af772db46c50337f1751ad0d2","NodeID":"a81bca75c97176b594079866e3eea8e7","Timestamp":"1490379453","Content":{"folder_list":"total 11812\ndrwxrwxr-x 2 nodeprime nodeprime     4096 Mar  8 13:19 credentials\n-rw-r--r-- 1 root      root             0 Nov 11  2015 dump\ndrwxrwxr-x 8 nodeprime nodeprime     4096 Mar 24 09:59 tester1\ndrwxrwxr-x 2 nodeprime nodeprime     4096 Mar 13 13:54 fremontdc\ndrwxrwxr-x 5 nodeprime nodeprime     4096 Feb 27 05:48 tester2\n-rwxr-xr-x 1 nodeprime nodeprime 11748296 Mar  8 13:19 hds-agent\ndrwxr-xr-x 3 root      root          4096 Nov 11  2015 setupbios\n-rw-r--r-- 1 nodeprime nodeprime   319839 Nov 11  2015 setupbios-2013-10-03.tgz\n","usbinfo":"Bus 001 Device 003: ID 046b:ff10 American Megatrends, Inc. Virtual Keyboard and Mouse\nBus 001 Device 002: ID 046b:ff01 American Megatrends, Inc. \nBus 001 Device 001: ID 1d6b:0002 Linux Foundation 2.0 root hub\nBus 004 Device 001: ID 1d6b:0001 Linux Foundation 1.1 root hub\nBus 003 Device 001: ID 1d6b:0001 Linux Foundation 1.1 root hub\nBus 002 Device 001: ID 1d6b:0001 Linux Foundation 1.1 root hub\n"}}
```

### Metrics

Metric script output is stored as a metric set named `user.scriptName`, where `scriptName` is the filename of the script minus its final extension. The Agent's expects metric scripts to always output two lines, the first line being column names and the second line being column values. Agent handles sending the info to DataHub in the correct format. 
In the examples directory, metrics collected with the script `cpu.py` will be stored in the metric set `user.cpu`. Its output to stdout looks like:

```
user_delta nice_delta
0.5 0.0
```

Note that column names and data values are space separated.

Examples for custom inventory and metric collectors are in ericsson-hds-agent/examples directory.
