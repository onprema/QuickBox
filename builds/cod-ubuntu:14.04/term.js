# Ubuntu 14:04 with sshd
FROM ubuntu:14.04
MAINTAINER "Lee Gaines" "@eightlimbed"
RUN apt-get update && apt-get install -y openssh-server git
RUN mkdir /var/run/sshd /root/.ssh
RUN echo "root:root" | chpasswd
RUN sed -ri 's/^#?PermitRootLogin\s+.*/PermitRootLogin yes/' /etc/ssh/sshd_config
EXPOSE 22
CMD ["/usr/sbin/sshd", "-D"]
RUN git clone https://github.com/chjj/term.js.git
