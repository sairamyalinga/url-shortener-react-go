import axios from "axios";
export const handleCopyToClipboard = (data) => {
  navigator.clipboard
    .writeText(data)
    .then(() => {
      alert("ShortURL copied to clipboard!");
    })
    .catch((err) => {
      console.error("Failed to copy text: ", err);
    });
};

let API = axios.create({
  baseURL: `${import.meta.env.VITE_API_URL}/api`,
  withCredentials: true,
});

API.interceptors.request.use((config) => {
  config.headers.Authorization = `Bearer ${localStorage.getItem('token')}`;
  return config;
}, (error) => {
  return Promise.reject(error);
});

API.interceptors.response.use(
  response => {
    return response;
  },
  error => {
    if (error.response && error.response.status === 401) { 
      localStorage.removeItem('token');
      window.location.href = '/';
    }
    if (error.response && error.response.data) {
      return Promise.reject(error.response.data.message);
    }
    return Promise.reject(error.message);
  }
);


export default API;



