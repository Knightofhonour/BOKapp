FROM public.ecr.aws/lambda/provided:al2 as build
# install compiler
RUN yum install -y golang
RUN go env -w GOPROXY=direct
# cache dependencies
# ADD go.mod go.sum ./
# RUN go mod download
RUN mkdir /core
COPY core/ /core

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./
# build
ADD . .
RUN GOARCH=amd64 GOOS=linux go build -o /random_entry
# copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /random_entry /random_entry
RUN chmod a+x /random_entry
ENTRYPOINT [ "/random_entry" ]