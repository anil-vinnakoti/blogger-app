import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:8080/api", // Gin backend
  withCredentials: true
});

export default api;
