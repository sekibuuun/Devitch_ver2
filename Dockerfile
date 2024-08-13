FROM node:latest

# Install bun
RUN curl -fsSL https://bun.sh/install | bash
ENV PATH="/root/.bun/bin:${PATH}"

# Install Go
RUN apt-get update && apt-get install -y golang-go

# Install git
RUN apt-get install -y git

# Install sudo
RUN apt-get install -y sudo

# Install lefthook
RUN curl -1sLf 'https://dl.cloudsmith.io/public/evilmartians/lefthook/setup.deb.sh' | sudo -E bash
RUN sudo apt install lefthook

# Set Go environment variables
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH

# Set working directory
WORKDIR /app

# Copy package.json and install dependencies
COPY frontend/package.json ./frontend/
RUN cd frontend && bun install

# Copy Go files
COPY backend ./backend

# Expose ports
EXPOSE 3000 8080 3306

#Install Air
RUN cd backend && go install github.com/air-verse/air@latest

# Start air
CMD ["air", "-c", "/backend/.air.toml"]

# Appliance terminal
RUN echo "source /usr/share/bash-completion/completions/git" >> ~/.bashrc

WORKDIR /usr/share/bash-completion/completions

RUN curl -O https://raw.githubusercontent.com/git/git/master/contrib/completion/git-prompt.sh
RUN curl -O https://raw.githubusercontent.com/git/git/master/contrib/completion/git-completion.bash
RUN chmod a+x git*.*

RUN ls -l $PWD/git*.* | awk '{print "source "$9}' >> ~/.bashrc

RUN echo "GIT_PS1_SHOWDIRTYSTATE=true" >> ~/.bashrc
RUN echo "GIT_PS1_SHOWUNTRACKEDFILES=true" >> ~/.bashrc
RUN echo "GIT_PS1_SHOWUPSTREAM=auto" >> ~/.bashrc

RUN echo 'export PS1="\[\033[01;32m\]\u@\h\[\033[01;33m\] \w \[\033[01;31m\]\$(__git_ps1 \"(%s)\") \\n\[\033[01;34m\]\\$ \[\033[00m\]"' >> ~/.bashrc

# Install bash-completion and other necessary packages
RUN sudo apt install bash-completion

RUN echo "source /usr/share/bash-completion/bash_completion" >> ~/.bashrc

# Set working directory
WORKDIR /app