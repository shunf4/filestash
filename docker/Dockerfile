FROM debian:stable-slim
MAINTAINER mickael@kerjean.me

ARG http_proxy=http://192.168.1.14:8118
ARG https_proxy=http://192.168.1.14:8118

RUN apt-get update && \
    #################
    # Install
    apt-get install -y libglib2.0-0 curl gnupg && \
    cd /tmp/ && \
    curl -s https://f.home.rivage.tk/filestash_linux-amd64.tar | tar xf - && \
    mv filestash /app/ && \
    #################
    # Optional dependencies
    apt-get install -y curl emacs zip poppler-utils > /dev/null && \
    # org-mode: html export
    curl https://raw.githubusercontent.com/mickael-kerjean/filestash/master/server/.assets/emacs/htmlize.el > /usr/share/emacs/site-lisp/htmlize.el && \
    # org-mode: markdown export
    curl https://raw.githubusercontent.com/mickael-kerjean/filestash/master/server/.assets/emacs/ox-gfm.el > /usr/share/emacs/site-lisp/ox-gfm.el && \
    # org-mode: pdf export (with a light latex distribution)
    cd && apt-get install -y wget perl > /dev/null && \
    export CTAN_REPO="http://mirror.las.iastate.edu/tex-archive/systems/texlive/tlnet" && \
    curl -sL "https://yihui.name/gh/tinytex/tools/install-unx.sh" | sh && \
    mv ~/.TinyTeX /usr/share/tinytex && \
    /usr/share/tinytex/bin/x86_64-linux/tlmgr install wasy && \
    /usr/share/tinytex/bin/x86_64-linux/tlmgr install ulem && \
    /usr/share/tinytex/bin/x86_64-linux/tlmgr install marvosym && \
    /usr/share/tinytex/bin/x86_64-linux/tlmgr install wasysym && \
    /usr/share/tinytex/bin/x86_64-linux/tlmgr install xcolor && \
    /usr/share/tinytex/bin/x86_64-linux/tlmgr install listings && \
    /usr/share/tinytex/bin/x86_64-linux/tlmgr install parskip && \
    /usr/share/tinytex/bin/x86_64-linux/tlmgr install float && \
    /usr/share/tinytex/bin/x86_64-linux/tlmgr install wrapfig && \
    /usr/share/tinytex/bin/x86_64-linux/tlmgr install sectsty && \
    ln -s /usr/share/tinytex/bin/x86_64-linux/pdflatex /usr/local/bin/pdflatex && \
    apt-get purge -y --auto-remove perl wget && \
    # Cleanup
    find /usr/share/ -name 'doc' | xargs rm -rf && \
    find /usr/share/emacs -name '*.pbm' | xargs rm -f && \
    find /usr/share/emacs -name '*.png' | xargs rm -f && \
    find /usr/share/emacs -name '*.xpm' | xargs rm -f && \
    #################
    # Finalise the image
    useradd filestash && \
    chown -R filestash:filestash /app/ && \
    rm -rf /var/lib/apt/lists/* && \
    rm -rf /tmp/* && \
    timeout 1 /app/filestash | grep -q start

RUN cd /tmp && \
    curl -s https://f.home.rivage.tk/filestash_linux-amd64.tar | tar xf -

RUN cd /tmp && \
    ls -l /tmp && \
    ls -l /tmp/filestash && \
    ls -l /app && \
    rm -rf /app/filestash && \
    mv filestash/filestash /app/ && \
    chown -R filestash:filestash /app/ && \
    chmod +x /app/filestash && \
    rm -rf /var/lib/apt/lists/* && \
    timeout 1 /app/filestash | grep -q start

EXPOSE 8334
VOLUME ["/app/data/"]
WORKDIR "/app"
USER filestash
CMD ["/app/filestash"]
