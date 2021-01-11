

```bash
sudo lsblk
NAME   MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
sdb      8:16   0   50G  0 disk
sdc      8:32   0   10G  0 disk
sda      8:0    0  100G  0 disk
└─sda1   8:1    0  100G  0 part /

mkdir -p /pod-50g
mkdir -p /pod-10g

sudo mkfs.ext4 -m 0 -F -E lazy_itable_init=0,lazy_journal_init=0,discard /dev/sdb
sudo mkfs.ext4 -m 0 -F -E lazy_itable_init=0,lazy_journal_init=0,discard /dev/sdc

sudo mount -o discard,defaults /dev/sdb /pod-50g
sudo mount -o discard,defaults /dev/sdc /pod-10g

sudo lsblk


sudo chmod a+w /pod-50g
sudo chmod a+w /pod-10g

echo UUID=`sudo blkid -s UUID -o value /dev/sdb` /pod-50g ext4 discard,defaults,nofail 0 2 | sudo tee -a /etc/fstab
echo UUID=`sudo blkid -s UUID -o value /dev/sdc` /pod-10g ext4 discard,defaults,nofail 0 2 | sudo tee -a /etc/fstab

cat /etc/fstab

# LABEL=cloudimg-rootfs	/	 ext4	defaults	0 0
# UUID=2289b0ce-0cd2-4625-857a-daab61a5ac35 /pod-50g ext4 discard,defaults,nofail 0 2
# UUID=a83da77a-c8b8-4ebb-89ec-0271261c9b5c /pod-10g ext4 discard,defaults,nofail 0 2
```