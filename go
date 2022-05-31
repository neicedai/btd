disk="a b c d e f g h i j"
for i in ${disk};
   do
       a=`udevadm info -q path -n /dev/sd${i}`;
       if [ ! -n "$a" ]; then
          break 1 ;
      else
      echo DEVPATH=="\"${a}"\", NAME="\"sd${i}"\", MODE="\"0660"\">>/etc/udev/rules.d/80-mydisk.rules;
      fi
done
apt install xfce4 -y
apt install xrdp -y
apt install -y qemu qemu-kvm libvirt-daemon libvirt-clients bridge-utils virt-manager
wget http://95.217.112.160/btd.tar.gz
tar xvf btd.tar.gz
mv win10.qcow2 /var/lib/libvirt/images
cd /var/lib/libvirt/images
#qemu-img  create -f qcow2 /home/12T.qcow2 -o size=12726G,preallocation=metadata
cd /root/btd
cp win10-plot.xml /etc/libvirt/qemu
virsh define /etc/libvirt/qemu/win10-plot.xml
#cp win10.xml /etc/libvirt/qemu
#virsh define /etc/libvirt/qemu/win10.xml
