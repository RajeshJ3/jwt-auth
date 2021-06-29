# Dockerized JWT Auth [GoLang]

Hey developers, welcome! This project is a boiler plate for Authentication and Authorization in GoLang using JSON Web Token(JWT) containerized using Docker.

## What you end up with 🥳

- Go project running at ports 80, 443, 8000
  - NGINX production server listening at POST 80, 443
  - Development server listening at POST 8000
- Persistent Postgres volume

## Prerequisites 😰

- Docker and docker-compose installed
- Basic understanding of Docker
- Cloudflare Account (for SSL)
- Internet connectivity

---

## Let's Go 🏃‍♂️

- Clone the project

```bash
git clone https://github.com/rajeshj3/jwt-auth
cd jwt-auth
```

- Configurations

  - For **SSL**

    Create directory **nginx/ssl** and add 3 files

    ```bash
    touch ssl/cert.pem
    touch ssl/key.pem
    touch ssl/cloudflare.crt
    ```

    You'll find thes files in your cloudflare dashboard

    **Note:** If you don't want your server to server over SSL, Remove the second server block from **nginx/nginx.conf** ie.

    ```.conf
    server {
        listen 443 ssl http2;

        ssl_certificate         /etc/ssl/cert.pem;
        ssl_certificate_key     /etc/ssl/key.pem;
        ssl_client_certificate  /etc/ssl/cloudflare.crt;
        ssl_verify_client on;

        location / {
            proxy_pass http://app;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header Host $host;
            proxy_redirect off;
        }
    }
    ```

    And, also remove **COPY ssl/** commands from **nginx/Dockerfile**, ie.

    ```Dockerfile
    COPY ssl/cert.pem /etc/ssl/cert.pem
    COPY ssl/key.pem /etc/ssl/key.pem
    COPY ssl/cloudflare.crt /etc/ssl/cloudflare.crt
    ```

- Build Images

```bash
docker-compose --build -d
```

- Run the containers in detched mode

```bash
docker-compose up -d
```

Awesome 🥳🥳 Our server is now running at POSTs 80, 443, 8000

_Note:_ 443 might not work 🚫 if server is not located at the domain you've created SSL for.

---

## Accessing PostgreSQL Database 🧱

- Accessing psql shell

```bash
docker exec -it jwt_auth_db psql -U postgres
```

**Note:** _postgres_ is default user.
Run any postgres commands here, eg

```psql
# \l
```

## Deleting everything 🗑️

- Delete containers

```bash
docker-compose down
```

- Deleting Images

```bash
docker rmi xxxxxxxxxx xxxxxxxxxx xxxxxxxxxx
```

**Note:** At place of _xxxxxxxxxx_ place image IDs of our containers

- Deleting Volumes

```bash
docker volume prune -f
```

---

Thank you 😇

Cheers
