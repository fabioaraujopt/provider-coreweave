FROM alpine:3.23.4
RUN apk --no-cache add ca-certificates

ARG TARGETOS
ARG TARGETARCH

ADD "bin/${TARGETOS}_${TARGETARCH}/provider" /usr/local/bin/provider

ENV USER_ID=65532

# All CoreWeave resources are Terraform Plugin Framework resources that Upjet
# runs in-process (no-fork architecture). The image therefore needs neither
# the Terraform CLI nor the terraform-provider-coreweave binary. The values
# below are only reported as metadata and required by the provider's flags.
ARG TERRAFORM_VERSION
ARG TERRAFORM_PROVIDER_SOURCE
ARG TERRAFORM_PROVIDER_VERSION

ENV TERRAFORM_VERSION=${TERRAFORM_VERSION}
ENV TERRAFORM_PROVIDER_SOURCE=${TERRAFORM_PROVIDER_SOURCE}
ENV TERRAFORM_PROVIDER_VERSION=${TERRAFORM_PROVIDER_VERSION}

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
