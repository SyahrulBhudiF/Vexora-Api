from dotenv import load_dotenv
import os

# Load environment variables
load_dotenv()

# Get secret key from environment variable
API_SECRET_KEY = os.getenv('SECRET_KEY')
if not API_SECRET_KEY:
    raise ValueError("API_SECRET_KEY must be set in .env file")
