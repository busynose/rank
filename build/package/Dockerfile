FROM makeblock/alpine

ARG tmp_api_version=default_value
ENV API_VERSION=$tmp_api_version

ADD next-maker /app/
ADD web /app/web

WORKDIR /app

EXPOSE 8000

CMD ["./next-maker"]