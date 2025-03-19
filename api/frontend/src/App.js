// src/App.js
import React from 'react';
import VerificationForm from './components/VerificationForm';

function App() {
  return (
    <div className="min-h-screen py-8 px-4">
      <header className="max-w-2xl mx-auto mb-8 text-center">
        <h1 className="text-3xl font-bold text-brand-blue mb-2">Location Verification System</h1>
        <p className="text-gray-300">Verify if coordinates are inside or outside a specified location</p>
      </header>
      
      <main>
        <VerificationForm />
      </main>
      
      <footer className="max-w-2xl mx-auto mt-12 pt-6 border-t border-gray-700 text-center text-gray-400 text-sm">
        &copy; {new Date().getFullYear()} Location Verification System
      </footer>
    </div>
  );
}

export default App;