FROM bash

COPY wontexit.sh /wontexit.sh

ENTRYPOINT wontexit.sh

