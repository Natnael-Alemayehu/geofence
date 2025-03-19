// src/services/api.js
import axios from 'axios';

// Replace with your actual API URL
const API_URL = 'http://localhost:3000/v1';

const verifyLocation = async (locationData) => {
  try {
    const response = await axios.post(`${API_URL}/verify_location`, locationData);
    return response.data;
  } catch (error) {
    console.error('Error verifying location:', error);
    throw error;
  }
};

export { verifyLocation };