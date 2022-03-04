FROM alpine

RUN mkdir /dist
COPY zuoye  /dist   
EXPOSE 7000
CMD [ "/dist/zuoye" ]