FROM ubuntu
ADD ./cmd/uniclub_project /bin/uniclub_project 
COPY .env . 
CMD ["/bin/uniclub_project"]
