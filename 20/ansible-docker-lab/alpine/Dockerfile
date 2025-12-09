# Use Alpine as the base image
FROM alpine:latest

# Install OpenSSH server and other utilities
RUN apk update && \
    apk add --no-cache openssh && \
    apk add --no-cache python3 && \
    apk add --no-cache sudo && \
    mkdir -p /var/run/sshd

# Create a new user 'ansible'
RUN adduser -D -s /bin/ash ansible

# Set password for 'ansible' (optional, passwordless SSH will be used)
RUN echo "ansible:password" | chpasswd

# Add 'ansible' user to sudoers for passwordless sudo access
RUN echo 'ansible ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

# Generate SSH host keys
RUN ssh-keygen -A


# Create .ssh directory for the ansible user and set permissions
RUN mkdir -p /home/ansible/.ssh && \
    chown ansible:ansible /home/ansible/.ssh && \
    chmod 700 /home/ansible/.ssh

# Copy the host's public SSH key to the container (assuming you have id_ed25519.pub in build context)
COPY ./id_ed25519.pub /home/ansible/.ssh/authorized_keys

# Set correct ownership and permissions for authorized_keys file
RUN chown ansible:ansible /home/ansible/.ssh/authorized_keys && \
    chmod 600 /home/ansible/.ssh/authorized_keys

# Expose SSH port
EXPOSE 22

# Start SSH service and keep the container running
CMD ["/usr/sbin/sshd", "-D"]

