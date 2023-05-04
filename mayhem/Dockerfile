FROM golang:1.18 as builder

COPY . /teacup
WORKDIR /teacup/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN go get github.com/AdaLogics/go-fuzz-headers
RUN apt update && apt install -y clang

RUN cd fuzz_surf_awsu && go-fuzz-build -libfuzzer -o fuzz_surf_awsu.a && \
    clang -fsanitize=fuzzer fuzz_surf_awsu.a -o fuzz_surf_awsu.libfuzzer

RUN cd fuzz_surf_jsonutil && go-fuzz-build -libfuzzer -o fuzz_surf_jsonutil.a && \
    clang -fsanitize=fuzzer fuzz_surf_jsonutil.a -o fuzz_surf_jsonutil.libfuzzer

RUN cd fuzz_surf_proto && go-fuzz-build -libfuzzer -o fuzz_surf_proto.a && \
    clang -fsanitize=fuzzer fuzz_surf_proto.a -o fuzz_surf_proto.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /teacup/mayhem/fuzz_surf_awsu/fuzz_surf_awsu.libfuzzer /
COPY --from=builder /teacup/mayhem/fuzz_surf_jsonutil/fuzz_surf_jsonutil.libfuzzer /
COPY --from=builder /teacup/mayhem/fuzz_surf_proto/fuzz_surf_proto.libfuzzer /