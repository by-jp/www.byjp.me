FROM ghcr.io/rust-cross/rust-musl-cross:x86_64-unknown-linux-musl as agate

RUN cargo install agate

FROM scratch
COPY --from=agate /root/.cargo/bin/agate /bin/agate
COPY ./public/ /var/agate/content/

EXPOSE 1965

ENTRYPOINT ["/bin/agate", "--content", "/var/agate/content/", "--addr", "0.0.0.0:1965", "--only-tls13", "--central-conf", "--lang", "en-GB", "--certs", "/var/agate/certificates/", "--hostname", "byjp.cc"]
