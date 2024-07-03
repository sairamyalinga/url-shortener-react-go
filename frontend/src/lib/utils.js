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

const token = localStorage.getItem('token')
console.log(token)
let API = axios.create({
  baseURL: "http://localhost:5050/api",
  withCredentials: true,
  headers: {
      Authorization: `Bearer ${token}`

  }
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
      return Promise.reject(error.response.data);
    }
    return Promise.reject(error.message);
  }
);


export default API;



