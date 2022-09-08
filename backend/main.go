package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Ram struct {
	Total      int `json:"total"`
	Free       int `json:"free"`
	Used       int `json:"used"`
	Percentage int `json:"percentage"`
}

type Process struct {
	Pid      int       `json:"pid"`
	Name     string    `json:"name"`
	User     int       `json:"user"`
	State    int       `json:"state"`
	Memory   float32   `json:"memory"`
	Children []Process `json:"children"`
}

type ArrayProcess []Process

var body = `
[
  {
    "pid": 1,
    "name": "systemd",
    "user": 0,
    "state": 1,
    "memory": 3041,
    "children": [
      {
        "pid": 170,
        "name": "systemd-journal",
        "user": 0,
        "state": 1,
        "memory": 3271
      },
      {
        "pid": 201,
        "name": "systemd-udevd",
        "user": 0,
        "state": 1,
        "memory": 1303
      },
      {
        "pid": 322,
        "name": "multipathd",
        "user": 0,
        "state": 1,
        "memory": 4500
      },
      {
        "pid": 438,
        "name": "systemd-network",
        "user": 100,
        "state": 1,
        "memory": 1762
      },
      {
        "pid": 442,
        "name": "systemd-resolve",
        "user": 101,
        "state": 1,
        "memory": 2795
      },
      {
        "pid": 477,
        "name": "accounts-daemon",
        "user": 0,
        "state": 1,
        "memory": 2082
      },
      { "pid": 490, "name": "chronyd", "user": 112, "state": 1, "memory": 624 },
      {
        "pid": 495,
        "name": "dbus-daemon",
        "user": 103,
        "state": 1,
        "memory": 1032
      },
      {
        "pid": 498,
        "name": "google_osconfig",
        "user": 0,
        "state": 1,
        "memory": 3855
      },
      {
        "pid": 503,
        "name": "networkd-dispat",
        "user": 0,
        "state": 1,
        "memory": 4448
      },
      {
        "pid": 505,
        "name": "rsyslogd",
        "user": 104,
        "state": 1,
        "memory": 1129
      },
      { "pid": 515, "name": "snapd", "user": 0, "state": 1, "memory": 9407 },
      { "pid": 521, "name": "atd", "user": 0, "state": 1, "memory": 555 },
      {
        "pid": 525,
        "name": "containerd",
        "user": 0,
        "state": 1,
        "memory": 8745
      },
      {
        "pid": 562,
        "name": "google_guest_ag",
        "user": 0,
        "state": 1,
        "memory": 3380
      },
      { "pid": 583, "name": "polkitd", "user": 0, "state": 1, "memory": 2228 },
      { "pid": 584, "name": "agetty", "user": 0, "state": 1, "memory": 461 },
      { "pid": 609, "name": "agetty", "user": 0, "state": 1, "memory": 428 },
      {
        "pid": 613,
        "name": "unattended-upgr",
        "user": 0,
        "state": 1,
        "memory": 4752
      },
      { "pid": 711, "name": "dockerd", "user": 0, "state": 1, "memory": 20828 },
      { "pid": 719, "name": "sshd", "user": 0, "state": 1, "memory": 1634 },
      {
        "pid": 873,
        "name": "systemd-logind",
        "user": 0,
        "state": 1,
        "memory": 1753
      },
      { "pid": 876, "name": "cron", "user": 0, "state": 1, "memory": 676 },
      {
        "pid": 5964,
        "name": "systemd",
        "user": 1001,
        "state": 1,
        "memory": 2172
      }
    ]
  },
  {
    "pid": 2,
    "name": "kthreadd",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": [
      { "pid": 3, "name": "rcu_gp", "user": 0, "state": 1026, "memory": 0 },
      { "pid": 4, "name": "rcu_par_gp", "user": 0, "state": 1026, "memory": 0 },
      { "pid": 5, "name": "netns", "user": 0, "state": 1026, "memory": 0 },
      {
        "pid": 7,
        "name": "kworker/0:0H",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      {
        "pid": 9,
        "name": "kworker/0:1H",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      {
        "pid": 10,
        "name": "mm_percpu_wq",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      {
        "pid": 11,
        "name": "rcu_tasks_rude_",
        "user": 0,
        "state": 1,
        "memory": 0
      },
      {
        "pid": 12,
        "name": "rcu_tasks_trace",
        "user": 0,
        "state": 1,
        "memory": 0
      },
      { "pid": 13, "name": "ksoftirqd/0", "user": 0, "state": 1, "memory": 0 },
      { "pid": 14, "name": "rcu_sched", "user": 0, "state": 1026, "memory": 0 },
      { "pid": 15, "name": "migration/0", "user": 0, "state": 1, "memory": 0 },
      {
        "pid": 16,
        "name": "idle_inject/0",
        "user": 0,
        "state": 1,
        "memory": 0
      },
      { "pid": 18, "name": "cpuhp/0", "user": 0, "state": 1, "memory": 0 },
      { "pid": 19, "name": "cpuhp/1", "user": 0, "state": 1, "memory": 0 },
      {
        "pid": 20,
        "name": "idle_inject/1",
        "user": 0,
        "state": 1,
        "memory": 0
      },
      { "pid": 21, "name": "migration/1", "user": 0, "state": 1, "memory": 0 },
      { "pid": 22, "name": "ksoftirqd/1", "user": 0, "state": 1, "memory": 0 },
      {
        "pid": 24,
        "name": "kworker/1:0H",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      { "pid": 25, "name": "kdevtmpfs", "user": 0, "state": 1, "memory": 0 },
      {
        "pid": 26,
        "name": "inet_frag_wq",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      { "pid": 27, "name": "kauditd", "user": 0, "state": 1, "memory": 0 },
      { "pid": 29, "name": "khungtaskd", "user": 0, "state": 1, "memory": 0 },
      { "pid": 30, "name": "oom_reaper", "user": 0, "state": 1, "memory": 0 },
      { "pid": 31, "name": "writeback", "user": 0, "state": 1026, "memory": 0 },
      { "pid": 32, "name": "kcompactd0", "user": 0, "state": 1, "memory": 0 },
      { "pid": 33, "name": "ksmd", "user": 0, "state": 1, "memory": 0 },
      { "pid": 34, "name": "khugepaged", "user": 0, "state": 1, "memory": 0 },
      {
        "pid": 80,
        "name": "kintegrityd",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      { "pid": 81, "name": "kblockd", "user": 0, "state": 1026, "memory": 0 },
      {
        "pid": 82,
        "name": "blkcg_punt_bio",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      {
        "pid": 83,
        "name": "tpm_dev_wq",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      { "pid": 84, "name": "ata_sff", "user": 0, "state": 1026, "memory": 0 },
      { "pid": 85, "name": "md", "user": 0, "state": 1026, "memory": 0 },
      {
        "pid": 86,
        "name": "edac-poller",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      {
        "pid": 87,
        "name": "devfreq_wq",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      { "pid": 88, "name": "watchdogd", "user": 0, "state": 1, "memory": 0 },
      {
        "pid": 90,
        "name": "kworker/1:1H",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      { "pid": 92, "name": "kswapd0", "user": 0, "state": 1, "memory": 0 },
      {
        "pid": 93,
        "name": "ecryptfs-kthrea",
        "user": 0,
        "state": 1,
        "memory": 0
      },
      { "pid": 95, "name": "kthrotld", "user": 0, "state": 1026, "memory": 0 },
      {
        "pid": 96,
        "name": "acpi_thermal_pm",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      { "pid": 97, "name": "scsi_eh_0", "user": 0, "state": 1, "memory": 0 },
      {
        "pid": 98,
        "name": "scsi_tmf_0",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      { "pid": 99, "name": "nvme-wq", "user": 0, "state": 1026, "memory": 0 },
      {
        "pid": 100,
        "name": "nvme-reset-wq",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      {
        "pid": 101,
        "name": "nvme-delete-wq",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      {
        "pid": 102,
        "name": "vfio-irqfd-clea",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      { "pid": 104, "name": "mld", "user": 0, "state": 1026, "memory": 0 },
      {
        "pid": 105,
        "name": "ipv6_addrconf",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      { "pid": 115, "name": "kstrp", "user": 0, "state": 1026, "memory": 0 },
      {
        "pid": 118,
        "name": "zswap-shrink",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      {
        "pid": 119,
        "name": "kworker/u5:0",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      {
        "pid": 126,
        "name": "charger_manager",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      { "pid": 127, "name": "jbd2/sda1-8", "user": 0, "state": 1, "memory": 0 },
      {
        "pid": 128,
        "name": "ext4-rsv-conver",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      { "pid": 131, "name": "hwrng", "user": 0, "state": 1, "memory": 0 },
      { "pid": 228, "name": "cryptd", "user": 0, "state": 1026, "memory": 0 },
      { "pid": 318, "name": "kaluad", "user": 0, "state": 1026, "memory": 0 },
      {
        "pid": 319,
        "name": "kmpath_rdacd",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      { "pid": 320, "name": "kmpathd", "user": 0, "state": 1026, "memory": 0 },
      {
        "pid": 321,
        "name": "kmpath_handlerd",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      {
        "pid": 733,
        "name": "bpfilter_umh",
        "user": 0,
        "state": 1,
        "memory": 143
      },
      {
        "pid": 5928,
        "name": "kworker/1:0",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      {
        "pid": 5929,
        "name": "kworker/u4:1",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      {
        "pid": 5966,
        "name": "kworker/1:2",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      {
        "pid": 6370,
        "name": "kworker/0:1",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      {
        "pid": 6832,
        "name": "kworker/u4:0",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      {
        "pid": 7259,
        "name": "kworker/0:0",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      {
        "pid": 7351,
        "name": "kworker/1:1",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      {
        "pid": 7467,
        "name": "kworker/u4:2",
        "user": 0,
        "state": 1026,
        "memory": 0
      },
      {
        "pid": 7468,
        "name": "kworker/0:2",
        "user": 0,
        "state": 1026,
        "memory": 0
      }
    ]
  },
  {
    "pid": 3,
    "name": "rcu_gp",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 4,
    "name": "rcu_par_gp",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 5,
    "name": "netns",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 7,
    "name": "kworker/0:0H",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 9,
    "name": "kworker/0:1H",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 10,
    "name": "mm_percpu_wq",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 11,
    "name": "rcu_tasks_rude_",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 12,
    "name": "rcu_tasks_trace",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 13,
    "name": "ksoftirqd/0",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 14,
    "name": "rcu_sched",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 15,
    "name": "migration/0",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 16,
    "name": "idle_inject/0",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 18,
    "name": "cpuhp/0",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 19,
    "name": "cpuhp/1",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 20,
    "name": "idle_inject/1",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 21,
    "name": "migration/1",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 22,
    "name": "ksoftirqd/1",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 24,
    "name": "kworker/1:0H",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 25,
    "name": "kdevtmpfs",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 26,
    "name": "inet_frag_wq",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 27,
    "name": "kauditd",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 29,
    "name": "khungtaskd",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 30,
    "name": "oom_reaper",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 31,
    "name": "writeback",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 32,
    "name": "kcompactd0",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 33,
    "name": "ksmd",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 34,
    "name": "khugepaged",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 80,
    "name": "kintegrityd",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 81,
    "name": "kblockd",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 82,
    "name": "blkcg_punt_bio",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 83,
    "name": "tpm_dev_wq",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 84,
    "name": "ata_sff",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 85,
    "name": "md",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 86,
    "name": "edac-poller",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 87,
    "name": "devfreq_wq",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 88,
    "name": "watchdogd",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 90,
    "name": "kworker/1:1H",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 92,
    "name": "kswapd0",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 93,
    "name": "ecryptfs-kthrea",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 95,
    "name": "kthrotld",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 96,
    "name": "acpi_thermal_pm",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 97,
    "name": "scsi_eh_0",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 98,
    "name": "scsi_tmf_0",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 99,
    "name": "nvme-wq",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 100,
    "name": "nvme-reset-wq",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 101,
    "name": "nvme-delete-wq",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 102,
    "name": "vfio-irqfd-clea",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 104,
    "name": "mld",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 105,
    "name": "ipv6_addrconf",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 115,
    "name": "kstrp",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 118,
    "name": "zswap-shrink",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 119,
    "name": "kworker/u5:0",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 126,
    "name": "charger_manager",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 127,
    "name": "jbd2/sda1-8",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 128,
    "name": "ext4-rsv-conver",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 131,
    "name": "hwrng",
    "user": 0,
    "state": 1,
    "memory": 0,
    "children": []
  },
  {
    "pid": 170,
    "name": "systemd-journal",
    "user": 0,
    "state": 1,
    "memory": 3271,
    "children": []
  },
  {
    "pid": 201,
    "name": "systemd-udevd",
    "user": 0,
    "state": 1,
    "memory": 1303,
    "children": []
  },
  {
    "pid": 228,
    "name": "cryptd",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 318,
    "name": "kaluad",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 319,
    "name": "kmpath_rdacd",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 320,
    "name": "kmpathd",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 321,
    "name": "kmpath_handlerd",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 322,
    "name": "multipathd",
    "user": 0,
    "state": 1,
    "memory": 4500,
    "children": []
  },
  {
    "pid": 438,
    "name": "systemd-network",
    "user": 100,
    "state": 1,
    "memory": 1762,
    "children": []
  },
  {
    "pid": 442,
    "name": "systemd-resolve",
    "user": 101,
    "state": 1,
    "memory": 2795,
    "children": []
  },
  {
    "pid": 477,
    "name": "accounts-daemon",
    "user": 0,
    "state": 1,
    "memory": 2082,
    "children": []
  },
  {
    "pid": 490,
    "name": "chronyd",
    "user": 112,
    "state": 1,
    "memory": 624,
    "children": [
      { "pid": 493, "name": "chronyd", "user": 112, "state": 1, "memory": 46 }
    ]
  },
  {
    "pid": 493,
    "name": "chronyd",
    "user": 112,
    "state": 1,
    "memory": 46,
    "children": []
  },
  {
    "pid": 495,
    "name": "dbus-daemon",
    "user": 103,
    "state": 1,
    "memory": 1032,
    "children": []
  },
  {
    "pid": 498,
    "name": "google_osconfig",
    "user": 0,
    "state": 1,
    "memory": 3855,
    "children": []
  },
  {
    "pid": 503,
    "name": "networkd-dispat",
    "user": 0,
    "state": 1,
    "memory": 4448,
    "children": []
  },
  {
    "pid": 505,
    "name": "rsyslogd",
    "user": 104,
    "state": 1,
    "memory": 1129,
    "children": []
  },
  {
    "pid": 515,
    "name": "snapd",
    "user": 0,
    "state": 1,
    "memory": 9407,
    "children": []
  },
  {
    "pid": 521,
    "name": "atd",
    "user": 0,
    "state": 1,
    "memory": 555,
    "children": []
  },
  {
    "pid": 525,
    "name": "containerd",
    "user": 0,
    "state": 1,
    "memory": 8745,
    "children": []
  },
  {
    "pid": 562,
    "name": "google_guest_ag",
    "user": 0,
    "state": 1,
    "memory": 3380,
    "children": []
  },
  {
    "pid": 583,
    "name": "polkitd",
    "user": 0,
    "state": 1,
    "memory": 2228,
    "children": []
  },
  {
    "pid": 584,
    "name": "agetty",
    "user": 0,
    "state": 1,
    "memory": 461,
    "children": []
  },
  {
    "pid": 609,
    "name": "agetty",
    "user": 0,
    "state": 1,
    "memory": 428,
    "children": []
  },
  {
    "pid": 613,
    "name": "unattended-upgr",
    "user": 0,
    "state": 1,
    "memory": 4752,
    "children": []
  },
  {
    "pid": 711,
    "name": "dockerd",
    "user": 0,
    "state": 1,
    "memory": 20828,
    "children": []
  },
  {
    "pid": 719,
    "name": "sshd",
    "user": 0,
    "state": 1,
    "memory": 1634,
    "children": [
      { "pid": 5961, "name": "sshd", "user": 0, "state": 1, "memory": 2097 }
    ]
  },
  {
    "pid": 733,
    "name": "bpfilter_umh",
    "user": 0,
    "state": 1,
    "memory": 143,
    "children": []
  },
  {
    "pid": 873,
    "name": "systemd-logind",
    "user": 0,
    "state": 1,
    "memory": 1753,
    "children": []
  },
  {
    "pid": 876,
    "name": "cron",
    "user": 0,
    "state": 1,
    "memory": 676,
    "children": []
  },
  {
    "pid": 5928,
    "name": "kworker/1:0",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 5929,
    "name": "kworker/u4:1",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 5961,
    "name": "sshd",
    "user": 0,
    "state": 1,
    "memory": 2097,
    "children": [
      { "pid": 6062, "name": "sshd", "user": 1001, "state": 1, "memory": 1442 }
    ]
  },
  {
    "pid": 5964,
    "name": "systemd",
    "user": 1001,
    "state": 1,
    "memory": 2172,
    "children": [
      {
        "pid": 5965,
        "name": "(sd-pam)",
        "user": 1001,
        "state": 1,
        "memory": 1139
      }
    ]
  },
  {
    "pid": 5965,
    "name": "(sd-pam)",
    "user": 1001,
    "state": 1,
    "memory": 1139,
    "children": []
  },
  {
    "pid": 5966,
    "name": "kworker/1:2",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 6062,
    "name": "sshd",
    "user": 1001,
    "state": 1,
    "memory": 1442,
    "children": [
      { "pid": 6063, "name": "bash", "user": 1001, "state": 1, "memory": 1251 }
    ]
  },
  {
    "pid": 6063,
    "name": "bash",
    "user": 1001,
    "state": 1,
    "memory": 1251,
    "children": [
      { "pid": 7484, "name": "cat", "user": 1001, "state": 0, "memory": 146 }
    ]
  },
  {
    "pid": 6370,
    "name": "kworker/0:1",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 6832,
    "name": "kworker/u4:0",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 7259,
    "name": "kworker/0:0",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 7351,
    "name": "kworker/1:1",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 7467,
    "name": "kworker/u4:2",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 7468,
    "name": "kworker/0:2",
    "user": 0,
    "state": 1026,
    "memory": 0,
    "children": []
  },
  {
    "pid": 7484,
    "name": "cat",
    "user": 1001,
    "state": 0,
    "memory": 146,
    "children": []
  }
]
`

var conn = MySQLConn()

func MySQLConn() *sql.DB {
	db, err := sql.Open("mysql", "admin:7P4,;C<8Io^jG&p&@tcp(35.202.232.209)/modules")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error en la conexión a la base de datos")
	} else {
		fmt.Println("Connected to MySQL")
	}
	return db
}

func postRam(data string) {
	fmt.Println("Insertando RAM en la base de datos")
	fmt.Println(data)
	var ram Ram
	json.Unmarshal([]byte(data), &ram)
	fmt.Println(ram)

	stmt, err := conn.Prepare("INSERT INTO ram(total, used, free, percentage) VALUES(?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	_, err = stmt.Exec(ram.Total, ram.Used, ram.Free, ram.Percentage)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Ram insertada")
}

func postProcesses(data string) {
	fmt.Println("Insertando procesos en la base de datos")
	fmt.Println(data)
	var processes []Process
	json.Unmarshal([]byte(data), &processes)
	fmt.Println(processes)
	for _, process := range processes {
		stmt, err := conn.Prepare("INSERT INTO process(pid, name, user, state, memory) VALUES(?, ?, ?, ?, ?)")
		if err != nil {
			fmt.Println(err)
		}
		_, err = stmt.Exec(process.Pid, process.Name, process.User, process.State, (process.Memory/(1024.0*1024.0))*100.0)
		if err != nil {
			fmt.Println(err)
		}
		for _, child := range process.Children {
			stmt, err := conn.Prepare("INSERT INTO process(pid, name, user, state, memory, pid_padre) VALUES(?, ?, ?, ?, ?, ?)")
			if err != nil {
				fmt.Println(err)
			}
			_, err = stmt.Exec(child.Pid, child.Name, child.User, child.State, (child.Memory/(1024.0*1024.0))*100.0, process.Pid)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	fmt.Println("Procesos insertados")

}

func main() {
	/*for {
		fmt.Println("Obteniendo datos ...")
		cmd := exec.Command("sh", "-c", "cat /proc/ram_201709450")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		postRam(string(out[:]))
		// ------------------------------------------------------------------
		time.Sleep(1 * time.Second)
		proccess := exec.Command("sh", "-c", "cat /proc/cpu_201709450")
		out, err = proccess.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		postProcesses(string(out[:]))
		time.Sleep(8 * time.Second)
	}*/
	var process []Process
	err := json.Unmarshal([]byte(body), &process)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, p := range process {
			fmt.Println(p)
		}
	}
}
