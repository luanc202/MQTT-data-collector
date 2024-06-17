# How to run

### Dependencies

Only Docker, unless you want to run everything in your machine without the use of Docker.

### Build steps

Build the `data-collector-server` with

```bash
docker build . -t data-collector-server
```

And then start both containers with

```bash
docker compose up
```

Make sure to replace the environment variables needed for it to work. An example file is provided as `.env.example` in this repo.
