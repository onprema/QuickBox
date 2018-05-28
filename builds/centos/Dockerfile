FROM centos:6
RUN \
	yum -y install openssh-clients openssh-server && \
	yum -y clean all && \
	echo "root:root" | chpasswd && \
	mkdir -p /var/run/sshd /root/.ssh
COPY entrypoint.sh /tmp
EXPOSE 22
ENTRYPOINT ["/tmp/entrypoint.sh"]

