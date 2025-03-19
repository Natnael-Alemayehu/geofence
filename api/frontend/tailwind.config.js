// tailwind.config.js
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'primary-bg': '#1a3a32',
        'header-bg': '#0f2a24',
        'brand-blue': '#4a90e2',
        'module-bg': '#2a4a40',
      },
      borderRadius: {
        'custom': '10px',
      }
    },
  },
  plugins: [],
}