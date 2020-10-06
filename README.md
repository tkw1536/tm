# tm - Tiny Mirror

![CI Status](https://github.com/tkw1536/tm/workflows/Publish%20Docker%20Image/badge.svg)

This folder contains a tiny image which serves and keeps a mirror of a remote rsync location up-to-date. 
It is written in go, bundled using docker. 
It is available as a GitHub Package.

Usage:

```
    docker run -v $VOLUME:/data/ -t -i --rm -e DELAY=$DELAY -e REMOTE=$REMOTE -p 8080:8080 docker.pkg.github.com/tkw1536/tm/tm:latest
```

Where:

- `$DELAY` is the interval between running syncs from the remote, e.g. `1h`
- `$REMOTE` is the remote to sync with, starting with `rsync://`

The code is licensed under the Unlicense, hence in the public domain.