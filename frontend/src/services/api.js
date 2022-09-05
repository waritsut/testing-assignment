import axios from 'axios';

const API = axios.create({
  baseURL: process.env.VUE_APP_API_HOST_NAME
});

export default API;
