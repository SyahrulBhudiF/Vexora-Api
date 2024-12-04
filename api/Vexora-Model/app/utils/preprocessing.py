import cv2
import numpy as np
from typing import Tuple


class ImagePreprocessor:
    def __init__(self, target_size: Tuple[int, int] = (128, 128)):
        self.target_size = target_size
        self.face_cascade = cv2.CascadeClassifier(
            cv2.data.haarcascades + 'haarcascade_frontalface_default.xml'
        )

    def preprocess_image(self, img_path: str) -> np.ndarray:
        """
        Preprocesses image for emotion detection.

        Args:
            img_path: Path to input image

        Returns:
            Preprocessed image array

        Raises:
            ValueError: If face detection or preprocessing fails
        """

        try:
            # Read image in color
            img = cv2.imread(img_path, cv2.IMREAD_COLOR)
            if img is None:
                raise ValueError("Gambar tidak ditemukan atau tidak dapat dimuat.")

            # Convert to grayscale after reading in color
            gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)

            faces = self.face_cascade.detectMultiScale(
                gray,
                scaleFactor=1.1,
                minNeighbors=5,
                minSize=(30, 30)
            )

            if len(faces) == 0:
                raise ValueError("Wajah tidak terdeteksi dalam gambar.")

            x, y, w, h = faces[0]
            face = gray[y:y + h, x:x + w]

            face_resized = cv2.resize(
                face,
                self.target_size,
                interpolation=cv2.INTER_LANCZOS4
            )
            face_normalized = face_resized.astype('float32') / 255.0

            # Repeat the grayscale channel 3 times to simulate RGB
            face_final = np.repeat(face_normalized[..., np.newaxis], 3, axis=-1)
            face_final = np.expand_dims(face_final, axis=0)
            return face_final

        except Exception as e:
            raise ValueError(f"Kesalahan dalam memproses gambar: {str(e)}")
