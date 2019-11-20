FROM golang

COPY ./web ./
RUN cd ./src/
CMD ./wiki


