#!/bin/bash
write_knockd_conf() {
cat <<EOF > /etc/knockd.conf
[options]
    logfile = /var/log/knockd.log
[SSH]
    sequence = $1
    tcpflags = syn
    seq_timeout = 5
    start_command = /sbin/iptables -I ssh_access 1 -s %IP% -p tcp --dport 22 -j ACCEPT
    cmd_timeout = 20
    stop_command = /sbin/iptables -D ssh_access -s %IP% -p tcp --dport 22 -j ACCEPT
EOF
}

while true 
do
    EPOCH=$(date +%s)
    if [ $(($EPOCH % 30)) == "0" ]
    then
        CODE=$(cat /etc/knockd/knockd.code.01)
        SEQ1=$(oathtool --base32 --totp "$CODE" -d 8|sed -e 's/\(....\)\(....\)/\1,\2/g')
        CODE=$(cat /etc/knockd/knockd02.code.02)
        SEQ2=$(oathtool --base32 --totp "$CODE" -d 8|sed -e 's/\(....\)\(....\)/\1,\2/g')
        write_knockd_conf $SEQ1,$SEQ2
        iptables -F ssh_access
        service knockd reload
    fi
    if [ $(($EPOCH % 600)) == "0" ]
    then
        aws s3 sync s3://knockd/ /etc/knockd/
    fi
    sleep 1
done
