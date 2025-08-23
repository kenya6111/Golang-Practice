FROM ubuntu

RUN groupadd -r appgroup && useradd -r -g appgroup testuser

USER testuser

EXPOSE 8089

CMD ["whoami"]