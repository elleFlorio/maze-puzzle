FROM ubuntu
COPY maze-puzzle /bin/maze-puzzle
ENTRYPOINT ["maze-puzzle"]