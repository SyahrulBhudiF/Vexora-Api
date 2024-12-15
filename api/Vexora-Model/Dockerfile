# Production stage
FROM python:3.10-slim-bookworm AS runtime

# Install required libraries for runtime
RUN apt-get update && apt-get install -y libwebp7 libwebp-dev ca-certificates libgl1 libglib2.0-0

# Set the working directory
WORKDIR /

# Install Python dependencies
COPY api/Vexora-Model/requirements.txt requirements.txt
RUN pip3 install --no-cache-dir -r requirements.txt --break-system-packages

COPY api/Vexora-Model/ .

CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]