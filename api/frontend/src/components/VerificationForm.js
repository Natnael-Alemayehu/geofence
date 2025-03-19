// src/components/VerificationForm.js
import React, { useState } from 'react';
import { verifyLocation } from '../services/api';

const VerificationForm = () => {
  const [formData, setFormData] = useState({
    location_id: '',
    latitude: '',
    longitude: ''
  });

  const [verificationResult, setVerificationResult] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: name === 'latitude' || name === 'longitude' ? parseFloat(value) || '' : value
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    // Form validation
    if (!formData.location_id || !formData.latitude || !formData.longitude) {
      setError('All fields are required');
      return;
    }
    
    setIsLoading(true);
    setError(null);
    
    try {
      const result = await verifyLocation(formData);
      setVerificationResult(result);
    } catch (err) {
      setError(err.response?.data?.message || 'An error occurred while verifying the location');
    } finally {
      setIsLoading(false);
    }
  };

  const handleReset = () => {
    setFormData({
      location_id: '',
      latitude: '',
      longitude: ''
    });
    setVerificationResult(null);
    setError(null);
  };

  return (
    <div className="max-w-2xl mx-auto mt-10">
      <div className="mb-8">
        <div className="section-header">
          <h2 className="text-xl">Location Verification</h2>
        </div>
        
        <div className="card">
          <form onSubmit={handleSubmit}>
            <div className="mb-4">
              <label className="block mb-2 text-sm font-medium">Location ID</label>
              <input
                type="text"
                name="location_id"
                value={formData.location_id}
                onChange={handleChange}
                className="input-field w-full"
                placeholder="Enter location ID"
              />
            </div>
            
            <div className="mb-4">
              <label className="block mb-2 text-sm font-medium">Latitude</label>
              <input
                type="number"
                step="any"
                name="latitude"
                value={formData.latitude}
                onChange={handleChange}
                className="input-field w-full"
                placeholder="Enter latitude"
              />
            </div>
            
            <div className="mb-6">
              <label className="block mb-2 text-sm font-medium">Longitude</label>
              <input
                type="number"
                step="any"
                name="longitude"
                value={formData.longitude}
                onChange={handleChange}
                className="input-field w-full"
                placeholder="Enter longitude"
              />
            </div>
            
            {error && (
              <div className="mb-4 p-3 bg-red-600 bg-opacity-30 border border-red-700 rounded-custom text-white">
                {error}
              </div>
            )}
            
            <div className="flex space-x-4">
              <button
                type="submit"
                className="btn btn-primary flex-1"
                disabled={isLoading}
              >
                {isLoading ? 'Verifying...' : 'Verify Location'}
              </button>
              <button
                type="button"
                onClick={handleReset}
                className="btn bg-gray-600 text-white hover:bg-gray-700"
              >
                Reset
              </button>
            </div>
          </form>
        </div>
      </div>
      
      {verificationResult && (
        <div className="mb-8">
          <div className="section-header">
            <h2 className="text-xl">Verification Result</h2>
          </div>
          
          <div className="card">
            <div className="mb-4">
              <div className="flex items-center mb-4">
                <span className={`w-3 h-3 rounded-full mr-2 ${
                  verificationResult.status === 'Inside' ? 'bg-green-500' : 'bg-red-500'
                }`}></span>
                <span className="text-xl font-semibold text-brand-blue">
                  Status: {verificationResult.status}
                </span>
              </div>
              
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-gray-400">Latitude</p>
                  <p>{verificationResult.latitude}</p>
                </div>
                <div>
                  <p className="text-sm text-gray-400">Longitude</p>
                  <p>{verificationResult.longitude}</p>
                </div>
              </div>
              
              {verificationResult.location_name && (
                <div className="mt-4">
                  <p className="text-sm text-gray-400">Location Name</p>
                  <ul className="list-disc list-inside">
                    {verificationResult.location_name.map((name, index) => (
                      <li key={index}>{name}</li>
                    ))}
                  </ul>
                </div>
              )}
            </div>
          </div>
        </div>
      )}
      
      <div className="bg-header-bg rounded-custom p-4 mt-8">
        <h3 className="text-brand-blue font-semibold mb-2">Key Information</h3>
        <ul className="space-y-1 text-sm">
          <li>• "Inside" status indicates the coordinates are within the specified location</li>
          <li>• "Outside" status indicates the coordinates are outside the specified location</li>
          <li>• For accurate results, ensure coordinates are in decimal format</li>
        </ul>
      </div>
    </div>
  );
};

export default VerificationForm;