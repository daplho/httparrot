FROM golang:1.8-onbuild

CMD ["app", "-body=true", "-header=true" , "-port=5555"]
