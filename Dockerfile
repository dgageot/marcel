FROM docker:dind
COPY marcel /usr/local/bin/marcel
ENTRYPOINT ["/usr/local/bin/marcel"]
