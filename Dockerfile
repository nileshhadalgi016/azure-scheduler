FROM registry.access.redhat.com/ubi8/go-toolset:1.18.4-8.1669838000

# Labels
LABEL name="awsvmscheduler" \
    maintainer="nileshhadalgi3@gmail.com" \
    vendor="Nilesh" \
    version="1.0.0" \
    release="1" \
    summary="This service enables Azure cloud vm start/stop." \
    description="This service enables Azure cloud vm start/stop."

ENV command ""
ENV AZURE_VM_NAME ""
ENV AZURE_RG_NAME ""

# copy code to the build path
COPY . .

RUN go mod download

RUN go build -o main

CMD ["sh", "-c", "./main -c $command -vm $AZURE_VM_NAME -rg $AZURE_RG_NAME  "]