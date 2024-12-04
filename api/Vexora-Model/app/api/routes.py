# app/api/routes.py
from fastapi import APIRouter, File, UploadFile, HTTPException, Header
from fastapi.responses import JSONResponse
import tensorflow as tf
import os
from tempfile import NamedTemporaryFile
import shutil
from ..config import API_SECRET_KEY
from ..utils.preprocessing import ImagePreprocessor

router = APIRouter()


def verify_secret_key(x_secret_key: str) -> bool:
    """Verifies if the provided secret key matches the one in .env"""
    return x_secret_key == API_SECRET_KEY


class EmotionDetector:
    def __init__(self, model_path: str):
        try:
            self.model = tf.keras.models.load_model(model_path, compile=False)
            self.model.compile(
                optimizer='adam',
                loss='categorical_crossentropy',
                metrics=['accuracy']
            )
        except Exception as e:
            print(f"Error loading model: {str(e)}")
            raise Exception(f"Failed to load model: {str(e)}")

        self.preprocessor = ImagePreprocessor()
        self.emotions = {0: 'angry', 1: 'happy', 2: 'neutral', 3: 'sad'}

    def predict(self, image_path: str) -> str:
        try:
            processed_img = self.preprocessor.preprocess_image(image_path)
            predictions = self.model.predict(processed_img, verbose=0)
            predicted_class = predictions.argmax()
            return self.emotions.get(predicted_class, "Unknown")
        except Exception as e:
            print(f"Prediction error: {str(e)}")
            raise Exception(f"Failed to make prediction: {str(e)}")


MODEL_PATH = os.path.join(os.path.dirname(os.path.dirname(__file__)), 'models', 'model_cnn.keras')

if not os.path.exists(MODEL_PATH):
    raise Exception(f"Model file not found at {MODEL_PATH}")

try:
    detector = EmotionDetector(MODEL_PATH)
except Exception as e:
    print(f"Failed to initialize EmotionDetector: {str(e)}")
    raise


@router.post("/mood-detection")
async def detect_mood(image: UploadFile = File(...),
                      x_secret_key: str = Header(..., alias="X-Secret-Key")) -> JSONResponse:
    """
    Detect mood from image with secret key authentication.

    Args:
        image: Upload image file
        x_secret_key: Secret key for authentication (passed in header)

    Returns:
        JSON response with detected emotion or error message
    """
    # Verify secret key first
    if not verify_secret_key(x_secret_key):
        return JSONResponse(
            status_code=401,
            content={
                "success": False,
                "shouldNotify": True,
                "message": "Invalid secret key",
                "data": None
            }
        )

    try:
        # Save uploaded file temporarily
        with NamedTemporaryFile(delete=False) as temp_file:
            shutil.copyfileobj(image.file, temp_file)
            temp_path = temp_file.name

        try:
            # Process image and get prediction
            emotion = detector.predict(temp_path)

            return JSONResponse(
                content={
                    "success": True,
                    "data": emotion,
                }
            )

        except Exception as e:
            raise HTTPException(
                status_code=500,
                detail=f"Error processing image: {str(e)}"
            )

        finally:
            # Clean up temporary file
            if os.path.exists(temp_path):
                os.unlink(temp_path)

    except HTTPException as he:
        return JSONResponse(
            status_code=he.status_code,
            content={
                "success": False,
                "shouldNotify": True,
                "message": he.detail,
                "data": None
            }
        )
    except Exception as e:
        return JSONResponse(
            status_code=500,
            content={
                "success": False,
                "shouldNotify": True,
                "message": f"Unexpected error: {str(e)}",
                "data": None
            }
        )
