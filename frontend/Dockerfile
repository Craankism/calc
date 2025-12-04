## Frontend-only Dockerfile
## Serves static frontend files using nginx

FROM nginx:alpine

# Remove default nginx config
RUN rm /etc/nginx/conf.d/default.conf

# Copy custom nginx configuration
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copy frontend static files into nginx webroot
# Expecting these files to exist in the `frontend/` build context:
# - index.html
# - dev.index.html
# - static/ (CSS, JS, images)
COPY index.html /usr/share/nginx/html/index.html
COPY static /usr/share/nginx/html/static

# Expose HTTP port
EXPOSE 80

# Start nginx in foreground
CMD ["nginx", "-g", "daemon off;"]