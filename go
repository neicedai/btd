apt install xfce4 -y
apt install xrdp -y
apt install virt-manager -y
wget http://95.217.112.160/btd.tar.gz
tar xvf btd.tar.gz
mv new-btd.qcow2 /var/lib/libvirt/images
cd /var/lib/libvirt/images
cp new-btd.qcow2 temp.qcow2
qemu-img  create -f qcow2 /home/12T.qcow2 -o size=12726G,preallocation=metadata
